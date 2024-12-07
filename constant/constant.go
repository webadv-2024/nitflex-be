package constant

import "time"

const (
	ErrorMessage_InternalServerError       = "Internal Server Error"
	ErrorMessage_UsernameExisted           = "Username Existed"
	ErrorMessage_EmailExisted              = "Email Existed"
	ErrorMessage_InvalidUsernameOrPassword = "Invalid Username Or Password"
	ErrorMessage_InvalidGoogleToken        = "invalid google token"
)

const (
	AccessTokenExpriesIn  = 1 * time.Hour
	RefreshTokenExpriesIn = 30 * 24 * time.Hour
)
