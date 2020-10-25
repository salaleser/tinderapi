package tinderapi

type Data struct {
	Account       Account       `json:"account"`
	Results       []User        `json:"results"`
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
	Swipenote     struct {
		Remaining int `json:"remaining"`
	} `json:"swipenote"`
	Spotify    `json:"spotify"`
	SuperLikes SuperLikes `json:"super_likes"`
	TinderU    struct {
		Status string `json:"status"`
	} `json:"tinder_u"`
	Travel struct {
		IsTraveling bool `json:"is_traveling"`
	} `json:"travel"`
	Tutorials []string `json:"tutorials"`
}
