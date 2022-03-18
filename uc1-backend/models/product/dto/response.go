package dto

type Response struct {
	Status string      `json:"status"`
	Error  interface{} `json:"error"`
	Data   interface{} `json:"data"`
}
