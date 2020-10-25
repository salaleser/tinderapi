package tinderapi

type CropInfo struct {
	User                CropInfoDetails `json:"user"`
	Algo                CropInfoDetails `json:"algo"`
	ProcessedByBullseye bool            `json:"processed_by_bullseye"`
	UserCustomized      bool            `json:"user_customized"`
}
