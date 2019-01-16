package sla

import (
	client "github.com/therealfakemoot/slayer/client"
)

type Checker struct {
	Service  *client.IssueService
	Enforcer Enforcer
	Renderer Renderer
}

func (c *Checker) Render() string {
	return c.Renderer.Render(c.Enforcer.Report(c.Service.Get()))
}

func (c *Checker) Report() ComplianceReport {
	return c.Enforcer.Report(c.Service.Get())
}

type Renderer interface {
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
