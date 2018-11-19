package main

import (
	"flag"
	"fmt"

	"github.com/andygrunwald/go-jira"
	log "github.com/sirupsen/logrus"
)

const BASE = "https://homesmediasolutions.atlassian.net"

var (
	USER  string
	TOKEN string
)

func main() {
	flag.StringVar(&USER, "user", "", "username")
	flag.StringVar(&TOKEN, "token", "", "auth token")
	flag.Parse()

	tp := jira.BasicAuthTransport{
		Username: USER,
		Password: TOKEN,
	}

	authCtx := log.WithFields(log.Fields{
		"user":  USER,
		"base":  BASE,
		"token": TOKEN,
	})

	jc, err := jira.NewClient(tp.Client(), BASE)
	if err != nil {
		authCtx.WithError(err).Error("unable to create client")
		return
	}

	// proj := "RMXC"

	issues, r, err := jc.Issue.Search("project = RMXC", nil)
	if err != nil {
		authCtx.WithFields(log.Fields{
			"request":  fmt.Sprintf("%+v", r.Request),
			"response": fmt.Sprintf("%+v", r),
		}).WithError(err).Error("unable to fetch issues")
		return
	}

	fmt.Println(issues)
}
