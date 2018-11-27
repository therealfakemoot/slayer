package issues

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	sla "git.ndumas.com/ndumas/slayer/sla"
)

// BASE is the URL for the homes.com Jira instance.
const BASE = "https://homesmediasolutions.atlassian.net"

func Get(t sla.Target, auth sla.Auth) (issues []jira.Issue, err error) {
	board := t.Board
	filter := t.Filter

	if board == 0 && filter == 0 {
		return issues, errors.New("no filter or board provided")
	}

	if board != 0 {

	}

	if filter != 0 {

	}

	return issues, nil
}

func project(user, token, proj string) ([]jira.Issue, error) {
	var issues []jira.Issue
	authCtx := log.WithFields(log.Fields{
		"user":  user,
		"base":  BASE,
		"token": token,
	})

	tp := jira.BasicAuthTransport{
		Username: user,
		Password: token,
	}

	jc, err := jira.NewClient(tp.Client(), BASE)
	if err != nil {
		authCtx.WithError(err).Error("unable to create client")
		return issues, err
	}

	issues, r, err := jc.Issue.Search("project = "+proj, nil)
	if err != nil {
		authCtx.WithFields(log.Fields{
			"request":  fmt.Sprintf("%+v", r.Request),
			"response": fmt.Sprintf("%+v", r),
		}).WithError(err).Error("unable to fetch issues")
		return issues, err
	}

	return issues, nil
}
