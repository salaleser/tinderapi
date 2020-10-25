package tinderapi

type Track struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Album      Album    `json:"album"`
	Artists    []Artist `json:"artists"`
	PreviewURL string   `json:"preview_url"`
	URI        string   `json:"uri"`
}
