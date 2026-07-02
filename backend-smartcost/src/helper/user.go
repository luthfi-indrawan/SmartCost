package helper

func (h *Helper) GeneratePermissions(role string) []string {
	mapPermissions := map[string][]string{
		"owner":   {"products.read", "products.write", "reports.read", "users.manage", "cashier.operate"},
		"cashier": {"products.read", "transactions.operate", "transactions.return"},
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