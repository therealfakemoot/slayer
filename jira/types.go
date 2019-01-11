package client

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
)

type Board struct {
}

type Filter struct {
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
