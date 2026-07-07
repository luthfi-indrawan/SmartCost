package controller

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (c *controller) CreateTransaction(ctx context.Context, req *RequestCreateTransaction) (*ResponseCreateTransaction, error) {

	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	now := time.Now()
	transactionID := uuid.NewString()

	if req.Type == "hold" && (req.HoldNote == nil || *req.HoldNote == "") {
		return nil, fmt.Errorf("hold_note is required for hold bill")
	}
	if len(req.Items) == 0 {
		return nil, fmt.Errorf("items cannot be empty")
	}

	var totalSubtotal int
	var responseItems []ResponseTransactionItem

	var bulkInsertItemsQueue []struct {
		id        string
		productID string
		qty       int
		price     int
		subtotal  int
		notes     *string
	}

	for _, item := range req.Items {
		var productName, productSKU string
		err := tx.QueryRowContext(ctx,
			`SELECT name, sku FROM products WHERE id = $1 AND is_active = true AND deleted_at IS NULL`,
			item.ProductID,
		).Scan(&productName, &productSKU)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("product not found or inactive: %s", item.ProductID)
			}
			return nil, err
		}

		resolvedPrice, err := c.resolvePrice(ctx, item.ProductID, item.Qty)
		if err != nil {
			return nil, err
		}

		if item.PriceAtTime != resolvedPrice {
			return nil, fmt.Errorf("price mismatch for product %s: expected %d, got %d", item.ProductID, resolvedPrice, item.PriceAtTime)
		}

		itemSubtotal := resolvedPrice * item.Qty
		totalSubtotal += itemSubtotal

		_, err = tx.ExecContext(ctx,
			`UPDATE products SET stock = stock - $1 WHERE id = $2`,
			item.Qty, item.ProductID,
		)
		if err != nil {
			return nil, err
		}

		itemID := uuid.NewString()

		bulkInsertItemsQueue = append(bulkInsertItemsQueue, struct {
			id        string
			productID string
			qty       int
			price     int
			subtotal  int
			notes     *string
		}{
			id:        itemID,
			productID: item.ProductID,
			qty:       item.Qty,
			price:     resolvedPrice,
			subtotal:  itemSubtotal,
			notes:     item.Notes,
		})

		responseItems = append(responseItems, ResponseTransactionItem{
			ID:        itemID,
			Product:   &ProductRef{ID: item.ProductID, Name: productName, SKU: productSKU},
			Qty:       item.Qty,
			UnitPrice: resolvedPrice,
			Subtotal:  itemSubtotal,
			Notes:     item.Notes,
		})
	}

	total := totalSubtotal - req.DiscountAmount + req.TaxAmount

	status := "PENDING"
	var paymentMethod, amountPaid, changeAmount interface{}
	var completedAt *time.Time

	if req.Type == "direct" {
		status = "COMPLETED"
		if req.Payment == nil {
			return nil, fmt.Errorf("payment is required for direct transaction")
		}
		if req.Payment.AmountPaid < total {
			return nil, fmt.Errorf("amount_paid must be >= total")
		}
		if req.Payment.Change != req.Payment.AmountPaid-total {
			return nil, fmt.Errorf("change must equal amount_paid - total")
		}
		paymentMethod = req.Payment.Method
		amountPaid = req.Payment.AmountPaid
		changeAmount = req.Payment.Change
		completedAt = &now
	}

	_, err = tx.ExecContext(ctx,
		`INSERT INTO transactions 
		(id, transaction_code, type, status, cashier_id, hold_note, subtotal, discount_amount, tax_amount, total, 
		 payment_method, amount_paid, change_amount, completed_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $15)`,
		transactionID, "", req.Type, status, req.CashierID, req.HoldNote,
		totalSubtotal, req.DiscountAmount, req.TaxAmount, total,
		paymentMethod, amountPaid, changeAmount, completedAt, now,
	)
	if err != nil {
		return nil, err
	}

	for _, item := range bulkInsertItemsQueue {
		_, err = tx.ExecContext(ctx,
			`INSERT INTO transaction_items (id, transaction_id, product_id, qty, unit_price, subtotal, notes, created_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
			item.id, transactionID, item.productID, item.qty, item.price, item.subtotal, item.notes, now,
		)
		if err != nil {
			return nil, err
		}
	}

	var transactionCode string
	err = tx.QueryRowContext(ctx,
		`SELECT transaction_code FROM transactions WHERE id = $1`,
		transactionID,
	).Scan(&transactionCode)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	cashierName, _ := c.getCashierName(ctx, req.CashierID)

	var respPayment *ResponsePayment
	if req.Type == "direct" && req.Payment != nil {
		respPayment = &ResponsePayment{
			Method:     req.Payment.Method,
			AmountPaid: req.Payment.AmountPaid,
			Change:     req.Payment.Change,
		}
	}

	return &ResponseCreateTransaction{
		ID:              transactionID,
		TransactionCode: transactionCode,
		Type:            req.Type,
		Status:          status,
		Cashier:         &CashierRef{ID: req.CashierID, Name: cashierName},
		HoldNote:        req.HoldNote,
		Items:           responseItems,
		Summary: ResponseSummary{
			Subtotal:       totalSubtotal,
			DiscountAmount: req.DiscountAmount,
			TaxAmount:      req.TaxAmount,
			Total:          total,
		},
		Payment:     respPayment,
		CreatedAt:   now,
		CompletedAt: completedAt,
	}, nil
}
