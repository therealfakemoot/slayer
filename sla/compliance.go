package sla

import (
	"os"
)

type ComplianceReport map[string]IssueCompliance

type IssueCompliance map[string]bool

type ReportRenderer interface {
	Render() error
}

type TerminalRenderer struct {
	o *os.File
}

func (tr TerminalRenderer) Render() error {
	tr.o.Write([]byte(""))
	return nil
}
