package utils

type Response struct {
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
}

func NewSuccessResponse(data interface{}, message string, statusCode int) Response {
	return Response{
		Data:       data,
		Message:    message,
		StatusCode: statusCode,
	}
}

func NewErrorResponse(message string, statusCode int) Response {
	return Response{
		Message:    message,
		StatusCode: statusCode,
	}
}
