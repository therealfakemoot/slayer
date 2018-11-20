package sla

import (
	jira "github.com/andygrunwald/go-jira"
	log "github.com/sirupsen/logrus"
)

type Rule func(i *jira.Issue) bool

func Enforce(rules []Rule, i *jira.Issue) bool {
	for _, c := range rules {
		if !c(i) {
			return false
		}
	}
	return true
}

func UpdatedWithin(d *time.Duration) Rule {
	return func(i *jira.Issue) bool {
		return time.Now().Sub(i.Fields.Updated) < d
	}
}

func AssignedToGroup(g string) Rule {
	return func(i *jira.Issue) bool {
		return false
	}
}
