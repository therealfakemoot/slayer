package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	// conf "git.ndumas.com/ndumas/slayer/conf"
	issues "git.ndumas.com/ndumas/slayer/issues"
	sla "git.ndumas.com/ndumas/slayer/sla"
)

func main() {
	config := *flag.String("config", "conf.toml", "config file path")
	debug := *flag.Bool("debug", false, "debug mode")
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

	var cr sla.ComplianceReport
	for name, target := range conf.Targets {
		targetCtx := log.WithFields(log.Fields{
			"name":   name,
			"target": target,
		})
		ic := make(sla.IssueCompliance)
		// this may be the time to start using errgroup and goroutines
		allIssues, err := issues.Get(target, conf.Auth)
		if err != nil {
			targetCtx.WithError(err).Error("unable to fetch issues")
		}
		for _, rule := range target.Rules {
			for _, i := range allIssues {
				ic[rule.Key()] = rule.Check(i)
			}
		}
		fmt.Printf("%s: %+v\n", name, target)
	}

	fmt.Printf("%+v", cr)
}
