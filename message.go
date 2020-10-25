package tinderapi

import "time"

type Message struct {
	ID        string    `json:"_id"`
	MatchID   string    `json:"match_id"`
	SentDate  time.Time `json:"sent_date"`
	Message   string    `json:"message"`
	To        string    `json:"to"`
	From      string    `json:"from"`
	Timestamp int       `json:"timestamp"`
}
