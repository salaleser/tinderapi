package tinderapi

type EmailSettings struct {
	Email         string `json:"email"`
	EmailSettings struct {
		Promotions bool `json:"promotions"`
		Messages   bool `json:"messages"`
		NewMatches bool `json:"new_matches"`
	} `json:"email_settings"`
}
