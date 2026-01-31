package auth

type RegisterRequestPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequestPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ValidateOtpRequestPayload struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}