package models

type GoogleLoginRequest struct {
	AccessToken string `json:"access_token" binding:"required"`
}

type GoogleUserInfo struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Picture  string `json:"picture"`
	Verified bool   `json:"verified_email"`
}
