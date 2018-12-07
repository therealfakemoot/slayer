package client

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"

	sla "git.ndumas.com/ndumas/slayer/sla"
)

type IssueProvider interface {
	Get(sla.Target, sla.Auth) ([]jira.Issue, error)
	Board(sla.Target) ([]jira.Issue, error)
	Filter(sla.Target) ([]jira.Issue, error)
}

type ResponseMeta struct {
	Expand     string
	StartAt    int
	MaxResults int
	Total      int
}

type BoardResponse struct {
	ResponseMeta
	Issues []jira.Issue
}

func (br BoardResponse) String() string {
	return fmt.Sprintf("%d-%d/%d Issues", br.StartAt, br.StartAt+br.MaxResults, br.Total)
}

type Reporter func([]jira.Issue, []sla.Rule) sla.ComplianceReport
