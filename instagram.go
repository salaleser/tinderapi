package tinderapi

import "time"

type Instagram struct {
	Username              string    `json:"username"`
	LastFetchTime         time.Time `json:"last_fetch_time"`
	CompletedInitialFetch bool      `json:"completed_initial_fetch"`
	Photos                []struct {
		Image     string `json:"image"`
		Thumbnail string `json:"thumbnail"`
		Ts        string `json:"ts"`
	} `json:"photos"`
	MediaCount int `json:"media_count"`
}
