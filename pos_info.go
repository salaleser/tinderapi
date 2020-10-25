package tinderapi

type PosInfo struct {
	Country  Country `json:"country"`
	Timezone string  `json:"timezone"`
}
