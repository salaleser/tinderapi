package structs

type Facebook struct {
	CommonConnections []interface{} `json:"common_connections,omitempty"`
	ConnectionCount   int           `json:"connection_count,omitempty"`
	CommonInterests   []interface{} `json:"common_interests,omitempty"`
}
