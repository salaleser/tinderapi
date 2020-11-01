package structs

import "time"

type Message struct {
	ID        string    `json:"_id,omitempty"`
	MatchID   string    `json:"match_id,omitempty"`
	SentDate  time.Time `json:"sent_date,omitempty"`
	Message   string    `json:"message,omitempty"`
	To        string    `json:"to,omitempty"`
	From      string    `json:"from,omitempty"`
	Timestamp int       `json:"timestamp,omitempty"`
	Media     struct {
		Width  interface{} `json:"width,omitempty"`
		Height interface{} `json:"height,omitempty"`
	} `json:"media,omitempty"`
	CreatedDate time.Time `json:"created_date,omitempty"`
}
