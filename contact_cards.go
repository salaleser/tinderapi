package tinderapi

type ContactCards struct {
	PopulatedCards []struct {
		ContactType string `json:"contact_type,omitempty"`
		ContactID   string `json:"contact_id,omitempty"`
		Deeplink    string `json:"deeplink,omitempty"`
		Default     bool   `json:"default,omitempty"`
	} `json:"populated_cards,omitempty"`
	AvailableCards []string `json:"available_cards,omitempty"`
}
