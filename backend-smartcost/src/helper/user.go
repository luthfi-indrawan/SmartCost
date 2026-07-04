package helper

import (
	"backend-smartcost/src/constants"

	"github.com/gin-gonic/gin"
)

func (h *Helper) GeneratePermissions(role string) []string {
	mapPermissions := map[string][]string{
		constants.RoleOwner:   {"products.read", "products.write", "reports.read", "users.manage", "cashier.operate"},
		constants.RoleCashier: {"products.read", "transactions.operate", "transactions.return"},
	}

	var permissions []string

	for k, p := range mapPermissions {
		if role == k {
			permissions = p
			break
		}
	}

	return permissions
}

func (h *Helper) IsOwner(c *gin.Context) bool {
	role := c.GetString(constants.RoleKey)
	return role == constants.RoleOwner
}

func (h *Helper) IsCashier(c *gin.Context) bool {
	role := c.GetString(constants.RoleKey)
	return role == constants.RoleCashier
}