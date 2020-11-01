package structs

import "time"

type User struct {
	Type                        string                   `json:"type,omitempty"`
	Instagram                   Instagram                `json:"instagram,omitempty"`
	Facebook                    Facebook                 `json:"facebook,omitempty"`
	Spotify                     Spotify                  `json:"spotify,omitempty"`
	DistanceMi                  int                      `json:"distance_mi,omitempty"`
	ContentHash                 string                   `json:"content_hash,omitempty"`
	SNumber                     int                      `json:"s_number,omitempty"`
	Teaser                      Teaser                   `json:"teaser,omitempty"`
	Teasers                     []Teaser                 `json:"teasers,omitempty"`
	ExperimentInfo              ExperimentInfo           `json:"experiment_info,omitempty"`
	CommonFriends               []interface{}            `json:"common_friends,omitempty"`
	CommonFriendCount           int                      `json:"common_friend_count,omitempty"`
	ConnectionCount             int                      `json:"connection_count,omitempty"`
	CommonConnections           []interface{}            `json:"common_connections,omitempty"`
	BirthDateInfo               string                   `json:"birth_date_info,omitempty"`
	CommonLikes                 []interface{}            `json:"common_likes,omitempty"`
	CommonLikeCount             int                      `json:"common_like_count,omitempty"`
	CommonInterests             []interface{}            `json:"common_interests,omitempty"`
	IsTinderU                   bool                     `json:"is_tinder_u,omitempty"`
	IsTraveling                 bool                     `json:"is_traveling,omitempty"`
	ShowGenderOnProfile         bool                     `json:"show_gender_on_profile,omitempty"`
	HideAge                     bool                     `json:"hide_age,omitempty"`
	HideDistance                bool                     `json:"hide_distance,omitempty"`
	AgeFilterMax                int                      `json:"age_filter_max,omitempty"`
	AgeFilterMin                int                      `json:"age_filter_min,omitempty"`
	Bio                         string                   `json:"bio,omitempty"`
	BirthDate                   time.Time                `json:"birth_date,omitempty"`
	CreateDate                  time.Time                `json:"create_date,omitempty"`
	CrmID                       string                   `json:"crm_id,omitempty"`
	ID                          string                   `json:"_id,omitempty"`
	PosInfo                     PosInfo                  `json:"pos_info,omitempty"`
	Discoverable                bool                     `json:"discoverable,omitempty"`
	DistanceFilter              int                      `json:"distance_filter,omitempty"`
	GlobalMode                  GlobalMode               `json:"global_mode,omitempty"`
	Gender                      int                      `json:"gender,omitempty"`
	GenderFilter                int                      `json:"gender_filter,omitempty"`
	Name                        string                   `json:"name,omitempty"`
	Photos                      []Photo                  `json:"photos,omitempty"`
	PhotosProcessing            bool                     `json:"photos_processing,omitempty"`
	PhotoOptimizerEnabled       bool                     `json:"photo_optimizer_enabled,omitempty"`
	PingTime                    time.Time                `json:"ping_time,omitempty"`
	RecentlyActive              bool                     `json:"recently_active,omitempty"`
	Jobs                        []Job                    `json:"jobs,omitempty"`
	Schools                     []School                 `json:"schools,omitempty"`
	Badges                      []Badge                  `json:"badges,omitempty"`
	Username                    string                   `json:"username,omitempty"`
	PhoneID                     string                   `json:"phone_id,omitempty"`
	InterestedIn                []int                    `json:"interested_in,omitempty"`
	Pos                         Pos                      `json:"pos,omitempty"`
	AutoplayVideo               string                   `json:"autoplay_video,omitempty"`
	TopPicksDiscoverable        bool                     `json:"top_picks_discoverable,omitempty"`
	PhotoTaggingEnabled         bool                     `json:"photo_tagging_enabled,omitempty"`
	City                        City                     `json:"city,omitempty"`
	ShowOrientationOnProfile    bool                     `json:"show_orientation_on_profile,omitempty"`
	ShowSameOrientationFirst    ShowSameOrientationFirst `json:"show_same_orientation_first,omitempty"`
	SexualOrientations          []SexualOrientation      `json:"sexual_orientations,omitempty"`
	UserInterests               UserInterests            `json:"user_interests,omitempty"`
	RecommendedSortDiscoverable bool                     `json:"recommended_sort_discoverable,omitempty"`
	SelfieVerification          string                   `json:"selfie_verification,omitempty"`
	NoonlightProtected          bool                     `json:"noonlight_protected,omitempty"`
	UserPresenceDisabled        bool                     `json:"user_presence_disabled,omitempty"`
	SyncSwipeEnabled            bool                     `json:"sync_swipe_enabled,omitempty"`
}
