package tinderapi

type Data struct {
	Account Account  `json:"account"`
	User    User     `json:"user"`
	Results []Result `json:"results"`
}
