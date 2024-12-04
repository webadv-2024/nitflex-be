package models

type GoogleLoginRequest struct {
	AccessToken string `json:"access_token" binding:"required"`
} 