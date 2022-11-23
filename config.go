package main

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Name              string          `json:"name"`
	Organisations     []Organisation  `json:"organisations"`
	OrganisationsFile string          `json:"organisations_file"`
	MinDiff           int             `json:"min_diff"`
	Mastodon          MastodonAccount `json:"mastodon"`
	Twitter           TwitterAccount  `json:"twitter"`
}

type MastodonAccount struct {
	Server       string `json:"server"`
	ClientKey    string `json:"client_key"`
	ClientSecret string `json:"client_secret"`
	AccessToken  string `json:"access_token"`
}

type TwitterAccount struct {
	ConsumerKey    string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
	AccessToken    string `json:"access_token"`
	AccessSecret   string `json:"access_secret"`
}

func (c Config) hasMastodon() bool {
	if c.Mastodon.Server != "" &&
		c.Mastodon.ClientKey != "" &&
		c.Mastodon.ClientSecret != "" &&
		c.Mastodon.AccessToken != "" {
		return true
	} else {
		return false
	}
}

func (c Config) hasTwitter() bool {
	if c.Twitter.ConsumerKey != "" &&
		c.Twitter.ConsumerSecret != "" &&
		c.Twitter.AccessToken != "" &&
		c.Twitter.AccessSecret != "" {
		return true
	} else {
		return false
	}
}

func (c Config) loadFile(path string) *Config {
	log.WithField("filename", path).Debug("Loading config file")

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
	return &c
}

func NewConfig() Config {
	return Config{
		Name:    "Default bot",
		MinDiff: 10,
	}
}
