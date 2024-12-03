package models

type BaseResponse struct {
	code    int    `json:"code"`
	message string `json:"message"`
}
