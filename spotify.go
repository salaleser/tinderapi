package tinderapi

import "time"

type Spotify struct {
	SpotifyConnected     bool        `json:"spotify_connected"`
	SpotifyThemeTrack    Track       `json:"spotify_theme_track"`
	SpotifyUsername      string      `json:"spotify_username"`
	SpotifyUserType      string      `json:"spotify_user_type"`
	SpotifyLastUpdatedAt time.Time   `json:"spotify_last_updated_at"`
	SpotifyTopArtists    []TopArtist `json:"spotify_top_artists"`
}
