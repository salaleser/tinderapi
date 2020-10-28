package tinderapi

import "time"

type SendMessage struct {
	ID       string    `json:"_id"`
	From     string    `json:"from"`
	To       string    `json:"to"`
	MatchID  string    `json:"match_id"`
	SentDate time.Time `json:"sent_date"`
	Message  string    `json:"message"`
	Media    struct {
		Width  interface{} `json:"width"`
		Height interface{} `json:"height"`
	} `json:"media"`
	CreatedDate time.Time `json:"created_date"`
}
