package apns

import (
	"crypto/tls"
	"fmt"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
	"log"
	"notification/config"
	"notification/utils"
)

var cert tls.Certificate
var client *apns2.Client

func init() {
	fmt.Println("init apns")
	var err error
	cert, err = certificate.FromPemFile(
		utils.ToAbsolutePath(config.Config.APNSKeyPath),
		"",
	)
	if err != nil {
		log.Fatal("Cert Error: ", err)
	}
	client = apns2.NewClient(cert).Production()
}
