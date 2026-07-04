package controller

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func (c *controller) CompleteHoldBill(ctx context.Context, req *RequestCompleteHoldBill) (*ResponseCompleteHoldBill, error) {
	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Verify transaction exists and is hold bill
	var txnType, txnStatus, txnCashierID string
	var currentTotal int
	err = tx.QueryRowContext(ctx,
		`SELECT type, status, cashier_id, total 
		FROM transactions WHERE id = $1`,
		req.TransactionID,
	).Scan(&txnType, &txnStatus, &txnCashierID, &currentTotal)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("transaction not found")
		}
		return nil, err
	}

	// Validation
	if txnType != "hold" {
		return nil, fmt.Errorf("transaction is not a hold bill")
	}
	if txnStatus != "PENDING" {
		return nil, fmt.Errorf("transaction already completed")
	}
	if txnCashierID != req.CashierID {
		return nil, fmt.Errorf("forbidden: not your hold bill")
	}

	// Validate payment
	if req.Payment.AmountPaid < currentTotal {
		return nil, fmt.Errorf("amount_paid must be >= total")
	}
	expectedChange := req.Payment.AmountPaid - currentTotal
	if req.Payment.Change != expectedChange {
		return nil, fmt.Errorf("change must equal amount_paid - total")
	}

	now := time.Now()

	// Update transaction
	_, err = tx.ExecContext(ctx,
		`UPDATE transactions 
		SET status = 'COMPLETED', 
			payment_method = $1,
			amount_paid = $2,
			change_amount = $3,
			completed_at = $4,
			updated_at = $4
		WHERE id = $5`,
		req.Payment.Method, req.Payment.AmountPaid, req.Payment.Change, now, req.TransactionID,
	)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &ResponseCompleteHoldBill{
		ID:              req.TransactionID,
		TransactionCode: "", // Will be fetched if needed
		Status:          "COMPLETED",
		Payment: ResponsePayment{
			Method:     req.Payment.Method,
			AmountPaid: req.Payment.AmountPaid,
			Change:     req.Payment.Change,
		},
		CompletedAt: now,
	}, nil
}