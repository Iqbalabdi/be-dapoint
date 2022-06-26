package response

type ApiResponse struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type ApiResponseSuccess struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
