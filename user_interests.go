package tinderapi

type UserInterests struct {
	SelectedInterests  []Interest `json:"selected_interests"`
	AvailableInterests []Interest `json:"available_interests"`
	MinInterests       int        `json:"min_interests"`
	MaxInterests       int        `json:"max_interests"`
}
