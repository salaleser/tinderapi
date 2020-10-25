package tinderapi

import "time"

type Match struct {
	Seen               Seen         `json:"seen"`
	ID                 string       `json:"_id"`
	Closed             bool         `json:"closed"`
	CommonFriendCount  int          `json:"common_friend_count"`
	CommonLikeCount    int          `json:"common_like_count"`
	CreatedDate        time.Time    `json:"created_date"`
	Dead               bool         `json:"dead"`
	LastActivityDate   time.Time    `json:"last_activity_date"`
	MessageCount       int          `json:"message_count"`
	Messages           []Message    `json:"messages"`
	Muted              bool         `json:"muted"`
	Participants       []string     `json:"participants"`
	Pending            bool         `json:"pending"`
	IsSuperLike        bool         `json:"is_super_like"`
	IsBoostMatch       bool         `json:"is_boost_match"`
	IsSuperBoostMatch  bool         `json:"is_super_boost_match"`
	IsExperiencesMatch bool         `json:"is_experiences_match"`
	IsFastMatch        bool         `json:"is_fast_match"`
	IsOpener           bool         `json:"is_opener"`
	Person             User         `json:"person"`
	Following          bool         `json:"following"`
	FollowingMoments   bool         `json:"following_moments"`
	Readreceipt        Readreceipt  `json:"readreceipt"`
	LikedContent       LikedContent `json:"liked_content"`
}
