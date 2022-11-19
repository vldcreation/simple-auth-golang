package user_test

import (
	"context"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/vldcreation/simple-auth-golang/internal/constants"
	"github.com/vldcreation/simple-auth-golang/internal/models"
	"github.com/vldcreation/simple-auth-golang/internal/repository"
	"github.com/vldcreation/simple-auth-golang/internal/usecase/user"
	uc_utils "github.com/vldcreation/simple-auth-golang/internal/usecase/utils"
)

func TestRegistrationUsecase(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	constants.DB = sqlx.NewDb(db, "sqlmock")

	newPg := repository.NewPostgreSQL(constants.DB)

	encryptPassword, err := uc_utils.EncryptPassword("password")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when encrypt password", err)
	}

	patchesEncryptedPassword := gomonkey.ApplyFunc(uc_utils.EncryptPassword,
		func(password string) (string, error) {
			return encryptPassword, nil
		})
	patchesGenToken := gomonkey.ApplyFunc(uc_utils.GenerateToken,
		func(userId int64, fullname string, username string, email string) (res models.GenerateTokenResponse, err error) {
			res.AccessToken = "accesstoken$#@#@$@#"
			return res, nil
		})

	defer patchesEncryptedPassword.Reset()
	defer patchesGenToken.Reset()

	args := []driver.Value{
		"Fullname Test",
		"test",
		"test@gmail.com",
		encryptPassword,
	}

	mock.ExpectQuery(`insert into edufund.users`).
		WithArgs(args...).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	mock.ExpectQuery(`select`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(
			[]string{"id", "fullname", "username", "email"}).
			AddRow(1, "Fullname Test", "test", "test@gmail.com"))

	req := models.RegisterRequest{
		Fullname: "Fullname Test",
		Username: "test",
		Email:    "test@gmail.com",
		Password: "password",
	}

	res, err := user.RegistrationUsecase(ctx, newPg, req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, int64(1), res.RegisterScanner.UserId)
	assert.Equal(t, req.Fullname, res.RegisterScanner.Fullname)
	assert.Equal(t, req.Username, res.RegisterScanner.Username)
	assert.Equal(t, req.Email, res.RegisterScanner.Email)
	assert.NotEmpty(t, res.Token)

	if err != nil {
		t.Error("invalid return ")
	}
}

func TestLoginUsecase(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	constants.DB = sqlx.NewDb(db, "sqlmock")

	newPg := repository.NewPostgreSQL(constants.DB)

	mock.ExpectQuery(`select u."id", u."fullname", u."username", u."email"`).
		WithArgs("test", "test@gmail.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "fullname", "username", "email"}).
			AddRow(1, "Fullname test", "test", "test@gmail.com"))

	patchesComparePassword := gomonkey.ApplyFunc(uc_utils.CompareHashAndPassword,
		func(hash string, password string) bool {
			return true
		})

	patchesGenToken := gomonkey.ApplyFunc(uc_utils.GenerateToken,

		func(userId int64, fullname string, username string, email string) (res models.GenerateTokenResponse, err error) {
			res.AccessToken = "accesstoken$#@#@$@#"
			return res, nil
		})

	defer patchesComparePassword.Reset()
	defer patchesGenToken.Reset()

	encryptPassword, err := uc_utils.EncryptPassword("password")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when encrypt password", err)
	}

	mock.ExpectQuery(`SELECT u."id", u."fullname", u."username", u."email", u."password"`).
		WithArgs("test", "test@gmail.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "fullname", "username", "email", "password"}).
			AddRow(1, "Fullname test", "test", "test@gmail.com", encryptPassword))

	req := models.AuthRequest{
		Username: "test",
		Email:    "test@gmail.com",
		Password: "password",
	}

	res, err := user.LoginUsecase(ctx, newPg, req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
}
