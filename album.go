package tinderapi

type Album struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Images []struct {
		Height int    `json:"height"`
		Width  int    `json:"width"`
		URL    string `json:"url"`
	} `json:"images"`
}
