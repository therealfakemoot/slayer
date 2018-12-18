package client

import (
	jira "github.com/andygrunwald/go-jira"
)

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
