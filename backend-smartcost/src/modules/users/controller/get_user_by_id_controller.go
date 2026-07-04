package controller

import (
	"context"
	"database/sql"
	"errors"
)

func (c *controller) GetUsersByID(
	ctx context.Context,
	req *GetUsersByIDRequest,
) (*GetUsersByIDResponse, error) {

	query := `
SELECT
	u.id,
	u.name,
	u.email,
	u.role,
	u.phone,
	u.is_active,
	u.created_at,
	u.updated_at,

	COALESCE(
		SUM(t.total)
		FILTER (
			WHERE
				t.status = 'COMPLETED'
				AND t.created_at >= CURRENT_DATE
		),
		0
	) AS total_sales_today,

	COALESCE(
		SUM(t.total)
		FILTER (
			WHERE
				t.status = 'COMPLETED'
				AND t.created_at >= DATE_TRUNC('month', CURRENT_DATE)
		),
		0
	) AS total_sales_month,

	COUNT(*)
	FILTER (
		WHERE
			t.status = 'COMPLETED'
			AND t.created_at >= CURRENT_DATE
	) AS transaction_count_today,

	COUNT(*)
	FILTER (
		WHERE
			t.status = 'COMPLETED'
			AND t.created_at >= DATE_TRUNC('month', CURRENT_DATE)
	) AS transaction_count_month

FROM users u
LEFT JOIN transactions t
	ON t.cashier_id = u.id

WHERE
	u.deleted_at IS NULL
	AND u.id = $1

GROUP BY
	u.id,
	u.name,
	u.email,
	u.role,
	u.phone,
	u.is_active,
	u.created_at,
	u.updated_at
`

	var result GetUsersByIDResponse

	err := c.db.QueryRowContext(
		ctx,
		query,
		req.UserID,
	).Scan(
		&result.ID,
		&result.Name,
		&result.Email,
		&result.Role,
		&result.Phone,
		&result.IsActive,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.Stats.TotalSalesToday,
		&result.Stats.TotalSalesMonth,
		&result.Stats.TransactionCountToday,
		&result.Stats.TransactionCountMonth,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	result.Permissions = c.helper.GeneratePermissions(result.Role)

	return &result, nil
}
