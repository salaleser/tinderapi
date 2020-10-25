package tinderapi

import "time"

type User struct {
	Type                        string                   `json:"type"`
	Instagram                   Instagram                `json:"instagram,omitempty"`
	Facebook                    Facebook                 `json:"facebook"`
	Spotify                     Spotify                  `json:"spotify,omitempty"`
	DistanceMi                  int                      `json:"distance_mi"`
	ContentHash                 string                   `json:"content_hash"`
	SNumber                     int                      `json:"s_number"`
	Teaser                      Teaser                   `json:"teaser,omitempty"`
	Teasers                     []Teaser                 `json:"teasers"`
	ExperimentInfo              ExperimentInfo           `json:"experiment_info,omitempty"`
	CommonFriends               []interface{}            `json:"common_friends"`
	CommonFriendCount           int                      `json:"common_friend_count"`
	ConnectionCount             int                      `json:"connection_count"`
	CommonConnections           []interface{}            `json:"common_connections"`
	BirthDateInfo               string                   `json:"birth_date_info"`
	CommonLikes                 []interface{}            `json:"common_likes"`
	CommonLikeCount             int                      `json:"common_like_count"`
	CommonInterests             []interface{}            `json:"common_interests"`
	IsTinderU                   bool                     `json:"is_tinder_u"`
	IsTraveling                 bool                     `json:"is_traveling"`
	ShowGenderOnProfile         bool                     `json:"show_gender_on_profile"`
	HideAge                     bool                     `json:"hide_age"`
	HideDistance                bool                     `json:"hide_distance"`
	AgeFilterMax                int                      `json:"age_filter_max"`
	AgeFilterMin                int                      `json:"age_filter_min"`
	Bio                         string                   `json:"bio"`
	BirthDate                   time.Time                `json:"birth_date"`
	CreateDate                  time.Time                `json:"create_date"`
	CrmID                       string                   `json:"crm_id"`
	ID                          string                   `json:"_id"`
	PosInfo                     PosInfo                  `json:"pos_info"`
	Discoverable                bool                     `json:"discoverable"`
	DistanceFilter              int                      `json:"distance_filter"`
	GlobalMode                  GlobalMode               `json:"global_mode"`
	Gender                      int                      `json:"gender"`
	GenderFilter                int                      `json:"gender_filter"`
	Name                        string                   `json:"name"`
	Photos                      []Photo                  `json:"photos"`
	PhotosProcessing            bool                     `json:"photos_processing"`
	PhotoOptimizerEnabled       bool                     `json:"photo_optimizer_enabled"`
	PingTime                    time.Time                `json:"ping_time"`
	RecentlyActive              bool                     `json:"recently_active"`
	Jobs                        []Job                    `json:"jobs"`
	Schools                     []School                 `json:"schools"`
	Badges                      []Badge                  `json:"badges"`
	Username                    string                   `json:"username"`
	PhoneID                     string                   `json:"phone_id"`
	InterestedIn                []int                    `json:"interested_in"`
	Pos                         Pos                      `json:"pos"`
	AutoplayVideo               string                   `json:"autoplay_video"`
	TopPicksDiscoverable        bool                     `json:"top_picks_discoverable"`
	PhotoTaggingEnabled         bool                     `json:"photo_tagging_enabled"`
	City                        City                     `json:"city"`
	ShowOrientationOnProfile    bool                     `json:"show_orientation_on_profile"`
	ShowSameOrientationFirst    ShowSameOrientationFirst `json:"show_same_orientation_first"`
	SexualOrientations          []SexualOrientation      `json:"sexual_orientations"`
	UserInterests               UserInterests            `json:"user_interests"`
	RecommendedSortDiscoverable bool                     `json:"recommended_sort_discoverable"`
	SelfieVerification          string                   `json:"selfie_verification"`
	NoonlightProtected          bool                     `json:"noonlight_protected"`
	UserPresenceDisabled        bool                     `json:"user_presence_disabled"`
	SyncSwipeEnabled            bool                     `json:"sync_swipe_enabled"`
}
