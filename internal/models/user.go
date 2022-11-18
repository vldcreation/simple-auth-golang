package models

import "github.com/golang-jwt/jwt/v4"

const (
	AppKey string = "edufund-pretest"
)

type (
	JwtClaim struct {
		UserId   int64  `json:"user_id"`
		Fullname string `json:"fullname"`
		Username string `json:"username"`
		Email    string `json:"email"`
		jwt.StandardClaims
	}

	RegisterRequest struct {
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	RegisterResponse struct {
		RegisterScanner RegisterScanner       `json:"user"`
		Token           GenerateTokenResponse `json:"token"`
	}

	RegisterScanner struct {
		UserId   int64  `json:"user_id"`
		Fullname string `json:"fullname"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	GenerateTokenResponse struct {
		UserId      int64  `json:"user_id"`
		Fullname    string `json:"fullname"`
		Username    string `json:"username"`
		Email       string `json:"email"`
		AccessToken string `json:"access_token"`
	}
)
