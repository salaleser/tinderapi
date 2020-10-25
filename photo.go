package tinderapi

type Photo struct {
	ID              string            `json:"id"`
	CropInfo        CropInfo          `json:"crop_info,omitempty"`
	URL             string            `json:"url"`
	ProcessedFiles  []ProcessedObject `json:"processedFiles"`
	FileName        string            `json:"fileName"`
	Extension       string            `json:"extension"`
	ProcessedVideos []ProcessedObject `json:"processedVideos,omitempty"`
	WebpQf          []int             `json:"webp_qf"`
}
