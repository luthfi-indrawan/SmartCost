package controller

import "context"

type IController interface {
	AddProduct(ctx context.Context, req *RequestAddProduct) (*ResponseAddProduct, error)
	GetProductList(ctx context.Context, req *RequestListQuery) (*ResponseProductList, error)
	GetProductDetail(ctx context.Context, id string) (*ResponseGetProduct, error)
	UpdateProduct(ctx context.Context, req *RequestUpdateProduct) (*ResponseUpdateProduct, error)
	DeleteProduct(ctx context.Context, id string) (*ResponseDeleteProduct, error)
}