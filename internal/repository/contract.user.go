package repository

import "context"

type AddNewUser interface {
	AddNewUser(
		/*req*/ ctx context.Context, request AddNewUserRequest) (
		/*res*/ response AddNewUserResponse, err error,
	)
}

type AddNewUserRequest struct {
	Fullname string
	Username string
	Email    string
	Password string
}

type AddNewUserResponse struct {
	LastInsertID int64
}

type GetUserByID interface {
	GetUserByID(
		/*req*/ ctx context.Context, request GetUserByIDRequest) (
		/*res*/ response GetUserByIDResponse, err error,
	)
}

type GetUserByIDRequest struct {
	ID int64
}

type GetUserByIDResponse struct {
	ID       int64
	Fullname string
	Email    string
	Username string
}

type GetUserByEmailOrUsername interface {
	GetUserByEmailOrUsername(
		/*req*/ ctx context.Context, request GetUserByEmailOrUsernameRequest) (
		/*res*/ response GetUserByEmailOrUsernameResponse, err error,
	)
}

type GetUserByEmailOrUsernameRequest struct {
	Username string
	Email    string
}

type GetUserByEmailOrUsernameResponse struct {
	ID       int64
	Fullname string
	Username string
	Email    string
}
