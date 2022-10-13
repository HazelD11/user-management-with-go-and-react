package web

type webResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors"`
}
