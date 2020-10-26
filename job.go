package tinderapi

type Job struct {
	Company struct {
		Name      string `json:"name,omitempty"`
		Displayed bool   `json:"displayed",omitempty`
	} `json:"company,omitempty"`
	Title struct {
		Displayed bool   `json:"displayed,omitempty"`
		Name      string `json:"name,omitempty"`
	} `json:"title,omitempty"`
}
