package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	jira "github.com/andygrunwald/go-jira"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	sla "git.ndumas.com/ndumas/slayer/sla"
)

type Jira struct {
	Auth   AuthOptions
	Base   *url.URL
	Client *http.Client
}

func (jc *Jira) Do(req *http.Request) (resp *http.Response, err error) {

	req.Header.Add("Authorization", "Basic bmljaG9sYXMuZHVtYXNAaG9tZXMuY29tOndLN2tIaU9SS3RmSkd4VFZ3MWwzODIxRQ")

	return jc.Client.Do(req)
}

func (jc *Jira) Get(t sla.Target, auth sla.Auth) (issues []jira.Issue, err error) {
	board := t.Board
	filter := t.Filter

	var br BoardResponse

	if board == 0 && filter == 0 {
		return issues, errors.New("no filter or board provided")
	}

	authCtx := log.WithFields(log.Fields{
		"user":  auth.User,
		"base":  jc.Base,
		"token": auth.Token,
	})

	/*
		tp := jira.BasicAuthTransport{
			Username: auth.User,
			Password: auth.Token,
		}

		jc, err := jira.NewClient(tp.Client(), BASE)
	*/

	if err != nil {
		authCtx.WithError(err).Error("unable to create client")
		return issues, err
	}

	if board != 0 {
		// var boardIssues []jira.Issue
		// boardIssues := new([]jira.Issue)

		endpoint := &url.URL{Path: fmt.Sprintf("rest/agile/1.0/board/%d/issue", board)}
		endpoint = jc.Base.ResolveReference(endpoint)

		payload := strings.NewReader("")
		boardCtx := authCtx.WithFields(log.Fields{
			"endpoint": endpoint,
		})

		// req, err := jc.NewRequest("GET", endpoint, payload)
		req, err := http.NewRequest("GET", endpoint.String(), payload)
		if err != nil {
			err = errors.Wrap(err, "unable to create board request")
			boardCtx.WithError(err).Error("unable to create board request")
			return issues, err
		}

		// _, err = jc.Do(req, boardIssues)
		resp, err := jc.Do(req)
		if err != nil {
			err = errors.Wrap(err, "unable to retrieve issues for board")
			boardCtx.WithError(err).Error("fuck")
		}
		// issues = append(issues, *boardIssues...)
		defer resp.Body.Close()
		defer req.Body.Close()
		rawBody, _ := ioutil.ReadAll(resp.Body)
		body := strings.NewReader(string(rawBody))
		err = json.NewDecoder(body).Decode(&br)
		if err != nil {
			boardCtx.WithError(err).Error("unable to parse JSON")
		}
		fmt.Printf("%+v\n", br.Issues)
		// fmt.Printf("%s\n", rawBody)
	}

	if filter != 0 {

	}

	return issues, nil
}
