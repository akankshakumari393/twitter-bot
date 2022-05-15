package auth

import (
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func GetTwitterV1Client(creds *Credentials) (*twitter.Client, error) {
	// consumer details
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	// user token
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)
	// create an httpClient
	httpClient := config.Client(oauth1.NoContext, token)
	// create a twiiter Client
	client := twitter.NewClient(httpClient)
	// verify Credentials
	verifyParam := &twitter.AccountVerifyParams{}
	_, _, err := client.Accounts.VerifyCredentials(verifyParam)
	if err != nil {
		log.Fatalf("failed to verify client %s", err)
	}
	// log.Default().Printf("%+v", user)
	return client, nil
}

// not working as the app currently doesn't use OAuth2.0

func GetTwitterV2Client(creds *Credentials) (*twitter.Client, error) {
	// consumer details
	config := &clientcredentials.Config{
		ClientID:     creds.ConsumerKey,
		ClientSecret: creds.ConsumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext)

	// Twitter client
	client := twitter.NewClient(httpClient)
	// verify Credentials
	verifyParam := &twitter.AccountVerifyParams{}
	_, _, err := client.Accounts.VerifyCredentials(verifyParam)
	if err != nil {
		log.Fatalf("failed to verify client %s", err)
	}
	// log.Default().Printf("%+v", user)
	return client, nil
}
