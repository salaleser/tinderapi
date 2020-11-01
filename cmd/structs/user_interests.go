package structs

type UserInterests struct {
	SelectedInterests  []Interest `json:"selected_interests,omitempty"`
	AvailableInterests []Interest `json:"available_interests,omitempty"`
	MinInterests       int        `json:"min_interests,omitempty"`
	MaxInterests       int        `json:"max_interests,omitempty"`
}
