package tinderapi

type CropInfoDetails struct {
	WidthPct   float64 `json:"width_pct"`
	XOffsetPct float64 `json:"x_offset_pct"`
	HeightPct  float64 `json:"height_pct"`
	YOffsetPct float64 `json:"y_offset_pct"`
}
