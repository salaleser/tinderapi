package tinderapi

type LikedContentInfo struct {
	UserID      string `json:"user_id,omitempty"`
	Type        string `json:"type,omitempty"`
	Photo       Photo  `json:"photo,omitempty"`
	IsSwipeNote bool   `json:"is_swipe_note,omitempty"`
}
