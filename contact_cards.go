package tinderapi

type ContactCards struct {
	PopulatedCards []struct {
		ContactType string `json:"contact_type"`
		ContactID   string `json:"contact_id"`
		Deeplink    string `json:"deeplink,omitempty"`
		Default     bool   `json:"default,omitempty"`
	} `json:"populated_cards"`
	AvailableCards []string `json:"available_cards"`
}
