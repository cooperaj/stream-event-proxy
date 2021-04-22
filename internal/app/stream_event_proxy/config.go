package stream_event_proxy

type Config struct {
	BroadcasterId string `required:"true"`
	ClientId      string `required:"true"`
	ClientSecret  string `required:"true"`
	Host          string `default:"0.0.0.0"`
	Port          int    `default:"5000"`
	ServiceUrl    string `required:"true"`
}
