package structs

type LikedContent struct {
	ByOpener LikedContentInfo `json:"by_opener,omitempty"`
	ByCloser LikedContentInfo `json:"by_closer,omitempty"`
}
