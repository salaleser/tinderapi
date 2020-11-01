package structs

type ExperimentInfo struct {
	UserInterests struct {
		SelectedInterests []struct {
			ID   string `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"selected_interests,omitempty"`
	} `json:"user_interests,omitempty"`
}
