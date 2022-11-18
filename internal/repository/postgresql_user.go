package repository

import (
	"context"

	"github.com/vldcreation/simple-auth-golang/internal/entity"
	"github.com/vldcreation/simple-auth-golang/internal/repository/postgresql_query"
)

func (x *postgresql) AddNewUser(
	/*req*/ ctx context.Context, request AddNewUserRequest) (
	/*res*/ response AddNewUserResponse, err error,
) {
	query := postgresql_query.AddNewUser
	args := entity.List{
		request.Fullname,
		request.Username,
		request.Email,
		request.Password,
	}

	row := func(i int) entity.List {
		return entity.List{
			&response.LastInsertID,
		}
	}

	err = new(SQL).BoxQuery(x.tc.QueryContext(ctx, query, args...)).Scan(row)
	if err != nil {
		return response, err
	}

	return response, err
}

func (x *postgresql) GetUserByID(
	/*req*/ ctx context.Context, request GetUserByIDRequest) (
	/*res*/ response GetUserByIDResponse, err error,
) {
	query := postgresql_query.GetUserByID
	args := entity.List{
		request.ID,
	}

	row := func(i int) entity.List {
		return entity.List{
			&response.ID,
			&response.Fullname,
			&response.Username,
			&response.Email,
		}
	}

	err = new(SQL).BoxQuery(x.tc.QueryContext(ctx, query, args...)).Scan(row)
	if err != nil {
		return response, err
	}

	return response, err
}

func (x *postgresql) GetUserByEmailOrUsername(
	/*req*/ ctx context.Context, request GetUserByEmailOrUsernameRequest) (
	/*res*/ response GetUserByEmailOrUsernameResponse, err error,
) {
	query := postgresql_query.GetUserByEmailOrUsername
	args := entity.List{
		request.Username,
		request.Email,
	}

	row := func(i int) entity.List {
		if i > (0) {
			return nil
		}

		return entity.List{
			&response.ID,
			&response.Fullname,
			&response.Username,
			&response.Email,
		}
	}

	err = new(SQL).BoxQuery(x.tc.QueryContext(ctx, query, args...)).Scan(row)
	if err != nil {
		return response, err
	}

	return response, err
}
