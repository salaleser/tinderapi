package tinderapi

type Like struct {
	Status         int  `json:"status"`
	Match          bool `json:"match"`
	LikesRemaining int  `json:"likes_remaining"`
	// XPadding       string `json:"X-Padding"`
}
