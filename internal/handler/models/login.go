package models

type LoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
