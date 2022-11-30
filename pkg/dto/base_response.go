package dto

type BaseResponse struct {
	Status string      `json:"status"`
	Error  *ErrorData  `json:"error"`
	Data   interface{} `json:"data"`
}

type ErrorData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
