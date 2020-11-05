package structs

type Page struct {
	Meta        *Meta  `json:"meta,omitempty"`
	Data        *Data  `json:"data,omitempty"`
	Error       *Error `json:"error,omitempty"`
	BatchErrors *[]struct {
		Error *Error  `json:"error,omitempty"`
		Key   *string `json:"key,omitempty"`
	} `json:"batch_errors,omitempty"`
}
