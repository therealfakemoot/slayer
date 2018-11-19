package main

import (
	"flag"
	"fmt"

	"github.com/andygrunwald/go-jira"
	log "github.com/sirupsen/logrus"
)

const BASE = "https://homesmediasolutions.atlassian.net"

var (
	TOKEN string
	USER  string
)

func main() {
	flag.StringVar(&TOKEN, "token", "", "auth token")
	flag.StringVar(&USER, "user", "", "username")
	flag.Parse()

	tp := jira.BasicAuthTransport{
		Username: USER,
		Password: TOKEN,
	}

	authCtx := log.WithFields(log.Fields{
		"user": USER,
		"base": BASE,
	})

	jc, err := jira.NewClient(tp.Client(), BASE)
	if err != nil {
		authCtx.WithError(err).Error("unable to create client")
		return
	}

	// proj := "RMXC"

	issues, r, err := jc.Issue.Search("project = RMXC", nil)
	fmt.Printf("%+v", r.Request)
	if err != nil {
		authCtx.WithError(err).Error("unable to fetch issues")
		return
	}

	fmt.Println(issues)
}
