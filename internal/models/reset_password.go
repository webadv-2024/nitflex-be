package models

type RequestResetPasswordRequest struct {
	Email string `json:"email"`
}

type RequestResetPasswordResponse struct {
	ResetPassworkUrl string `json:"reset_password_url"`
}

type ResetPasswordRequest struct {
	Password string `json:"password"`
}
