package tinderapi

type Spotify struct {
	SpotifyConnected  bool          `json:"spotify_connected"`
	SpotifyTopArtists []interface{} `json:"spotify_top_artists"`
	SpotifyThemeTrack struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Album struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Images []struct {
				Height int    `json:"height"`
				Width  int    `json:"width"`
				URL    string `json:"url"`
			} `json:"images"`
		} `json:"album"`
		Artists []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"artists"`
		PreviewURL string `json:"preview_url"`
		URI        string `json:"uri"`
	} `json:"spotify_theme_track"`
}
