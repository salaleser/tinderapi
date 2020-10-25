package tinderapi

type Likes struct {
	LikesRemaining   int   `json:"likes_remaining"`
	RateLimitedUntil int64 `json:"rate_limited_until"`
}
