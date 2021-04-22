package stream_event_proxy

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type EventPublisher struct {
	upgrader *websocket.Upgrader

	connections map[*websocket.Conn]bool
	events      chan *event
}

func NewEventPublisher() *EventPublisher {
	publisher := &EventPublisher{
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	publisher.connections = make(map[*websocket.Conn]bool)
	publisher.events = make(chan *event, 20)

	return publisher
}

func (e *EventPublisher) Run() {
	for {
		select {
		case event := <-e.events:
			for client := range e.connections {
				if err := client.WriteJSON(event); err != nil {
					log.Println("Websocket failed to write")

					if _, ok := e.connections[client]; ok {
						delete(e.connections, client)
					}

					client.Close()
				}
			}
			log.Println("wrote \"" + event.Type + "\" event to clients sockets")
		}
	}
}

func (e *EventPublisher) PublishEvent(event *event) {
	e.events <- event
}

func (e *EventPublisher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := e.upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	if err != nil {
		log.Println("client failed to negotiate websocket correctly")
		return
	}

	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(60 * time.Second)); return nil })

	e.connections[conn] = true
	log.Println("client websocket connection created")
}
