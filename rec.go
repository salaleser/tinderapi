package tinderapi

type Rec struct {
	Type      string    `json:"type"`
	User      User      `json:"user,omitempty"`
	Facebook  Facebook  `json:"facebook,omitempty"`
	Spotify   Spotify   `json:"spotify,omitempty"`
	Teaser    Teaser    `json:"teaser,omitempty"`
	Instagram Instagram `json:"instagram,omitempty"`
}
