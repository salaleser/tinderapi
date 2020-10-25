package tinderapi

type GlobalMode struct {
	IsEnabled           bool   `json:"is_enabled"`
	DisplayLanguage     string `json:"display_language"`
	LanguagePreferences []struct {
		Language   string `json:"language"`
		IsSelected bool   `json:"is_selected"`
	} `json:"language_preferences"`
}
