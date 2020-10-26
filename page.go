package tinderapi

type Page struct {
	Meta  Meta  `json:"meta,omitempty"`
	Data  Data  `json:"data,omitempty"`
	Error Error `json:"error,omitempty"`
}
