package tinderapi

type LikedContent struct {
	ByOpener LikedContentInfo `json:"by_opener"`
	ByCloser LikedContentInfo `json:"by_closer"`
}
