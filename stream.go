package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/r3labs/sse/v2"
)

type RecentChange struct {
	Schema string `json:"$schema"`
	Meta   struct {
		URI       string    `json:"uri"`
		RequestID string    `json:"request_id"`
		ID        string    `json:"id"`
		Dt        time.Time `json:"dt"`
		Domain    string    `json:"domain"`
		Stream    string    `json:"stream"`
		Topic     string    `json:"topic"`
		Partition int       `json:"partition"`
		Offset    int64     `json:"offset"`
	} `json:"meta"`
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Namespace int    `json:"namespace"`
	Title     string `json:"title"`
	Comment   string `json:"comment"`
	Timestamp int    `json:"timestamp"`
	User      string `json:"user"`
	Bot       bool   `json:"bot"`
	Minor     bool   `json:"minor"`
	Length    struct {
		Old int `json:"old"`
		New int `json:"new"`
	} `json:"length"`
	Revision struct {
		Old int `json:"old"`
		New int `json:"new"`
	} `json:"revision"`
	ServerURL        string `json:"server_url"`
	ServerName       string `json:"server_name"`
	ServerScriptPath string `json:"server_script_path"`
	Wiki             string `json:"wiki"`
	Parsedcomment    string `json:"parsedcomment"`
}

func (event RecentChange) getDiffURL() string {
	return fmt.Sprintf(
		"https://%s/w/index.php?diff=%d&oldid=%d",
		event.ServerName,
		event.Revision.New,
		event.Revision.Old,
	)
}

func (event RecentChange) getDiffLength() int {
	return int(math.Abs(float64(event.Length.Old - event.Length.New)))
}
func handleAnonymousEdit(event RecentChange, config *Config) {
	editIP := net.ParseIP(event.User)

	log.WithFields(log.Fields{
		"title": event.Title,
		"type":  event.Type,
		"bot":   event.Bot,
		"user":  event.User,
	}).Debug("Received anonymous Edit")

	for _, org := range config.Organisations {
		for _, network := range org.Networks {
			//fmt.Printf("Checking if %v contains %v\n", network, editIP)
			if network.Contains(editIP) {
				log.WithFields(log.Fields{
					"article": event.Title,
					"IP":      editIP,
					"org":     org.Name,
				}).Info("Found an anonymous edit from a known IP")

				if config.hasMastodon() {
					log.Debug("Config contains a Mastodon account. Sending Toot")
					sendToot(org, event, config.Mastodon)
				}
				if config.hasTwitter() {
					log.Debug("Config contains a Twitter account. Sending Tweet")
					sendTweet(org, event, config.Twitter)
				}

			}

		}
	}
}

func handleRecentChange(msg *sse.Event, config *Config) {
	// Got some data!
	var event RecentChange
	json.Unmarshal(msg.Data, &event)
	editIP := net.ParseIP(event.User)
	diffLength := event.getDiffLength()

	// Check criterya
	if event.Type == "edit" && event.Bot == false && editIP != nil && diffLength > config.MinDiff {
		handleAnonymousEdit(event, config)
	}
}
