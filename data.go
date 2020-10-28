package tinderapi

type Data struct {
	Account       Account       `json:"account,omitempty"`
	Results       []Rec         `json:"results,omitempty"`
	Matches       []Match       `json:"matches,omitempty"`
	Messages      []Message     `json:"messages,omitempty"`
	User          User          `json:"user,omitempty"`
	Boost         Boost         `json:"boost,omitempty"`
	ContactCards  ContactCards  `json:"contact_cards,omitempty"`
	EmailSettings EmailSettings `json:"email_settings,omitempty"`
	Instagram     Instagram     `json:"instagram,omitempty"`
	Likes         Likes         `json:"likes,omitempty"`
	Notifications []interface{} `json:"notifications,omitempty"`
	PlusControl   struct{}      `json:"plus_control,omitempty"`
	Products      struct{}      `json:"products,omitempty"`
	Purchase      Purchase      `json:"purchase,omitempty"`
	Readreceipts  Readreceipts  `json:"readreceipts,omitempty"`
	Swipenote     Swipenote     `json:"swipenote,omitempty"`
	Spotify       Spotify       `json:"spotify,omitempty"`
	SuperLikes    SuperLikes    `json:"super_likes,omitempty"`
	TinderU       TinderU       `json:"tinder_u,omitempty"`
	Travel        Travel        `json:"travel,omitempty"`
	Tutorials     []string      `json:"tutorials,omitempty"`
	NextPageToken string        `json:"next_page_token,omitempty"`
}
