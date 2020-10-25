package tinderapi

type Page struct {
	Meta  Meta  `json:"meta"`
	Data  Data  `json:"data"`
	Error Error `json:"error"`
}
