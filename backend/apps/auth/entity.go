package auth

import (
	"ariskaAdi-pretest-ai/utils"
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	ROLE_ADMIN Role = "admin"
	ROLE_USER  Role = "user"
)

type AuthEntity struct {
	Id        int    `db:"id"`
	UserPublicId  uuid.UUID `db:"public_id"`
	Username  string `db:"username"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	OTP       string `db:"otp"`
	Role 	Role   `db:"role"`
	Verified  bool   `db:"verified"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewFormRegisterRequest(req RegisterRequestPayload) AuthEntity {
	return AuthEntity{
		UserPublicId: uuid.New(),
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Verified: false,
		Role: ROLE_USER,
		OTP: utils.GenerateOtp(6),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewFormLoginRequest(req LoginRequestPayload) AuthEntity {
	return AuthEntity{
		Email: req.Email,
		Password: req.Password,
	}
}

func NewFormValidateOtpRequest(req ValidateOtpRequestPayload) AuthEntity {
	return AuthEntity{
		Email: req.Email,
		OTP: req.OTP,
	}
}