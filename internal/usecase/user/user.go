package user

import (
	"context"

	"github.com/vldcreation/simple-auth-golang/internal/models"
	"github.com/vldcreation/simple-auth-golang/internal/repository"
	uc_utils "github.com/vldcreation/simple-auth-golang/internal/usecase/utils"
)

func RegistrationUsecase(ctx context.Context, newPG repository.PostgreSQL, form models.RegisterRequest) (resp models.RegisterResponse, err error) {
	password, err := uc_utils.EncryptPassword(form.Password)
	if err != nil {
		return resp, err
	}

	value, err := newPG.AddNewUser(ctx, repository.AddNewUserRequest{
		Fullname: form.Fullname,
		Email:    form.Email,
		Password: password,
		Username: form.Username,
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

	auth, err := uc_utils.GenerateToken(user.ID, user.Username, user.Email)
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
