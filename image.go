package tinderapi

type Image struct {
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
	URL    string `json:"url,omitempty"`
}
