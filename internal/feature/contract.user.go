package feature

import (
	"context"

	"github.com/vldcreation/simple-auth-golang/internal/models"
)

const (
	MinimumFullname = 2
	MinimumPassword = 12
	MinimumUsername = 6
)

// SetupUser
//
// ----------------------------------------------------------------------------.
type SetupUser interface {
	SetupUser(
		/*req*/ ctx context.Context, request SetupUserRequest) (
		/*res*/ response SetupUserResponse, httpcode int, err error,
	)
}

type SetupUserRequest struct {
	Fullname        string `json:"fullname"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type SetupUserResponse struct {
	User  UserResponse                 `json:"user"`
	Token models.GenerateTokenResponse `json:"token"`
}

type UserResponse struct {
	ID       int64  `json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type EmailField string

func (x EmailField) Coallesce(def string) string {
	if x == "" {
		return string(def)
	}
	return string(x)
}

// UserLoginWithEmailOrUsername
//
// ----------------------------------------------------------------------------.
type AccountLogin interface {
	AccountLogin(
		/*req*/ ctx context.Context, request AccountLoginRequest) (
		/*res*/ response AccountLoginResponse, htppcode int, err error,
	)
}

type AccountLoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountLoginResponse struct {
	User  UserResponse                 `json:"user"`
	Token models.GenerateTokenResponse `json:"token"`
}
