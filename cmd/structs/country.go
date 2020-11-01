package structs

type Country struct {
	Name   string `json:"name,omitempty"`
	Cc     string `json:"cc,omitempty"`
	Alpha3 string `json:"alpha3,omitempty"`
}
