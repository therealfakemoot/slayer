package main

import (
	"flag"
	"fmt"

	jira "github.com/andygrunwald/go-jira"
	log "github.com/sirupsen/logrus"
)

// BASE is the URL for the homes.com Jira instance.
const BASE = "https://homesmediasolutions.atlassian.net"

func main() {
	// conf := flag.String("config", ".slayer", "config file path")
	user := flag.String("user", "", "username")
	token := flag.String("token", "", "auth token")
	debug := flag.Bool("debug", false, "debug mode")
	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	tp := jira.BasicAuthTransport{
		Username: *user,
		Password: *token,
	}

	authCtx := log.WithFields(log.Fields{
		"user":  user,
		"base":  BASE,
		"token": token,
	})

	jc, err := jira.NewClient(tp.Client(), BASE)
	if err != nil {
		authCtx.WithError(err).Error("unable to create client")
		return
	}

	issues, r, err := jc.Issue.Search("project = RMXC", nil)
	if err != nil {
		authCtx.WithFields(log.Fields{
			"request":  fmt.Sprintf("%+v", r.Request),
			"response": fmt.Sprintf("%+v", r),
		}).WithError(err).Error("unable to fetch issues")
		return
	}

	fmt.Println(issues[0])
}
