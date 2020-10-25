package tinderapi

type LikedContentInfo struct {
	UserID      string `json:"user_id"`
	Type        string `json:"type"`
	Photo       Photo  `json:"photo"`
	IsSwipeNote bool   `json:"is_swipe_note"`
}
