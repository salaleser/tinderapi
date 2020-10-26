package tinderapi

type CropInfo struct {
	User                CropInfoDetails `json:"user,omitempty"`
	Algo                CropInfoDetails `json:"algo,omitempty"`
	ProcessedByBullseye bool            `json:"processed_by_bullseye,omitempty"`
	UserCustomized      bool            `json:"user_customized,omitempty"`
}
