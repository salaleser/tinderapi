package tinderapi

type School struct {
	Name      string `json:"name,omitempty"`
	Displayed bool   `json:"displayed,omitempty"`
}
