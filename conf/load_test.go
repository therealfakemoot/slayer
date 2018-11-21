package conf

import (
	"strings"
	"testing"
	// sla "git.ndumas.com/ndumas/slayer/sla"
)

func TestLoadRules(t *testing.T) {
	t.Run("good load", func(t *testing.T) {
		r := strings.NewReader(`group Homes Connect
updated 48h`)

		rules, err := LoadRules(r)

		if len(rules) != 2 {
			t.Logf("expected 2 rules, got %d", len(rules))
			t.Fail()
		}

		if err != nil {
			t.Logf("%s", err)
			t.Fail()
		}
	})
}
