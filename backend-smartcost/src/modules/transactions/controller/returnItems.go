package controller

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (c *controller) ReturnItems(ctx context.Context, req *RequestReturnItems) (*ResponseReturnItems, error) {
	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Verify transaction is completed
	var txnStatus string
	var currentTotal, currentSubtotal int
	err = tx.QueryRowContext(ctx,
		`SELECT status, total, subtotal FROM transactions WHERE id = $1`,
		req.TransactionID,
	).Scan(&txnStatus, &currentTotal, &currentSubtotal)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("transaction not found")
		}
		return nil, err
	}
	if txnStatus != "COMPLETED" {
		return nil, fmt.Errorf("hold bill cannot be returned, cancel it instead")
	}

	now := time.Now()
	var totalRefund int
	var returnedItems []ResponseReturnedItem
	voidLogID := uuid.NewString()

	for _, item := range req.Items {
		// Get original item details
		var originalQty, unitPrice int
		var productID, productName string
		err := tx.QueryRowContext(ctx,
			`SELECT ti.qty, ti.unit_price, ti.product_id, p.name
			FROM transaction_items ti
			JOIN products p ON ti.product_id = p.id
			WHERE ti.id = $1 AND ti.transaction_id = $2`,
			item.TransactionItemID, req.TransactionID,
		).Scan(&originalQty, &unitPrice, &productID, &productName)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("transaction item not found: %s", item.TransactionItemID)
			}
			return nil, err
		}

		// Check total returned qty so far
		var totalReturnedSoFar sql.NullInt64
		err = tx.QueryRowContext(ctx,
			`SELECT COALESCE(SUM(qty_returned), 0) FROM void_logs 
			WHERE transaction_item_id = $1`,
			item.TransactionItemID,
		).Scan(&totalReturnedSoFar)
		if err != nil {
			return nil, err
		}

		availableQty := originalQty - int(totalReturnedSoFar.Int64)
		if item.Qty > availableQty {
			return nil, fmt.Errorf("return qty exceeds available: max %d", availableQty)
		}

		refundAmount := unitPrice * item.Qty
		totalRefund += refundAmount

		// Return stock
		_, err = tx.ExecContext(ctx,
			`UPDATE products SET stock = stock + $1 WHERE id = $2`,
			item.Qty, productID,
		)
		if err != nil {
			return nil, err
		}

		// Insert void log
		_, err = tx.ExecContext(ctx,
			`INSERT INTO void_logs 
			(id, transaction_id, transaction_item_id, cashier_id, product_id, qty_returned, refund_amount, reason, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $9)`,
			uuid.NewString(), req.TransactionID, item.TransactionItemID, req.CashierID,
			productID, item.Qty, refundAmount, item.Reason, now,
		)
		if err != nil {
			return nil, err
		}

		returnedItems = append(returnedItems, ResponseReturnedItem{
			TransactionItemID: item.TransactionItemID,
			ProductName:       productName,
			QtyReturned:       item.Qty,
			RefundAmount:      refundAmount,
			Reason:            item.Reason,
		})
	}

	// Update transaction totals
	newTotal := currentTotal - totalRefund
	newSubtotal := currentSubtotal - totalRefund

	_, err = tx.ExecContext(ctx,
		`UPDATE transactions 
		SET total = $1, subtotal = $2, updated_at = $3
		WHERE id = $4`,
		newTotal, newSubtotal, now, req.TransactionID,
	)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &ResponseReturnItems{
		TransactionID:  req.TransactionID,
		ReturnedItems:  returnedItems,
		NewTotal:       newTotal,
		VoidLogID:      voidLogID,
		CreatedAt:      now,
	}, nil
}