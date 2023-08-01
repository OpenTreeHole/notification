package apns

import (
	"crypto/tls"

	"github.com/rs/zerolog/log"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"

	"notification/config"
)

var cert tls.Certificate
var client *apns2.Client

func init() {
	var err error
	cert, err = certificate.FromPemFile(config.Config.APNSKeyPath, "")
	if err != nil {
		log.Fatal().Err(err).Str("scope", "init APNs").Msg("APNs cert error")
	}
	if config.Config.Mode == "dev" {
		client = apns2.NewClient(cert).Development()
		log.Debug().Msg("init apns; use development mode")
	} else {
		client = apns2.NewClient(cert).Production()
		log.Debug().Msg("init apns; use production mode")
	}
}
