package structs

type Album struct {
	ID     string  `json:"id,omitempty"`
	Name   string  `json:"name,omitempty"`
	Images []Image `json:"images,omitempty"`
}
