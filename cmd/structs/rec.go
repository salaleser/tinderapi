package structs

type Rec struct {
	Type           string    `json:"type,omitempty"`
	User           User      `json:"user,omitempty"`
	Facebook       Facebook  `json:"facebook,omitempty"`
	Spotify        Spotify   `json:"spotify,omitempty"`
	Teaser         Teaser    `json:"teaser,omitempty"`
	Instagram      Instagram `json:"instagram,omitempty"`
	DistanceMi     int       `json:"distance_mi,omitempty"`
	ContentHash    string    `json:"content_hash,omitempty"`
	SNumber        int       `json:"s_number,omitempty"`
	Teasers        []Teaser  `json:"teasers,omitempty"`
	ExperimentInfo struct {
		UserInterests struct {
			SelectedInterests []struct {
				ID   string `json:"id,omitempty"`
				Name string `json:"name,omitempty"`
			} `json:"selected_interests,omitempty"`
		} `json:"user_interests,omitempty"`
	} `json:"experiment_info,omitempty"`
	IsSuperlikeUpsell bool `json:"is_superlike_upsell,omitempty"`
}
