package tinderapi

type Seen struct {
	MatchSeen     bool   `json:"match_seen"`
	LastSeenMsgID string `json:"last_seen_msg_id"`
}
