package main

import (
	// "bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BurntSushi/toml"
	// "github.com/pkg/errors"
	// jira "github.com/andygrunwald/go-jira"
	log "github.com/sirupsen/logrus"
)

type conf struct {
	Auth auth
}

type auth struct {
	User  string `toml:"user"`
	Token string `toml:"token"`
}

func (a auth) String() string {
	return fmt.Sprintf("%s/%s", a.User, a.Token)
	// return fmt.Sprintf("%+v", a)
}

const BASE = "https://homesmediasolutions.atlassian.net/"

func loadConf() conf {
	var c conf

	m, err := toml.DecodeFile("conf.toml", &c)

	if err != nil {
		log.WithFields(log.Fields{
			"meta": m,
		}).WithError(err).Fatal("unable to load config data")
	}

	return c
}

func main() {
	log.SetLevel(log.DebugLevel)

	log.Debug("slayer starting")
	defer log.Debug("slayer terminating")

	conf := loadConf()
	log.WithFields(log.Fields{
		"config": conf,
	}).Debug("loaded config")

	t1 := fmt.Sprintf("%s:%s", conf.Auth.User, conf.Auth.Token)
	log.WithFields(log.Fields{
		"user:token": t1,
	}).Debug("token gen phase 1")

	t2 := []byte(t1)
	authToken := base64.StdEncoding.EncodeToString(t2)
	log.WithFields(log.Fields{
		"user:token": authToken,
	}).Debug("token gen phase 2")

	c := http.Client{}
	req, err := http.NewRequest("GET", BASE+"agile/1.0/board/67/backlog", nil)

	if err != nil {
		log.WithError(err).Fatal("could not prepare request")
		// log.Fatalf("could not prepare request: %s", err)
	}

	req.Header.Add("Authorization", "Basic "+authToken)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(conf.Auth.User, string(conf.Auth.Token))
	log.WithFields(log.Fields{
		"headers": req.Header,
	}).Debug("request prepared")

	resp, err := c.Do(req)
	defer resp.Body.Close()

	if err != nil {
		log.WithError(err).Fatal("error during http request")
		// log.Fatalf("error during http request: %s", err)
	}

	switch resp.StatusCode {
	case 401:
		log.WithFields(log.Fields{
			"url":     resp.Request.URL,
			"headers": resp.Request.Header,
			"host":    resp.Request.Host,
		}).Debug("request failed")
		log.WithFields(log.Fields{
			"code":    resp.StatusCode,
			"headers": resp.Header,
			// "body":    ioutil.ReadAll(resp.Body),
		}).Error("authorization error")
		// log.Printf("request: %+v", resp.Request)
		// log.Printf("response: %+v", resp)
		// log.Fatalf("authorization issue")
	}

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithError(err).Fatal("unable to read response body")
		// log.Fatalf("unable to read response body: %s", err)
	}

	log.WithFields(log.Fields{
		"body": string(raw),
	}).Debug("raw body")
}
