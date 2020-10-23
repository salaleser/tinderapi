package tinderapi

type Result struct {
	Type           string         `json:"type"`
	User           User           `json:"user,omitempty"`
	Instagram      Instagram      `json:"instagram,omitempty"`
	Facebook       Facebook       `json:"facebook"`
	Spotify        Spotify        `json:"spotify,omitempty"`
	DistanceMi     int            `json:"distance_mi"`
	ContentHash    string         `json:"content_hash"`
	SNumber        int            `json:"s_number"`
	Teaser         Teaser         `json:"teaser,omitempty"`
	Teasers        []Teaser       `json:"teasers"`
	ExperimentInfo ExperimentInfo `json:"experiment_info,omitempty"`
}
