package stream_event_proxy

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

type conn struct {
	ws *websocket.Conn
}

func (c *conn) keepalive() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()

	for {
		select {
		case <-ticker.C:
			c.ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Printf("failed to write ping message to websocket. err: %v\n", err)
				return
			}
		}
	}
}

type EventPublisher struct {
	upgrader *websocket.Upgrader

	connections         map[*conn]bool
	connectionObservers []func()
	events              chan *event
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

	publisher.connections = make(map[*conn]bool)
	publisher.events = make(chan *event, 20)

	return publisher
}

func (e *EventPublisher) Run() {
	for {
		select {
		case event := <-e.events:
			var count = len(e.connections)
			for client := range e.connections {
				client.ws.SetWriteDeadline(time.Now().Add(writeWait))
				if err := client.ws.WriteJSON(event); err != nil {
					count--

					if _, ok := e.connections[client]; ok {
						delete(e.connections, client)
					}

					client.ws.Close()
					log.Printf("websocket failed to write, removing. err: %v\n", err)
				}
			}
			log.Printf("wrote %s event to %d client socket/s\n", event.Type, count)
		}
	}
}

func (e *EventPublisher) PublishEvent(event *event) {
	e.events <- event
}

func (e *EventPublisher) AddConnection(ws *websocket.Conn) {
	c := &conn{
		ws: ws,
	}

	e.connections[c] = true
	go c.keepalive()

	for _, observer := range e.connectionObservers {
		observer()
	}
	e.connectionObservers = nil
}

func (e *EventPublisher) AddConnectionObserver(observer func()) {
	e.connectionObservers = append(e.connectionObservers, observer)
}

func (e *EventPublisher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := e.upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	if err != nil {
		log.Println("client failed to negotiate websocket correctly")
		return
	}

	conn.SetReadLimit(512)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	e.AddConnection(conn)
	log.Println("client websocket connection created")
}
