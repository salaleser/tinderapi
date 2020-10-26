package tinderapi

type Seen struct {
	MatchSeen     bool   `json:"match_seen,omitempty"`
	LastSeenMsgID string `json:"last_seen_msg_id,omitempty"`
}
