package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"

	sep "go.acpr.dev/stream-event-proxy/internal/app/stream_event_proxy"
)

func main() {
	var c sep.Config
	err := envconfig.Process("sep", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	client := sep.NewClient(&c)
	publisher := sep.NewEventPublisher()
	go publisher.Run()

	router.Handle("/ws", publisher)

	router.Handle("/follow", sep.NewFollowEvent(c.BroadcasterId, c.ServiceUrl, client, publisher)).Methods("POST")
	router.Handle("/follows", sep.NewFollows(c.BroadcasterId, client)).Methods("GET")

	spa := &sep.Spa{StaticPath: "./web/assets"}
	router.PathPrefix("/").Handler(spa)

	handler := handlers.RecoveryHandler()(handlers.CombinedLoggingHandler(os.Stdout, router))

	srv := &http.Server{
		Handler:      handler,
		Addr:         c.Host + ":" + strconv.Itoa(c.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
	log.Println("stream event proxy started")
}
