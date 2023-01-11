package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	log "github.com/sirupsen/logrus"
)

func sendTweet(org Organisation, event RecentChange, account TwitterAccount) {
	text := fmt.Sprintf("%q wurde anonym aus derm Netz %s bearbeitet.\n%s", event.Title, org.Name, event.getDiffURL())

	if len(text) > 280 {
		log.Error("Tweet text is  to long!")
		return
	}

	config := oauth1.NewConfig(account.ConsumerKey, account.ConsumerSecret)
	token := oauth1.NewToken(account.AccessToken, account.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	_, _, err := client.Statuses.Update(text, nil)

	if err != nil {
		log.Error("Couldn't send tweet:", err)
	}

}
