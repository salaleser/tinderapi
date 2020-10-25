package tinderapi

type SuperLikes struct {
	Remaining                    int         `json:"remaining"`
	AlcRemaining                 int         `json:"alc_remaining"`
	NewAlcRemaining              int         `json:"new_alc_remaining"`
	Allotment                    int         `json:"allotment"`
	SuperlikeRefreshAmount       int         `json:"superlike_refresh_amount"`
	SuperlikeRefreshInterval     int         `json:"superlike_refresh_interval"`
	SuperlikeRefreshIntervalUnit string      `json:"superlike_refresh_interval_unit"`
	ResetsAt                     interface{} `json:"resets_at"`
}
