package controller

import "backend-smartcost/src/types"

type (
	RequestGetListCategories  struct{}

	ResponseGetListCategories struct {
		Data []types.CategoryType
		Metadata types.MetadataType
	}
)

type (
	RequestAddCategory struct {
		Name string
		Color string
		Description *string
	}

	ResponseAddCategory struct {
		types.CategoryType
	}
)

type (
	RequestUpdateCatregory struct {
		CategoryID string
		Name *string
		Color *string
		Description *string
	}

	ResponseUpdateCategory struct {
		types.CategoryType
	}
)

type (
	RequestDeleteCategory struct {
		CategoryID string
	}

	ResponseDeleteCategory struct {
		
	}
)