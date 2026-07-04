package controller

import "context"

type IController interface {
	GetVoidLogs(ctx context.Context, req *RequestListQuery) (*ResponseVoidLogList, error)
}