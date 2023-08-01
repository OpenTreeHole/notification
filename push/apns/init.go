package apns

import (
	"crypto/tls"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"

	"notification/config"
)

var client *apns2.Client
var clientV2 *apns2.Client

func init() {
	// init apns for DanXi v1
	initAPNS(config.Config.APNSKeyPath, &client)

	// init apns for DanXi v2
	initAPNS(config.Config.APNSKeyPathV2, &clientV2)
}

// initAPNS init apns
func initAPNS(path string, client **apns2.Client) {
	var (
		err  error
		cert tls.Certificate
	)

	if strings.HasSuffix(path, ".p12") {
		cert, err = certificate.FromP12File(path, "")
	} else {
		cert, err = certificate.FromPemFile(path, "")
	}

	if err != nil {
		log.Warn().Err(err).Str("scope", "init APNs").Msg("APNs cert error")
		return
	}
	if config.Config.Mode == "dev" {
		*client = apns2.NewClient(cert).Development()
		log.Debug().Msg("init apns; use development mode")
	} else {
		*client = apns2.NewClient(cert).Production()
		log.Debug().Msg("init apns; use production mode")
	}
}
