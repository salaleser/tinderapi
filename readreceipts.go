package tinderapi

type Readreceipts struct {
	InternalRemaining  int `json:"internal_remaining"`
	PurchasedRemaining int `json:"purchased_remaining"`
	Remaining          int `json:"remaining"`
}
