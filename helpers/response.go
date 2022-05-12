package helpers

import "strings"

//Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// EmptyObj is used when data doesnt want to be null on json
type EmptyObj struct{}

// BuildResponse method is to inject data value to dynamic success response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}

	return res
}

// BuildErrorResponse metho is to inject errors value to dinamic error response
func BuildErrorResponse(message string, errors string, data interface{}) Response {
	splitError := strings.Split(errors, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Error:   splitError,
		Data:    data,
	}
	return res
}
