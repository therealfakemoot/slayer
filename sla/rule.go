package sla

import (
	"errors"
	"fmt"
	"strings"
	"time"

	jira "github.com/andygrunwald/go-jira"
	log "github.com/sirupsen/logrus"
)

var (
	ErrBadField      = errors.New("unrecognized field")
	ErrBadConstraint = errors.New("unable to parse constraint")
)

type RuleSet []Rule

// Apply applies a set of Rules to a set of issues. Returns a Compliance report detailing which rules and which issues did not meet SLA.
func (rs RuleSet) Apply(issues []jira.Issue) ComplianceReport {
	var cr ComplianceReport
	for _, i := range issues {
		var ic IssueCompliance
		for _, r := range rs {
			issueCtx := log.WithFields(log.Fields{
				"key": i.Key,
			})

			issueCtx.Debug("enforcing SLA")
			ic[r.Key()] = r.Check(i)
		}
		cr[i.Key] = ic
	}
	return cr
}

// Rule is a subcomponent of an SLA.
type Rule struct {
	Raw   string // original string used to build the rule
	Check RuleFunc
}

func (r Rule) Key() string {
	return r.Raw
}

type RuleFunc func(i jira.Issue) bool

func ParseRule(s string) (Rule, error) {
	var r Rule
	r.Raw = s

	failure := func(i jira.Issue) bool { return false }
	r.Check = failure

	raw := strings.Split(s, " ")

	field, constraint := raw[0], strings.Join(raw[1:], " ")

	parseCtx := log.WithFields(log.Fields{
		"field":      fmt.Sprintf("%q", field),
		"constraint": fmt.Sprintf("%q", constraint),
		"split":      fmt.Sprintf("%q", raw),
		"original":   fmt.Sprintf("%q", s),
	})

	switch field {
	case "updated":
		d, err := time.ParseDuration(constraint)
		if err != nil {
			return r, ErrBadConstraint
		}
		r.Check = UpdatedWithin(d)
		return r, nil
	case "group":
		r.Check = AssignedToGroup(constraint)
		return r, nil
	default:
		parseCtx.WithError(ErrBadField).Error("rule parse failure")
		return r, ErrBadField
	}
}

// UpdatedWithin asserts that an Issue has been updated with the duration d.
func UpdatedWithin(d time.Duration) RuleFunc {
	return func(i jira.Issue) bool {
		return time.Since(time.Time(i.Fields.Updated)) < d
	}
}

// AssignedToGroup asserts that an Issue is assigned to a specific development group matching the provided string.
func AssignedToGroup(g string) RuleFunc {
	return func(i jira.Issue) bool {
		return false
	}
}
