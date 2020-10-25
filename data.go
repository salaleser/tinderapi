package tinderapi

type Data struct {
	Account       Account       `json:"account"`
	Results       []Rec         `json:"results"`
	Matches       []Match       `json:"matches"`
	User          User          `json:"user"`
	Boost         Boost         `json:"boost"`
	ContactCards  ContactCards  `json:"contact_cards"`
	EmailSettings EmailSettings `json:"email_settings"`
	Instagram     Instagram     `json:"instagram"`
	Likes         Likes         `json:"likes"`
	Notifications []interface{} `json:"notifications"`
	PlusControl   struct{}      `json:"plus_control"`
	Products      struct{}      `json:"products"`
	Purchase      Purchase      `json:"purchase"`
	Readreceipts  Readreceipts  `json:"readreceipts"`
	Swipenote     Swipenote     `json:"swipenote"`
	Spotify       `json:"spotify"`
	SuperLikes    SuperLikes `json:"super_likes"`
	TinderU       TinderU    `json:"tinder_u"`
	Travel        Travel     `json:"travel"`
	Tutorials     []string   `json:"tutorials"`
	NextPageToken string     `json:"next_page_token"`
}
