package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	sla "github.com/therealfakemoot/slayer/sla"
)

type FlatFileService struct {
	Store string
}

func (ffs FlatFileService) Get() (issues []sla.Issue) {
	f, err := os.Open(ffs.Store)
	if err != nil {
		log.Fatal("could not open config file")
	}

	err = json.NewDecoder(f).Decode(&issues)
	if err != nil {
		log.Fatal("could not parse issues")
	}

	return issues
}

type Rule struct {
}

type Enforcer struct {
	Rules []Rule
}

func LoadRules(r io.Reader) (err error) {
	return err
}

func (e Enforcer) Report(is *sla.IssueService) (cr sla.ComplianceReport) {

	return cr
}

func main() {
	var store = flag.String("store", "issues.json", "issues file")

	flag.Parse()

	ffs := FlatFileService{Store: *store}
	fmt.Println(ffs.Get())
}
