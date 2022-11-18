package utils

import (
	"errors"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/vldcreation/simple-auth-golang/internal/models"
	"github.com/vldcreation/simple-auth-golang/internal/utils/utstring"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(pwd string) (result string, err error) {
	bt, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[admin][EncryptPassword] Error while encrypt password: %v ", err)

		return "", err
	}

	return string(bt), nil
}

func GenerateToken(userId int64, fullname string, username string, email string) (result models.GenerateTokenResponse, err error) {
	claims := models.JwtClaim{
		UserId:   userId,
		Fullname: fullname,
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    models.AppKey,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(utstring.Env("JWT_SECRET_KEY")))
	if err != nil {
		return result, errors.New("while validate and encrypts")
	}

	result = models.GenerateTokenResponse{
		UserId:      userId,
		Fullname:    fullname,
		Username:    username,
		Email:       email,
		AccessToken: signedToken,
	}

	return result, nil
}

func CompareHashAndPassword(hashedPwd string, plainPwd string) (result bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))

	return err == nil
}
