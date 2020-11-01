package structs

type Boost struct {
	Duration                     int           `json:"duration,omitempty"`
	Allotment                    int           `json:"allotment,omitempty"`
	AllotmentUsed                int           `json:"allotment_used,omitempty"`
	AllotmentRemaining           int           `json:"allotment_remaining,omitempty"`
	InternalRemaining            int           `json:"internal_remaining,omitempty"`
	PurchasedRemaining           int           `json:"purchased_remaining,omitempty"`
	Remaining                    int           `json:"remaining,omitempty"`
	SuperBoostPurchasedRemaining int           `json:"super_boost_purchased_remaining,omitempty"`
	SuperBoostRemaining          int           `json:"super_boost_remaining,omitempty"`
	BoostRefreshAmount           int           `json:"boost_refresh_amount,omitempty"`
	BoostRefreshInterval         int           `json:"boost_refresh_interval,omitempty"`
	BoostRefreshIntervalUnit     string        `json:"boost_refresh_interval_unit,omitempty"`
	Purchases                    []interface{} `json:"purchases,omitempty"`
}
