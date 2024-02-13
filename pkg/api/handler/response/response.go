package response

import "strings"

type Response struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"statuscode"`
	Message    string `json:"message"`
	Error      any    `json:"error,omitempty"`
	Data       any    `json:"data,omitempty"`
}

func SuccessResponse(statusCode int, message string, data any) Response {
	return Response{
		Success:    true,
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}

func ErrorResponse(statusCode int, message string, err error, data any) Response {
	// seperate error string by newline
	errArr := strings.Split(err.Error(), "\n")
	return Response{
		StatusCode: statusCode,
		Message:    message,
		Error:      errArr,
		Data:       data,
	}
}
