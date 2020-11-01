package structs

type UserV1 struct {
	Status  int  `json:"status,omitempty"`
	Results User `json:"results,omitempty"`
}
