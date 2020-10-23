package tinderapi

import "time"

type UserV1 struct {
	Status  int `json:"status"`
	Results struct {
		CommonFriends     []interface{} `json:"common_friends"`
		CommonFriendCount int           `json:"common_friend_count"`
		DistanceMi        int           `json:"distance_mi"`
		ConnectionCount   int           `json:"connection_count"`
		CommonConnections []interface{} `json:"common_connections"`
		Bio               string        `json:"bio"`
		BirthDate         time.Time     `json:"birth_date"`
		Name              string        `json:"name"`
		Instagram         struct {
			LastFetchTime         time.Time `json:"last_fetch_time"`
			CompletedInitialFetch bool      `json:"completed_initial_fetch"`
			Photos                []struct {
				Image     string `json:"image"`
				Thumbnail string `json:"thumbnail"`
				Ts        string `json:"ts"`
			} `json:"photos"`
			MediaCount int    `json:"media_count"`
			Username   string `json:"username"`
		} `json:"instagram"`
		Jobs    []interface{} `json:"jobs"`
		Schools []interface{} `json:"schools"`
		Teasers []struct {
			Type   string `json:"type"`
			String string `json:"string"`
		} `json:"teasers"`
		Gender        int           `json:"gender"`
		BirthDateInfo string        `json:"birth_date_info"`
		PingTime      time.Time     `json:"ping_time"`
		Badges        []interface{} `json:"badges"`
		Photos        []struct {
			ID             string   `json:"id"`
			CropInfo       CropInfo `json:"crop_info,omitempty"`
			URL            string   `json:"url"`
			ProcessedFiles []struct {
				URL    string `json:"url"`
				Height int    `json:"height"`
				Width  int    `json:"width"`
			} `json:"processedFiles"`
			FileName  string `json:"fileName"`
			Extension string `json:"extension"`
		} `json:"photos"`
		CommonLikes     []interface{} `json:"common_likes"`
		CommonLikeCount int           `json:"common_like_count"`
		CommonInterests []interface{} `json:"common_interests"`
		SNumber         int           `json:"s_number"`
		ID              string        `json:"_id"`
		IsTinderU       bool          `json:"is_tinder_u"`
	} `json:"results"`
}
