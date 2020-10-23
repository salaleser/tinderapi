package tinderapi

type ExperimentInfo struct {
	UserInterests struct {
		SelectedInterests []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"selected_interests"`
	} `json:"user_interests"`
}
