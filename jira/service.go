package client

import (
	"net/http"
	"time"

	jira "github.com/andygrunwald/go-jira"
	log "github.com/sirupsen/logrus"

	client "git.ndumas.com/ndumas/slayer/client"
	sla "git.ndumas.com/ndumas/slayer/sla"
)

func New(auth client.AuthOptions, base string, duration time.Duration) *JiraService {
	c := http.Client{
		Timeout: duration,
	}

	js := JiraService{
		Base:   base,
		Auth:   auth,
		Client: &c,
	}

	return &js
}

// JiraService is a provider for issues
type JiraService struct {
	Base   string
	Auth   client.AuthOptions
	Client *http.Client
}

func (js *JiraService) prepRequest(target string) (r *http.Request) {
	r, _ = http.NewRequest("GET", target, nil)

	return r
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

	return issues, err
}

// BoardBacklog will fetch all issues attached to a given board.
func (js *JiraService) BoardBacklog(id int) (issues []jira.Issue, err error) {

	return issues, err
}

// Filter fetches all issues returned by a specific filter.
func (js *JiraService) Filter(id int) (issues []jira.Issue, err error) {

	return issues, err
}
