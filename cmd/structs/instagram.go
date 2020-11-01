package structs

import "time"

type Instagram struct {
	Username              string    `json:"username,omitempty"`
	LastFetchTime         time.Time `json:"last_fetch_time,omitempty"`
	CompletedInitialFetch bool      `json:"completed_initial_fetch,omitempty"`
	Photos                []struct {
		Image     string `json:"image,omitempty"`
		Thumbnail string `json:"thumbnail,omitempty"`
		Ts        string `json:"ts,omitempty"`
	} `json:"photos,omitempty"`
	MediaCount int `json:"media_count,omitempty"`
}
