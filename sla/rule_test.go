package sla

import (
	"testing"
	"time"

	jira "github.com/andygrunwald/go-jira"
)

func TestParseRules(t *testing.T) {
	t.Run("parse errors", func(t *testing.T) {
		cases := []struct {
			actual   string
			expected error
		}{
			{"madeupfield 15d", ErrBadField},
			{"updated garbagedata", ErrBadConstraint},
		}

		for _, tt := range cases {
			_, err := ParseRule(tt.actual)

			if err != tt.expected {
				t.Fail()
			}
		}
	})

	t.Run("basic parsing", func(t *testing.T) {
		cases := []struct {
			actual string
		}{
			{"updated 24h"},
			{"group salesforce"},
			{"group Homes Connect"},
		}

		for _, tt := range cases {
			_, err := ParseRule(tt.actual)

			if err != nil {
				t.Logf("received: %s", err)
				t.Fail()
			}
		}
	})
}

func TestRuleApplication(t *testing.T) {
	t.Run("updated", func(t *testing.T) {
		var (
			i jira.Issue
			f jira.IssueFields
		)
		i.Fields = &f
		i.Fields.Updated = jira.Time(time.Now().AddDate(0, 0, -5))

		cases := []struct {
			actual   string
			expected bool
		}{
			{"updated 24h", false},
			{"updated 72h", false},
			{"updated 168h", true},
		}

		for _, tt := range cases {
			r, err := ParseRule(tt.actual)

			if err != nil {
				t.Logf("received: %s", err)
				t.Fail()
			}

			if r(&i) != tt.expected {
				t.Fail()
			}
		}

	})
}
