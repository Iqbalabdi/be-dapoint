package response

type ApiResponse struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type ApiResponseSuccess struct {
	Status string      `json:"status"`
	Count  int         `default:"1" json:"count"`
	Data   interface{} `json:"data"`
}
