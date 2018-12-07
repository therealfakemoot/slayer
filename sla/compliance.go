package sla

import (
	"io"
	"strings"
	"text/tabwriter"
	"text/template"

	jira "github.com/andygrunwald/go-jira"
	log "github.com/sirupsen/logrus"
)

var funcMap = template.FuncMap{
	"strip": func(s string) string {
		return strings.TrimSpace(s)
	},
	"trunc": func(s string, n int) string {
		return s[:n]
	},
}

/*
var HeaderTemplate = template.Must(template.New("header").Parse("{{$sep := .Separator}}{{ range .Columns }}{{$sep}}{{.}}{{ end }}{{$sep}}\n"))
var RowTemplate = template.Must(template.New("row").Parse("{{$sep := .Separator}}{{$sep}}{{.Issue.Key}}{{$sep}}{{.Issue.Fields.Summary}}{{$sep}}\n"))
*/

type ComplianceReport map[string]IssueCompliance

func (cr ComplianceReport) Render(rr ReportRenderer) (s string) {
	return rr.Render(cr)
}

type IssueCompliance struct {
	Rules map[string]bool
	Issue jira.Issue
}

type ReportRenderer interface {
	Render(ComplianceReport) string
}

type TermRenderer struct {
	Out io.Writer
}

var TableTemplate = template.Must(template.New("root").Funcs(funcMap).Parse(`{{ $.Separator }}{{ range .Columns }}{{.}}{{ $.Separator }}{{ end }}
{{ range $key, $ic := .Issues }}{{ $.Separator }}{{$key}}{{ $.Separator }}{{ range $rule, $pass := $ic.Rules }}{{$pass}}{{  $.Separator }}{{ end }}
{{end}}
`))

func (tr *TermRenderer) Render(cr ComplianceReport) (s string) {
	var (
		tw tabwriter.Writer
	)

	tw.Init(tr.Out, 15, 4, 0, '\t', 0)
	defer tw.Flush()

	/*
		var issues []jira.Issue
		for _, v := range cr {
			issues = append(issues, v.Issue)
		}
	*/

	templateData := struct {
		Columns   []string
		Issues    ComplianceReport
		Separator string
	}{
		Columns:   []string{"Key", "Updated", "Group"},
		Separator: "\t",
		Issues:    cr,
	}

	err := TableTemplate.Execute(&tw, templateData)

	if err != nil {
		log.WithError(err).Error("unable to render table")
	}

	return s
}
