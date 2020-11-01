package structs

type Purchase struct {
	Purchases           []interface{} `json:"purchases,omitempty"`
	SubscriptionExpired bool          `json:"subscription_expired,omitempty"`
}
