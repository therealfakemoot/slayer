package sla

import (
	client "git.ndumas.com/ndumas/slayer/client"
)

type Checker struct {
	Service  *client.IssueService
	Enforcer Enforcer
	Renderer ReportRenderer
}

func (c *Checker) Report() string {
	return c.Renderer.Render(c.Service.Get())
}

type ReportRenderer interface {
	Render(ComplianceReport) string
}

type Enforcer interface {
	Report(*client.IssueService) ComplianceReport
}

type ComplianceReport map[string]IssueCompliance

type IssueCompliance struct {
	Rules map[string]bool
	Meta  map[string]string
}
