package client

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"

	sla "git.ndumas.com/ndumas/slayer/sla"
)

type IssueService interface {
	Get(sla.Target) chan jira.Issue
	Board(sla.Target) error
	Filter(sla.Target) error
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
