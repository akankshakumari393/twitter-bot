package main

import (
	"log"
	"os"

	"github.com/akankshakumari393/twitter-bot/pkg/auth"
)

func main() {
	log.Println("TWITTER BOT")
	creds := auth.Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	// log.Printf("%+v\n", creds)

	client, err := auth.GetTwitterClient(&creds)
	if err != nil {
		log.Fatalf("Error getting Twitter Client %s", err)
	}
	tweet, resp, err := client.Statuses.Update("Hello from bot!", nil)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", tweet)
}
