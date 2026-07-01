package controller

import "context"

func (c *controller) Login(ctx context.Context, req *RequestLogin) (res *ResponseLogin, err error) {
	var (
		email = req.Email
		password = req.Password

	)

	c.db.QueryRowContext(
		ctx,
		`select 
			name,
			email,
			password_hash,
			role,
			phone,
			i
		`
	)
}