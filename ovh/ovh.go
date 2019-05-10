package ovh

import (
	"github.com/ovh/go-ovh/ovh"
	log "github.com/sirupsen/logrus"
)

func GetOvhClient() *ovh.Client {
	config := GetCurrentAPIConfig()
	client, err := ovh.NewClient(
		config.Endpoint,
		config.ApplicationKey,
		config.ApplicationSecret,
		config.ConsumerKey,
	)
	if err != nil {
		log.Fatalf("Unable to create ovh client: %v", err)
	}
	return client
}
