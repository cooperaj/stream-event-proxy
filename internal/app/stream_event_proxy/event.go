package stream_event_proxy

import (
	"encoding/json"
	"log"
	"math/rand"
	"strings"
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

	// Ensure that existing subscriptions for this client are removed before adding new
	cleanupSubscriptions(client, config.ServiceUrl, config.BroadcasterId)

	return client
}

func cleanupSubscriptions(client *helix.Client, serviceUrl string, broadcasterId string) {
	subs, err := client.GetEventSubSubscriptions(&helix.EventSubSubscriptionsParams{})
	if err != nil {
		log.Fatal(err)
	}

	var count int = 0
	for _, sub := range subs.Data.EventSubSubscriptions {
		if strings.HasPrefix(sub.Transport.Callback, serviceUrl) && sub.Condition.BroadcasterUserID == broadcasterId {
			client.RemoveEventSubSubscription(sub.ID)
			count++
		}
	}
	if count > 0 {
		log.Printf("removed %d prior subscriptions for user %s", count, broadcasterId)
	}
}
