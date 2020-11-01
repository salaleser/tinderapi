package structs

type Readreceipts struct {
	InternalRemaining  int `json:"internal_remaining,omitempty"`
	PurchasedRemaining int `json:"purchased_remaining,omitempty"`
	Remaining          int `json:"remaining,omitempty"`
}
