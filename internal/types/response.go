package types

type BillboardsResponse struct {
	// Code is the code of the response
	Code int `json:"code"`
	// Message is the message of the response
	Message string `json:"message"`
	// Data is the data of the response
	Data interface{} `json:"data"`
}

func SuccessResponse(data interface{}) *BillboardsResponse {
	return &BillboardsResponse{
		Code:    0,
		Message: "success",
		Data:    data,
	}
}

func ErrorResponse(code int, message string) *BillboardsResponse {
	if code >= 0 {
		code = -1
	}
	return &BillboardsResponse{
		Code:    code,
		Message: message,
	}
}
