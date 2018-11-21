package conf

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	// "github.com/spf13/viper"

	sla "git.ndumas.com/ndumas/slayer/sla"
	"github.com/pkg/errors"
)

func LoadRules(r io.Reader) ([]sla.Rule, error) {
	var rules []sla.Rule

	raw, err := ioutil.ReadAll(r)
	if err != nil {
		return rules, err
	}

	for idx, raw := range strings.Split(string(raw), "\n") {
		r, err := sla.ParseRule(raw)
		if err != nil {
			return rules, errors.Wrap(err, fmt.Sprintf("could not parse rule %d", idx))
		}
		rules = append(rules, r)
	}

	return rules, nil
}
