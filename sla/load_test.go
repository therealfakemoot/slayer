package sla

import (
	"strings"
	"testing"
)

func TestLoadRules(t *testing.T) {
	t.Run("good load", func(t *testing.T) {
		r := strings.NewReader(`[auth]
user= "your.user"
token = "bad token"
password= "BadPassword1"

[targets]

[targets.REMAX]
name = "RE/MAX"
board = 0
filter = 0
rules = ["updated 48h", "group any"]`)

		config, err := Load(r)

		if err != nil {
			t.Logf("%s", err)
			t.Fail()
		}

		remax, ok := config.Targets["REMAX"]
		if !ok {
			t.Logf("target REMAX not parsed")
			t.Fail()
		}

		if len(remax.Rules) != 2 {
			t.Logf("expected 2 rules, got %d", len(remax.Rules))
			t.Fail()
		}
	})
}
