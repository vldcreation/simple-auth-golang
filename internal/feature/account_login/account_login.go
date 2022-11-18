package account_login

import (
	"context"
	"errors"
	"log"
	"net/http"

	mail "github.com/AfterShip/email-verifier"
	"github.com/vldcreation/simple-auth-golang/internal/constants"
	"github.com/vldcreation/simple-auth-golang/internal/entity"
	"github.com/vldcreation/simple-auth-golang/internal/feature"
	"github.com/vldcreation/simple-auth-golang/internal/models"
	"github.com/vldcreation/simple-auth-golang/internal/repository"
	uc_user "github.com/vldcreation/simple-auth-golang/internal/usecase/user"
)

func New(c Configuration, d Dependency) feature.AccountLogin {
	return &feat{c, d}
}

type Configuration struct{}

type Dependency struct {
	Postgresql repository.SQLConn
}

type feat struct {
	Configuration
	Dependency
}

var (
	ErrInvalidEmail    = errors.New("Please provide a valid email address")
	ErrMinimumPassword = errors.New("Password should be at least 12 characters long")
	ErrUserNotExist    = errors.New("Invalid username / password")
)

func (x *feat) AccountLogin(
	/*req*/ ctx context.Context, request feature.AccountLoginRequest) (
	/*res*/ response feature.AccountLoginResponse, httpcode int, err error,
) {
	newPG := repository.NewPostgreSQL(x.Postgresql)
	env := entity.ENV()
	verifier := mail.NewVerifier()

	serverHostName := env.Get(constants.ServerHostName)
	serverMailAddress := env.Get(constants.ServerMailAddress)

	log.Printf("serverHostName: %s serverMailAddress %s", serverHostName, serverMailAddress)

	var emailField feature.EmailField = feature.EmailField(request.Email)
	verifyEmail := emailField.Coallesce(request.Username)

	ret, err := verifier.Verify(verifyEmail)
	if err != nil {
		log.Printf("[AccountLogin] - VerifyEmail: %s", err.Error())

		return response, http.StatusBadRequest, ErrInvalidEmail
	}
	if !ret.Syntax.Valid {
		log.Printf("[AccountLogin] - EmailSyntax: %v", ret)

		return response, http.StatusBadRequest, ErrInvalidEmail
	}

	if len(request.Password) < feature.MinimumPassword {
		return response, http.StatusBadRequest, ErrMinimumPassword
	}

	req := models.AuthRequest{
		Email:    request.Email,
		Password: request.Password,
	}

	res, err := uc_user.LoginUsecase(ctx, newPG, req)
	if err != nil {
		log.Printf("[AccountLogin] - LoginUsecase %v", err)

		return response, http.StatusBadRequest, err
	}

	response.User = feature.UserResponse{
		ID:       res.AuthScanner.UserId,
		Fullname: res.AuthScanner.Fullname,
		Username: res.AuthScanner.Username,
		Email:    res.AuthScanner.Email,
	}

	response.Token = res.Token

	return response, http.StatusOK, nil
}
