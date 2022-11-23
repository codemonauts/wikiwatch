package main

import (
	"encoding/json"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
)

type Organisation struct {
	Name     string   `json:"name"`
	CIDRList []string `json:"ranges"`
	Networks []*net.IPNet
}

func loadRanges(config *Config) {
	log.WithField("filename", config.OrganisationsFile).Info("Loading ip ranges from a seperate file")

	config.Organisations = nil

	data, err := os.ReadFile(config.OrganisationsFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &config.Organisations)
	if err != nil {
		log.Fatal(err)
	}

	for i := range config.Organisations {
		for _, cidr := range config.Organisations[i].CIDRList {
			_, net, err := net.ParseCIDR(cidr)
			if err != nil {
				log.WithFields(log.Fields{
					"organisation": config.Organisations[i].Name,
					"cidr":         cidr,
				}).Error("Could not parse given CIDR notation")
			}
			config.Organisations[i].Networks = append(config.Organisations[i].Networks, net)
		}
	}
	log.Debugf("Loaded %d organisations", len(config.Organisations))
}
