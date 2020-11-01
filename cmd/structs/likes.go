package structs

type Likes struct {
	LikesRemaining   int   `json:"likes_remaining,omitempty"`
	RateLimitedUntil int64 `json:"rate_limited_until,omitempty"`
}
