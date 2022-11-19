package user

import (
	"context"
	"errors"
	"log"

	"github.com/vldcreation/simple-auth-golang/internal/models"
	"github.com/vldcreation/simple-auth-golang/internal/repository"
	sconstants "github.com/vldcreation/simple-auth-golang/internal/service/constants"
	uc_utils "github.com/vldcreation/simple-auth-golang/internal/usecase/utils"
)

func RegistrationUsecase(ctx context.Context, newPG repository.PostgreSQL, form models.RegisterRequest) (resp models.RegisterResponse, err error) {
	password, err := uc_utils.EncryptPassword(form.Password)
	if err != nil {
		return resp, err
	}

	value, err := newPG.AddNewUser(ctx, repository.AddNewUserRequest{
		Fullname: form.Fullname,
		Username: form.Username,
		Email:    form.Email,
		Password: password,
	})
	if err != nil {
		return resp, err
	}

	user, err := newPG.GetUserByID(ctx, repository.GetUserByIDRequest{
		ID: value.LastInsertID,
	})
	if err != nil {
		return resp, err
	}

	auth, err := uc_utils.GenerateToken(user.ID, user.Fullname, user.Username, user.Email)
	if err != nil {
		return resp, err
	}

	resp.RegisterScanner = models.RegisterScanner{
		UserId:   user.ID,
		Fullname: user.Fullname,
		Username: user.Username,
		Email:    user.Email,
	}

	resp.Token = auth

	return resp, nil
}

func LoginUsecase(ctx context.Context, newPG repository.PostgreSQL, form models.AuthRequest) (resp models.AuthResponse, err error) {
	if form.Email == "" {
		form.Email = form.Username
	}

	if form.Username == "" {
		form.Username = form.Email
	}

	log.Printf("form: %+v", form)

	value, err := newPG.GetUserByEmailOrUsername(ctx, repository.GetUserByEmailOrUsernameRequest{
		Username: form.Username,
		Email:    form.Email,
	})
	if err != nil {
		return resp, err
	}

	log.Printf("value: %+v", value)

	if value.ID == 0 {
		log.Printf("value.ID: %+v", form)

		return resp, errors.New(sconstants.UserNotExist)
	}

	// we was check existing account , so we don't need to check again
	userLogin, err := newPG.UserLoginWithEmailOrUsername(ctx, repository.UserLoginWithEmailOrUsernameRequest{
		Username: form.Username,
		Email:    form.Email,
	})
	if err != nil {
		return resp, err
	}

	if isValid := uc_utils.CompareHashAndPassword(userLogin.Password, form.Password); !isValid {
		log.Printf("UserLogin :%v", userLogin)

		return resp, errors.New(sconstants.UserNotExist)
	}

	auth, err := uc_utils.GenerateToken(userLogin.ID, userLogin.Fullname, userLogin.Username, userLogin.Email)
	if err != nil {
		return resp, err
	}

	resp.AuthScanner = models.AuthScanner{
		UserId:   userLogin.ID,
		Fullname: userLogin.Fullname,
		Username: userLogin.Username,
		Email:    userLogin.Email,
	}

	resp.Token = auth

	return resp, nil
}
