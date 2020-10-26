package tinderapi

import "time"

type Message struct {
	ID        string    `json:"_id,omitempty"`
	MatchID   string    `json:"match_id,omitempty"`
	SentDate  time.Time `json:"sent_date,omitempty"`
	Message   string    `json:"message,omitempty"`
	To        string    `json:"to,omitempty"`
	From      string    `json:"from,omitempty"`
	Timestamp int       `json:"timestamp,omitempty"`
}
