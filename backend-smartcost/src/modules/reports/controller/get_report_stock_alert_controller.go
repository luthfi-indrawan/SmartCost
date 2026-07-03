package controller

import (
	"context"
	"time"
)

func (c *controller) GetReportStockAlert(
	ctx context.Context,
	req *GetReportStockAlertRequest,
) (*GetReportStockAlertResponse, error) {

	var res GetReportStockAlertResponse

	criticalQuery := `
		SELECT COUNT(*)
		FROM public.stock_alerts sa
		WHERE sa.is_resolved = false
		AND sa.alert_type = 'MINUS'
	`

	if err := c.db.QueryRowContext(
		ctx,
		criticalQuery,
	).Scan(&res.CriticalCount); err != nil {
		return nil, err
	}

	lowQuery := `
		SELECT COUNT(*)
		FROM public.stock_alerts sa
		WHERE sa.is_resolved = false
		AND sa.alert_type = 'LOW'
	`

	if err := c.db.QueryRowContext(
		ctx,
		lowQuery,
	).Scan(&res.LowCount); err != nil {
		return nil, err
	}

	query := `
		SELECT
			p.id,
			p.name,
			p.stock,
			p.min_stock_threshold,
			sa.alert_type,
			sa.updated_at
		FROM public.stock_alerts sa
		INNER JOIN public.products p
			ON p.id = sa.product_id
		WHERE sa.is_resolved = false
		ORDER BY
			CASE
				WHEN sa.alert_type = 'MINUS' THEN 1
				WHEN sa.alert_type = 'LOW' THEN 2
				ELSE 3
			END,
			sa.created_at DESC
	`

	rows, err := c.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res.Alerts = make([]StockAlertReport, 0)

	for rows.Next() {

		var (
			item        StockAlertReport
			lastUpdated time.Time
		)

		err := rows.Scan(
			&item.ProductID,
			&item.ProductName,
			&item.CurrentStock,
			&item.MinThreshold,
			&item.Status,
			&lastUpdated,
		)
		if err != nil {
			return nil, err
		}

		item.LastUpdated = lastUpdated.Format(time.RFC3339)

		res.Alerts = append(res.Alerts, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &res, nil
}
