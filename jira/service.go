package client

import (
	"net/http"
	"sync"

	jira "github.com/andygrunwald/go-jira"

	client "git.ndumas.com/ndumas/slayer/client"
	sla "git.ndumas.com/ndumas/slayer/sla"
)

func New(auth client.AuthOptions, base string) *JiraService {
	var js JiraService

	return &js
}

type JiraService struct {
	Auth   client.AuthOptions
	Client *http.Client
}

func (js JiraService) Get(target sla.Target) chan jira.Issue {
	c := make(chan jira.Issue)

	if target.Board != 0 {
		for i := range Board(js.Client, target.Board) {
			c <- i
		}
	}

	if target.Filter != 0 {
		for i := range Filter(js.Client, target.Filter) {
			c <- i
		}
	}

	return c
}
