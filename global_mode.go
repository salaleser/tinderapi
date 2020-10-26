package tinderapi

type GlobalMode struct {
	IsEnabled           bool   `json:"is_enabled,omitempty"`
	DisplayLanguage     string `json:"display_language,omitempty"`
	LanguagePreferences []struct {
		Language   string `json:"language,omitempty"`
		IsSelected bool   `json:"is_selected,omitempty"`
	} `json:"language_preferences,omitempty"`
}
