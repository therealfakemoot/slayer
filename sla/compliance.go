package sla

import (
	jira "github.com/andygrunwald/go-jira"
	// log "github.com/sirupsen/logrus"

	client "git.ndumas.com/ndumas/slayer/client"
)

type Checker struct {
	Service  client.IssueService
	Enforcer Enforcer
}

type Enforcer interface {
	Report(client.IssueService) ComplianceReport
}

type ComplianceReport map[string]IssueCompliance

func (cr ComplianceReport) Render(rr ReportRenderer) (s string) {
	return rr.Render(cr)
}

type IssueCompliance struct {
	Rules map[string]bool
	Issue jira.Issue
}
