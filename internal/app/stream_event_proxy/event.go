package stream_event_proxy

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/nicklaw5/helix"
)

type event struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type eventSubNotification struct {
	Subscription helix.EventSubSubscription `json:"subscription"`
	Challenge    string                     `json:"challenge"`
	Event        json.RawMessage            `json:"event"`
}

func generateWebhookSecret() string {
	rand.Seed(time.Now().UnixNano())

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 15) // our random string will be 15 characters
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func NewClient(config *Config) *helix.Client {
	var err error
	client, err := helix.NewClient(&helix.Options{
		ClientID:     config.ClientId,
		ClientSecret: config.ClientSecret,
	})
	if err != nil {
		log.Fatal(err)
	}

	tokenResp, err := client.RequestAppAccessToken([]string{})
	if err != nil {
		log.Fatal(err)
	}

	// Set the access token on the client
	client.SetAppAccessToken(tokenResp.Data.AccessToken)

	return client
}
