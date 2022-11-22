package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net"
	"os"
	"time"

	"github.com/r3labs/sse/v2"
)

const (
	minDiffLength = 10
)

var (
	organisations []Organisation
)

type Config struct {
}

type Organisation struct {
	Name     string   `json:"name"`
	CIDRList []string `json:"ranges"`
	Ranges   []*net.IPNet
}

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

func buildDiffURL(event RecentChange) string {
	return fmt.Sprintf(
		"https://%s/w/index.php?diff=%d&oldid=%d",
		event.ServerName,
		event.Revision.New,
		event.Revision.Old,
	)
}

func getDiffLength(event RecentChange) int {
	return int(math.Abs(float64(event.Length.Old - event.Length.New)))
}

func loadRanges() {
	data, err := os.ReadFile("newRanges.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &organisations)
	if err != nil {
		panic(err)
	}

	for i := range organisations {
		for _, cidr := range organisations[i].CIDRList {
			_, net, err := net.ParseCIDR(cidr)
			if err != nil {
				fmt.Println(err)
			}
			organisations[i].Ranges = append(organisations[i].Ranges, net)
		}
	}
}

func handleAnonymousEdit(event RecentChange) {
	editIP := net.ParseIP(event.User)
	fmt.Printf("Anonymous Edit from %s\n", editIP)

	for _, org := range organisations {
		for _, network := range org.Ranges {
			//fmt.Printf("Checking if %v contains %v\n", network, editIP)
			if network.Contains(editIP) {
				fmt.Printf("Found an anonymous edit from %q\n", org.Name)
				//TODO Send toot and tweet
			}

		}
	}
}

func main() {
	loadRanges()

	client := sse.NewClient("https://stream.wikimedia.org/v2/stream/recentchange")
	fmt.Println("Subscribing...")
	client.Subscribe("recentchange", func(msg *sse.Event) {
		// Got some data!
		var event RecentChange
		json.Unmarshal(msg.Data, &event)
		editIP := net.ParseIP(event.User)
		diffLength := getDiffLength(event)
		if event.Type == "edit" && event.Bot == false && editIP != nil && diffLength > minDiffLength {
			//fmt.Println(string(msg.Data))
			// fmt.Printf("%+v\n", event)
			// fmt.Println(buildDiffURL(event))
			handleAnonymousEdit(event)
		}

	})
}
