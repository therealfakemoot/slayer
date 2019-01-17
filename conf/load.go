package conf

import (
	"io"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

func LoadAuth(r io.Reader) (ao AuthOptions) {
	var f Full
	if _, err := toml.DecodeReader(r, &f); err != nil {
		log.WithError(err).Error("unable to load auth data")
	}
	return f.Auth
}
