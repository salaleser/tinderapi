package structs

type PosInfo struct {
	Country  Country `json:"country,omitempty"`
	Timezone string  `json:"timezone,omitempty"`
}
