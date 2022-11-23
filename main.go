package main

import (
	"flag"

	log "github.com/sirupsen/logrus"

	"github.com/r3labs/sse/v2"
)

func main() {
	loglevelFLag := flag.String("loglevel", "INFO", "Set loglevel")
	configFlag := flag.String("config", "./config.json", "Path to the config file")
	flag.Parse()

	l, err := log.ParseLevel(*loglevelFLag)
	if err != nil {
		log.Warn("Unknown loglevel provided. Defaulting to INFO")
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(l)
	}

	config := NewConfig().loadFile(*configFlag)
	if config.OrganisationsFile != "" {
		loadRanges(config)
	}

	if !(config.hasMastodon() || config.hasTwitter()) {
		log.Fatal("Configfile doesn't have a Twitter nor an Mastodon account")

	}

	log.Infof("Starting bot %q\n", config.Name)
	client := sse.NewClient("https://stream.wikimedia.org/v2/stream/recentchange")
	log.Info("Subscribing to recentchange eventstream")
	client.Subscribe("recentchange", func(msg *sse.Event) {
		handleRecentChange(msg, config)
	})
}
