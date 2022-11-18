package account_creation

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

func New(c Configuration, d Dependency) feature.SetupUser {
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
	ErrMinimumFullname = errors.New("Name should be 2 characters or more")
	ErrMinimumUsername = errors.New("Username should be 6 characters or more")
	ErrMinimumPassword = errors.New("Password should be at least 12 characters long")
	ErrInvalidEmail    = errors.New("Please provide a valid email address")
	ErrConfirmPassword = errors.New("Confirmation password does not match")
	ErrUserExist       = errors.New("user already exist with same username or email")
)

func (x *feat) SetupUser(
	/*req*/ ctx context.Context, request feature.SetupUserRequest) (
	/*res*/ response feature.SetupUserResponse, httpcode int, err error,
) {
	newPG := repository.NewPostgreSQL(x.Postgresql)
	env := entity.ENV()
	verifier := mail.NewVerifier()

	if len(request.Fullname) < feature.MinimumFullname {
		return response, http.StatusBadRequest, ErrMinimumFullname
	}

	if len(request.Username) < feature.MinimumUsername {
		return response, http.StatusBadRequest, ErrMinimumUsername
	}

	serverHostName := env.Get(constants.ServerHostName)
	serverMailAddress := env.Get(constants.ServerMailAddress)

	log.Printf("serverHostName: %s serverMailAddress %s", serverHostName, serverMailAddress)

	ret, err := verifier.Verify(request.Email)
	if err != nil {
		log.Printf("[SetupUser] - ValidateHostAndUser: %s", err.Error())

		return response, http.StatusBadRequest, ErrInvalidEmail
	}
	if !ret.Syntax.Valid {
		log.Printf("[SetupUser] - EmailSyntax: %v", ret)

		return response, http.StatusBadRequest, ErrInvalidEmail
	}

	if len(request.Password) < feature.MinimumPassword {
		return response, http.StatusBadRequest, ErrMinimumPassword
	}

	if request.Password != request.ConfirmPassword {
		return response, http.StatusBadRequest, ErrConfirmPassword
	}

	// check if user already exist
	curUser, err := newPG.GetUserByEmailOrUsername(ctx, repository.GetUserByEmailOrUsernameRequest{
		Email:    request.Email,
		Username: request.Username,
	})
	if err != nil {
		log.Printf("[SetupUser] - GetUserByEmailOrUsername %v", err)

		return response, http.StatusBadRequest, err
	}

	if curUser.ID != 0 {
		return response, http.StatusBadRequest, ErrUserExist
	}

	req := models.RegisterRequest{
		Fullname: request.Fullname,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	res, err := uc_user.RegistrationUsecase(ctx, newPG, req)
	if err != nil {
		log.Printf("[SetupUser] - RegistrationUsecase %v", err)

		return response, http.StatusBadRequest, err
	}

	response.User = feature.UserResponse{
		ID:       res.RegisterScanner.UserId,
		Fullname: res.RegisterScanner.Fullname,
		Username: res.RegisterScanner.Username,
		Email:    res.RegisterScanner.Email,
	}

	response.Token = res.Token

	return response, http.StatusOK, nil
}
