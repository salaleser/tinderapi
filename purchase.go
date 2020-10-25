package tinderapi

type Purchase struct {
	Purchases           []interface{} `json:"purchases"`
	SubscriptionExpired bool          `json:"subscription_expired"`
}
