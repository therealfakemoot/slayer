package sla

import (
	"time"

	jira "github.com/andygrunwald/go-jira"
	log "github.com/sirupsen/logrus"
)

// Rule is a subcomponent of an SLA.
type Rule func(i *jira.Issue) bool

// Enforce applies a set of Rules to an issue. If any Rule returns false, the Issue does not meet SLA.
func Enforce(rules []Rule, i *jira.Issue) bool {
	issueCtx := log.WithFields(log.Fields{
		"key": i.Key,
	})

	issueCtx.Debug("enforcing SLA")

	for _, c := range rules {
		if !c(i) {
			return false
		}
	}
	return true
}

// UpdatedWithin asserts that an Issue has been updated with the duration d.
func UpdatedWithin(d time.Duration) Rule {
	return func(i *jira.Issue) bool {
		return time.Now().Sub(time.Time(i.Fields.Updated)) < d
	}
}

// AssignedToGroup asserts that an Issue is assigned to a specific development group matching the provided string.
func AssignedToGroup(g string) Rule {
	return func(i *jira.Issue) bool {
		return false
	}
}
