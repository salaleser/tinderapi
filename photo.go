package tinderapi

type Photo struct {
	ID              string            `json:"id,omitempty"`
	CropInfo        CropInfo          `json:"crop_info,omitempty"`
	URL             string            `json:"url,omitempty"`
	ProcessedFiles  []ProcessedObject `json:"processedFiles,omitempty"`
	FileName        string            `json:"fileName,omitempty"`
	Extension       string            `json:"extension,omitempty"`
	ProcessedVideos []ProcessedObject `json:"processedVideos,omitempty"`
	WebpQf          []int             `json:"webp_qf,omitempty"`
}
