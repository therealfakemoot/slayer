package client

import (
	//"fmt"
	"net/http"
	"net/url"
	"time"

	jira "github.com/andygrunwald/go-jira"
	log "github.com/sirupsen/logrus"

	client "github.com/therealfakemoot/slayer/client"
	sla "github.com/therealfakemoot/slayer/sla"
)

// JiraService is a provider for issues
type JiraService struct {
	Base   *url.URL
	Auth   client.AuthOptions
	Client *http.Client
}

func New(auth client.AuthOptions, base string, duration time.Duration) *JiraService {
	c := http.Client{
		Timeout: duration,
	}

	u, err := url.Parse(base)
	if err != nil {
		log.WithError(err).Error("invalid base URL")
	}

	js := JiraService{
		Base:   base,
		Auth:   u,
		Client: &c,
	}

	return &js
}

func (js *JiraService) prepRequest(target string) (req *http.Request) {

	endpoint := js.Base.ResolveReference()

	prepCtx := log.WithFields(log.Fields{
		"endpoint": endpoint,
	})
	req, _ = http.NewRequest("GET", endpoint.String(), nil)
	req.Header.Add("Authorization", "Basic bmljaG9sYXMuZHVtYXNAaG9tZXMuY29tOndLN2tIaU9SS3RmSkd4VFZ3MWwzODIxRQ")

	return req
}

// Get operates on an SLA Target and retrieves all issues matching those targets.
func (js JiraService) Get(target sla.Target) (issues []jira.Issue) {
	getCtx := log.WithFields(log.Fields{
		"board":  target.Board,
		"filter": target.Filter,
	})

	if target.Board != 0 {
		boardIssues, err := js.Board(target.Board)
		if err != nil {
			getCtx.WithError(err).Error("couldn't fetch board issues")
		}
		issues = append(issues, boardIssues...)
	}

	if target.Filter != 0 {
		filterIssues, err := js.Board(target.Filter)
		if err != nil {
			getCtx.WithError(err).Error("couldn't fetch filter issues")
		}
		issues = append(issues, filterIssues...)
	}

	return issues
}

// Board will fetch all issues attached to a given board.
func (js *JiraService) Board(id int) (issues []jira.Issue, err error) {
	// boardEndpoint := url.URL{Path: fmt.Sprintf("rest/agile/1.0/board/%d/issue", id)}

	return issues, err
}

// BoardBacklog will fetch all backlogged issues attached to a given board.
func (js *JiraService) BoardBacklog(id int) (issues []jira.Issue, err error) {

	return issues, err
}

// Filter fetches all issues returned by a specific filter.
func (js *JiraService) Filter(id int) (issues []jira.Issue, err error) {
	// filterEndpoint := url.URL{Path: fmt.Sprintf("/api/2/filter/%d", id)}

	return issues, err
}
