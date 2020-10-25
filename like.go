package tinderapi

type Like struct {
	Status           int  `json:"status"`
	Match            bool `json:"match"`
	RateLimitedUntil int  `json:"rate_limited_until"`
	LikesRemaining   int  `json:"likes_remaining"`
	// XPadding       string `json:"X-Padding"`
}
