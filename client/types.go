package client

import (
	jira "github.com/andygrunwald/go-jira"

	sla "git.ndumas.com/ndumas/slayer/sla"
)

type IssueService interface {
	Get(sla.Target) chan jira.Issue
	Board(sla.Target) error
	Filter(sla.Target) error
}
