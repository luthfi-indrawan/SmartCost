package handler

type (
	DTOAddCategory struct {
		Name        string  `json:"name" binding:"required"`
		Color       string  `json:"color" binding:"required,hexcolor"`
		Description *string `json:"description" binding:"omitempty,max=500"`
	}
)

type (
	DTOUpdateCategory struct {
		Name        *string `json:"name" binding:"omitempty"`
		Color       *string `json:"color" binding:"omitempty,hexcolor"`
		Description *string `json:"description" binding:"omitempty,max=500"`
	}
)