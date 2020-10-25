package tinderapi

type Job struct {
	Company struct {
		Name      string `json:"name"`
		Displayed bool   `json:"displayed"`
	} `json:"company"`
	Title struct {
		Displayed bool   `json:"displayed"`
		Name      string `json:"name"`
	} `json:"title"`
}
