package structs

type EmailSettings struct {
	Email         string `json:"email,omitempty"`
	EmailSettings struct {
		Promotions bool `json:"promotions,omitempty"`
		Messages   bool `json:"messages,omitempty"`
		NewMatches bool `json:"new_matches,omitempty"`
	} `json:"email_settings,omitempty"`
}
