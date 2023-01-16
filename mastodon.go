package main

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/mattn/go-mastodon"
)

func sendToot(org Organisation, event RecentChange, account MastodonAccount) {
	text := fmt.Sprintf("%q wurde anonym aus dem Netz %s bearbeitet.\n%s", event.Title, org.Name, event.getDiffURL())

	if len(text) > 500 {
		log.Error("Toot text is  to long!")
		return
	}

	c := mastodon.NewClient(&mastodon.Config{
		Server:       account.Server,
		ClientID:     account.ClientKey,
		ClientSecret: account.ClientSecret,
		AccessToken:  account.AccessToken,
	})

	_, err := c.PostStatus(context.Background(), &mastodon.Toot{
		Status:     text,
		Visibility: "unlisted",
	})

	if err != nil {
		log.Error("Couldn't send toot:", err)
	}
}
