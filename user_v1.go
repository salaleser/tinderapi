package tinderapi

type UserV1 struct {
	Status  int  `json:"status"`
	Results User `json:"results"`
}
