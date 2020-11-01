package structs

type SuperLikes struct {
	Remaining                    int         `json:"remaining,omitempty"`
	AlcRemaining                 int         `json:"alc_remaining,omitempty"`
	NewAlcRemaining              int         `json:"new_alc_remaining,omitempty"`
	Allotment                    int         `json:"allotment,omitempty"`
	SuperlikeRefreshAmount       int         `json:"superlike_refresh_amount,omitempty"`
	SuperlikeRefreshInterval     int         `json:"superlike_refresh_interval,omitempty"`
	SuperlikeRefreshIntervalUnit string      `json:"superlike_refresh_interval_unit,omitempty"`
	ResetsAt                     interface{} `json:"resets_at,omitempty"`
}
