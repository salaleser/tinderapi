package tinderapi

import "time"

type Spotify struct {
	SpotifyConnected     bool      `json:"spotify_connected"`
	SpotifyThemeTrack    Track     `json:"spotify_theme_track"`
	SpotifyUsername      string    `json:"spotify_username"`
	SpotifyUserType      string    `json:"spotify_user_type"`
	SpotifyLastUpdatedAt time.Time `json:"spotify_last_updated_at"`
	SpotifyTopArtists    []struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Popularity int    `json:"popularity"`
		Selected   bool   `json:"selected"`
		TopTrack   Track  `json:"top_track"`
	} `json:"spotify_top_artists"`
}
