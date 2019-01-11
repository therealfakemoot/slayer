package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	jira "github.com/andygrunwald/go-jira"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	client "git.ndumas.com/ndumas/slayer/client"
)

type Jira struct {
	Auth   client.AuthOptions
	Base   *url.URL
	Client *http.Client
	Issues chan jira.Issue
}

func (jc *Jira) Do(req *http.Request) (resp *http.Response, err error) {

	req.Header.Add("Authorization", "Basic bmljaG9sYXMuZHVtYXNAaG9tZXMuY29tOndLN2tIaU9SS3RmSkd4VFZ3MWwzODIxRQ")

	return jc.Client.Do(req)
}

func parseResponse(r io.Reader) (br BoardResponse, err error) {
	err = json.NewDecoder(r).Decode(&br)

	return br, err
}

func (jc *Jira) Board(board, start, maxResults int) (issues chan jira.Issue) {
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

	br, err = parseResponse(body)

	pageCtx.Debug("channeling page")
	go func() {
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
		go jc.Board(board, i, 50)
	}

	if err != nil {
		boardCtx.WithError(err).Error("unable to parse JSON")
	}
	pageCtx.Debugf("%s\n", br)

	return issues
}
