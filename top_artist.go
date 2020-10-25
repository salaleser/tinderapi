package tinderapi

type TopArtist struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Popularity int    `json:"popularity"`
	Selected   bool   `json:"selected"`
	TopTrack   Track  `json:"top_track"`
}
