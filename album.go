package tinderapi

type Album struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Images []Image `json:"images"`
}
