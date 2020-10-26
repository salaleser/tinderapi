package tinderapi

type TopArtist struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Popularity int    `json:"popularity,omitempty"`
	Selected   bool   `json:"selected,omitempty"`
	TopTrack   Track  `json:"top_track,omitempty"`
}
