package sla

import (
	"io"

	// "github.com/spf13/viper"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type Auth struct {
	User     string
	Token    string
	Password string
}

type Config struct {
	Auth    Auth
	Targets map[string]Target
}

func Load(r io.Reader) (c Config, err error) {
	_, err = toml.DecodeReader(r, &c)
	if err != nil {
		return c, errors.Wrap(err, "error parsing configuration")
	}

	return c, nil
}
