package tinderapi

type Facebook struct {
	CommonConnections []interface{} `json:"common_connections"`
	ConnectionCount   int           `json:"connection_count"`
	CommonInterests   []interface{} `json:"common_interests"`
}
