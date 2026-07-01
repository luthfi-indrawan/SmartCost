package controller

import "context"

type IController interface {
	Login(ctx context.Context, req *RequestLogin) (res *ResponseLogin, err error)
	Logout(ctx context.Context, req *RequestLogout) (res *ResponseLogout, err error)
	Refresh(ctx context.Context, req *RequestRefresh) (res *ResponseRefresh, err error)
	Me(ctx context.Context, req *RequestMe) (res *ResponseMe, err error)
}