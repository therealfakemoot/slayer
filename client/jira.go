package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"

	jira "github.com/andygrunwald/go-jira"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	sla "git.ndumas.com/ndumas/slayer/sla"
)

type Jira struct {
	Auth   AuthOptions
	Base   *url.URL
	Client *http.Client
	Issues chan jira.Issue
}

func (jc *Jira) Do(req *http.Request) (resp *http.Response, err error) {

	req.Header.Add("Authorization", "Basic bmljaG9sYXMuZHVtYXNAaG9tZXMuY29tOndLN2tIaU9SS3RmSkd4VFZ3MWwzODIxRQ")

	return jc.Client.Do(req)
}

func (jc *Jira) Board(board, start, maxResults int, wg *sync.WaitGroup) {
	var br BoardResponse

	endpoint := &url.URL{Path: fmt.Sprintf("rest/agile/1.0/board/%d/issue", board)}
	q := endpoint.Query()
	q.Add("startAt", fmt.Sprintf("%d", start))
	q.Add("maxResults", fmt.Sprintf("%d", maxResults))
	endpoint.RawQuery = q.Encode()

	endpoint = jc.Base.ResolveReference(endpoint)

	payload := strings.NewReader("")
	boardCtx := log.WithFields(log.Fields{})

	req, err := http.NewRequest("GET", endpoint.String(), payload)
	if err != nil {
		err = errors.Wrap(err, "unable to create board request")
		boardCtx.WithError(err).Error("unable to create board request")
	}

	pageCtx := boardCtx.WithFields(log.Fields{
		"start":    start,
		"pageSize": maxResults,
	})

	pageCtx.Debug("fetching page")

	resp, err := jc.Do(req)
	if err != nil {
		err = errors.Wrap(err, "unable to execute API request")
		boardCtx.WithError(err).Error("yikes")
	}

	defer resp.Body.Close()
	defer req.Body.Close()
	pageCtx.Debug("parsing page")

	rawBody, _ := ioutil.ReadAll(resp.Body)
	body := strings.NewReader(string(rawBody))

	err = json.NewDecoder(body).Decode(&br)

	pageCtx.Debug("channeling page")
	go func() {
		defer wg.Done()
		for _, i := range br.Issues {
			jc.Issues <- i
			pageCtx.WithFields(log.Fields{
				"key": i.Key,
			}).Debug("issue pushed to channel")
		}
	}()

	if start != 0 {
	}

	pageCtx.Debug("recursing into Board")

	for i := br.MaxResults; i < br.Total; i += br.MaxResults {
		go jc.Board(board, i, 50, wg)
	}

	if err != nil {
		boardCtx.WithError(err).Error("unable to parse JSON")
	}
	pageCtx.Debugf("%s\n", br)
}

func (jc *Jira) Get(t sla.Target) (err error) {
	var wg sync.WaitGroup

	board := t.Board
	filter := t.Filter

	if board == 0 && filter == 0 {
		return errors.New("no filter or board provided")
	}

	authCtx := log.WithFields(log.Fields{
		"user":  jc.Auth.User,
		"base":  jc.Base,
		"token": jc.Auth.Token,
	})

	if err != nil {
		authCtx.WithError(err).Error("unable to create client")
		return err
	}

	if board != 0 {
		wg.Add(1)
		go jc.Board(board, 0, 50, &wg)
	}

	if filter != 0 {
	}

	wg.Wait()
	close(jc.Issues)
	return nil
}
