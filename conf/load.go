package conf

import (
	"io"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"

	client "github.com/therealfakemoot/slayer/client"
)

func LoadAuth(r io.Reader) (ac client.AuthConfig) {
	if _, err := toml.DecodeReader(r, ac); err != nil {
		log.WithError(err).Error("unable to load auth data")
	}
	return ac
}
