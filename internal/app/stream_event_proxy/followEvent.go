package stream_event_proxy

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nicklaw5/helix"
)

type FollowEvent struct {
	client    *helix.Client
	publisher *EventPublisher

	subscription  *helix.EventSubSubscriptionsResponse
	broadcasterId string
	webhookUrl    string
	webhookSecret string
}

func NewFollowEvent(broadcasterId string, serviceUrl string, client *helix.Client, publisher *EventPublisher) *FollowEvent {
	var followEvent = &FollowEvent{
		client:        client,
		publisher:     publisher,
		broadcasterId: broadcasterId,
		webhookSecret: generateWebhookSecret(),
		webhookUrl:    serviceUrl + "/follow",
	}

	publisher.AddConnectionObserver(func() {
		followEvent.createEventSubscription()
	})

	return followEvent
}

func (f *FollowEvent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()

	// verify that the notification came from twitch using the secret.
	if !helix.VerifyEventSubNotification(f.webhookSecret, r.Header, string(body)) {
		log.Println("no valid signature on subscription")
		return
	} else {
		log.Println("verified signature for subscription")
	}

	var vals eventSubNotification
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&vals)
	if err != nil {
		log.Println(err)
		return
	}

	// if there's a challenge in the request, respond with only the challenge to verify your eventsub.
	if vals.Challenge != "" {
		w.Write([]byte(vals.Challenge))
		return
	}

	var followEvent helix.EventSubChannelFollowEvent
	s, _ := vals.Event.MarshalJSON()
	err = json.NewDecoder(bytes.NewReader(s)).Decode(&followEvent)

	f.publisher.PublishEvent(&event{
		Type: "follow",
		Data: &followEvent,
	})

	log.Printf("got follow webhook: %s follows %s\n", followEvent.UserName, followEvent.BroadcasterUserName)
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func (f *FollowEvent) createEventSubscription() {
	sub, err := f.client.CreateEventSubSubscription(&helix.EventSubSubscription{
		Type:    helix.EventSubTypeChannelFollow,
		Version: "1",
		Condition: helix.EventSubCondition{
			BroadcasterUserID: f.broadcasterId,
		},
		Transport: helix.EventSubTransport{
			Method:   "webhook",
			Callback: f.webhookUrl,
			Secret:   f.webhookSecret,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	f.subscription = sub
}
