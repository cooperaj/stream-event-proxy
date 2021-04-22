package stream_event_proxy

import (
	"encoding/json"
	"net/http"

	"github.com/nicklaw5/helix"
)

type Follows struct {
	client        *helix.Client
	broadcasterId string
}

type followsResponse struct {
	Total   int                `json:"total"`
	Follows []helix.UserFollow `json:"follows"`
}

func NewFollows(broadcasterId string, client *helix.Client) *Follows {
	return &Follows{
		client:        client,
		broadcasterId: broadcasterId,
	}
}

func (f *Follows) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	follows, err := f.client.GetUsersFollows(&helix.UsersFollowsParams{
		ToID: f.broadcasterId,
	})
	if err != nil {
		http.Error(w, "Twitch API error", http.StatusInternalServerError)
	}

	data := &followsResponse{
		Total:   follows.Data.Total,
		Follows: follows.Data.Follows,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
