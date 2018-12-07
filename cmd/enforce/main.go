package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"

	jira "github.com/andygrunwald/go-jira"
	log "github.com/sirupsen/logrus"

	// conf "git.ndumas.com/ndumas/slayer/conf"
	client "git.ndumas.com/ndumas/slayer/client"
	sla "git.ndumas.com/ndumas/slayer/sla"
)

func enforce(c *client.Jira, t map[string]sla.Target) sla.ComplianceReport {
	cr := make(sla.ComplianceReport)
	for name, target := range t {
		targetCtx := log.WithFields(log.Fields{
			"name": name,
		})
		var ic sla.IssueCompliance
		ic.Rules = make(map[string]bool)
		// this may be the time to start using errgroup and goroutines
		go c.Get(target)

		for i := range c.Issues {
			for _, rule := range target.Rules {
				ic.Rules[rule.Key()] = rule.Check(i)
				ic.Issue = i

				targetCtx.WithFields(log.Fields{
					"key":  i.Key,
					"rule": rule.Key(),
				}).Debug("checking issue for compliance")
			}
			targetCtx.Debug("setting IssueCompliance")
			cr[i.Key] = ic
		}
	}

	log.WithFields(log.Fields{
		"resultCount": len(cr),
	}).Debug("compliance report generated")
	// log.Debugf("Compliance Report: %+v\n", cr)

	return cr
}

func main() {
	var (
		config string
		debug  bool
	)

	flag.StringVar(&config, "config", "conf.toml", "config file path")
	flag.BoolVar(&debug, "debug", false, "debug mode")

	flag.Parse()

	if debug {
		log.SetLevel(log.DebugLevel)
	}

	if config == "" {
		log.Error("no config provided")
	}

	f, err := os.Open(config)

	if err != nil {
		log.WithError(err).Error("unable to open rule file")
	}

	conf, err := sla.Load(f)
	if err != nil {
		log.WithError(err).Error("unable to load rules")
	}

	// Change this to parse a URL from the config file.
	base, err := url.Parse("https://homesmediasolutions.atlassian.net")
	if err != nil {
		log.WithError(err).Error("unable to parse Jira URL")
	}

	c := &client.Jira{
		Base:   base,
		Client: &http.Client{},
		Issues: make(chan jira.Issue, 5),
	}

	cr := enforce(c, conf.Targets)
	tr := sla.TermRenderer{Out: os.Stdout}
	fmt.Println(cr.Render(&tr))
}
