package tinderapi

type Photo struct {
	ID             string   `json:"id"`
	CropInfo       CropInfo `json:"crop_info,omitempty"`
	URL            string   `json:"url"`
	ProcessedFiles []struct {
		URL    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"processedFiles"`
	FileName        string `json:"fileName"`
	Extension       string `json:"extension"`
	ProcessedVideos []struct {
		URL    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"processedVideos,omitempty"`
}
