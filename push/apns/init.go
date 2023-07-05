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
	log.Debug().Msg("init apns")
	var err error
	cert, err = certificate.FromPemFile(config.Config.APNSKeyPath, "")
	if err != nil {
		log.Fatal().Err(err).Str("scope", "init apns").Msg("apns Cert Error")
	}
	client = apns2.NewClient(cert).Production()
}
