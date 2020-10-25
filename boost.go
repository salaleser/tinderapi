package tinderapi

type Boost struct {
	Duration                     int           `json:"duration"`
	Allotment                    int           `json:"allotment"`
	AllotmentUsed                int           `json:"allotment_used"`
	AllotmentRemaining           int           `json:"allotment_remaining"`
	InternalRemaining            int           `json:"internal_remaining"`
	PurchasedRemaining           int           `json:"purchased_remaining"`
	Remaining                    int           `json:"remaining"`
	SuperBoostPurchasedRemaining int           `json:"super_boost_purchased_remaining"`
	SuperBoostRemaining          int           `json:"super_boost_remaining"`
	BoostRefreshAmount           int           `json:"boost_refresh_amount"`
	BoostRefreshInterval         int           `json:"boost_refresh_interval"`
	BoostRefreshIntervalUnit     string        `json:"boost_refresh_interval_unit"`
	Purchases                    []interface{} `json:"purchases"`
}
