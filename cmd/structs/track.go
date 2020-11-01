package structs

type Track struct {
	ID         string   `json:"id,omitempty"`
	Name       string   `json:"name,omitempty"`
	Album      Album    `json:"album,omitempty"`
	Artists    []Artist `json:"artists,omitempty"`
	PreviewURL string   `json:"preview_url,omitempty"`
	URI        string   `json:"uri,omitempty"`
}
