package models

type JSONResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type JSONErrorResponse struct {
	Error string `json:"error"`
}
