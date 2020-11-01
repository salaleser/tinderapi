package structs

type Like struct {
	Status           int  `json:"status,omitempty"`
	Match            bool `json:"match,omitempty"`
	RateLimitedUntil int  `json:"rate_limited_until,omitempty"`
	LikesRemaining   int  `json:"likes_remaining,omitempty"`
	// XPadding       string `json:"X-Padding"`
}
