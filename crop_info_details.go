package tinderapi

type CropInfoDetails struct {
	WidthPct   float64 `json:"width_pct,omitempty"`
	XOffsetPct float64 `json:"x_offset_pct,omitempty"`
	HeightPct  float64 `json:"height_pct,omitempty"`
	YOffsetPct float64 `json:"y_offset_pct,omitempty"`
}
