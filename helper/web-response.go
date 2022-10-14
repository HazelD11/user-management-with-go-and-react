package helper

import "strings"

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors"`
}

func BuildResponse(code int, status string, data interface{}) WebResponse {
	res := WebResponse{
		Code:   code,
		Status: status,
		Errors: nil,
		Data:   data,
	}
	return res
}

func BuildErrorResponse(code int, status string, err string, data interface{}) WebResponse {
	splittedError := strings.Split(err, "\n")
	res := WebResponse{
		Code:   code,
		Status: "Success",
		Errors: splittedError,
		Data:   data,
	}
	return res
}
