package dto

type Response struct {
	StatusCode int         `json:"statusCode"`
	Status     string      `json:"status"`
	Error      interface{} `json:"error"`
	Data       interface{} `json:"data"`
}
