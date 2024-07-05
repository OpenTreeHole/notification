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
	client = newAPNSClient(config.Config.APNSKeyPath, config.Config.APNSKeyPassword)

	// init apns for DanXi v2
	clientV2 = newAPNSClient(config.Config.APNSKeyPathV2, config.Config.APNSKeyPasswordV2)
}

// newAPNSClient create apns client from cert file and password
func newAPNSClient(path string, password string) (client *apns2.Client) {
	var (
		err  error
		cert tls.Certificate
	)

	if strings.HasSuffix(path, ".p12") {
		cert, err = certificate.FromP12File(path, password)
	} else {
		cert, err = certificate.FromPemFile(path, password)
	}

	if err != nil {
		log.Warn().Err(err).Str("scope", "init APNs").Msg("APNs cert error")
		return nil
	}
	if config.Config.Mode == "dev" {
		client = apns2.NewClient(cert).Development()
		log.Debug().Msg("init apns; use development mode")
	} else {
		client = apns2.NewClient(cert).Production()
		log.Debug().Msg("init apns; use production mode")
	}

	return
}
