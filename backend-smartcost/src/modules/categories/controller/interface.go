package controller

import "context"

type IController interface {
	GetListCategories(ctx context.Context, req *RequestGetListCategories) (res *ResponseGetListCategories, err error)
	AddCategory(ctx context.Context, req *RequestAddCategory) (res *ResponseAddCategory, err error)
	UpdateCategory(ctx context.Context, req *RequestUpdateCatregory) (res *ResponseUpdateCategory, err error)
	DeleteCategory(ctx context.Context, req *RequestDeleteCategory) (res *ResponseDeleteCategory, err error)
}