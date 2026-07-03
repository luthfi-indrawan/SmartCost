package controller

import (
	"context"
)

type IController interface {
	CreateUsers(ctx context.Context, req *CreateUserRequest) (res *CreateUserResponse, err error)

	GetUsers(ctx context.Context, req *GetUsersRequest) (res *GetUsersResponse, err error)

	GetUsersByID(ctx context.Context, req *GetUsersByIDRequest) (res *GetUsersByIDResponse, err error)

	UpdateUsers(ctx context.Context, req *UpdateUserRequest) (res *UpdateUserResponse, err error)

	DeleteUsers(ctx context.Context, req *DeleteUserRequest) (res *DeleteUserResponse, err error)
}
