package tinderapi

import "time"

type Match struct {
	Seen               Seen         `json:"seen,omitempty"`
	ID                 string       `json:"_id,omitempty"`
	Closed             bool         `json:"closed,omitempty"`
	CommonFriendCount  int          `json:"common_friend_count,omitempty"`
	CommonLikeCount    int          `json:"common_like_count,omitempty"`
	CreatedDate        time.Time    `json:"created_date,omitempty"`
	Dead               bool         `json:"dead,omitempty"`
	LastActivityDate   time.Time    `json:"last_activity_date,omitempty"`
	MessageCount       int          `json:"message_count,omitempty"`
	Messages           []Message    `json:"messages,omitempty"`
	Muted              bool         `json:"muted,omitempty"`
	Participants       []string     `json:"participants,omitempty"`
	Pending            bool         `json:"pending,omitempty"`
	IsSuperLike        bool         `json:"is_super_like,omitempty"`
	IsBoostMatch       bool         `json:"is_boost_match,omitempty"`
	IsSuperBoostMatch  bool         `json:"is_super_boost_match,omitempty"`
	IsExperiencesMatch bool         `json:"is_experiences_match,omitempty"`
	IsFastMatch        bool         `json:"is_fast_match,omitempty"`
	IsOpener           bool         `json:"is_opener,omitempty"`
	Person             User         `json:"person,omitempty"`
	Following          bool         `json:"following,omitempty"`
	FollowingMoments   bool         `json:"following_moments,omitempty"`
	Readreceipt        Readreceipt  `json:"readreceipt,omitempty"`
	LikedContent       LikedContent `json:"liked_content,omitempty"`
}
