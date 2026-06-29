package handler

import (
	"backend-smartcost/src/helper"
	"backend-smartcost/src/modules/products/controller"
)

type Handler struct {
	controller controller.IController
	helper *helper.Helper
}

func NewHandler(ctrl controller.IController, helper *helper.Helper) *Handler {
	return &Handler{
		controller: ctrl,
		helper: helper,
	}
}