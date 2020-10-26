package tinderapi

import "time"

type Spotify struct {
	SpotifyConnected     bool        `json:"spotify_connected,omitempty"`
	SpotifyThemeTrack    Track       `json:"spotify_theme_track,omitempty"`
	SpotifyUsername      string      `json:"spotify_username,omitempty"`
	SpotifyUserType      string      `json:"spotify_user_type,omitempty"`
	SpotifyLastUpdatedAt time.Time   `json:"spotify_last_updated_at,omitempty"`
	SpotifyTopArtists    []TopArtist `json:"spotify_top_artists,omitempty"`
}
