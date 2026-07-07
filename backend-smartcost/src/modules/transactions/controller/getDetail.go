package controller

import (
	"context"
	"database/sql"
	"fmt"
)

func (c *controller) GetTransactionDetail(ctx context.Context, id string, currentUserID, currentUserRole string) (*ResponseTransactionDetail, error) {
	// First, check access and get transaction
	var txn ResponseTransactionDetail
	var cashierID, cashierName string
	var holdNote, paymentMethod sql.NullString
	var amountPaid, changeAmount sql.NullInt64
	var completedAt, updatedAt sql.NullTime

	err := c.db.QueryRowContext(ctx,
		`SELECT t.id, t.transaction_code, t.type, t.status,
			t.cashier_id, u.name,
			t.hold_note,
			t.subtotal, t.discount_amount, t.tax_amount, t.total,
			t.payment_method, t.amount_paid, t.change_amount,
			t.completed_at, t.created_at, t.updated_at
		FROM transactions t
		JOIN users u ON t.cashier_id = u.id
		WHERE t.id = $1`,
		id,
	).Scan(
		&txn.ID, &txn.TransactionCode, &txn.Type, &txn.Status,
		&cashierID, &cashierName,
		&holdNote,
		&txn.Summary.Subtotal, &txn.Summary.DiscountAmount, &txn.Summary.TaxAmount, &txn.Summary.Total,
		&paymentMethod, &amountPaid, &changeAmount,
		&completedAt, &txn.CreatedAt, &updatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("transaction not found")
		}
		return nil, err
	}

	// Access control
	if currentUserRole == "cashier" && cashierID != currentUserID {
		return nil, fmt.Errorf("forbidden: not your transaction")
	}

	txn.Cashier = &CashierRef{ID: cashierID, Name: cashierName}
	if holdNote.Valid {
		txn.HoldNote = &holdNote.String
	}
	if updatedAt.Valid {
		txn.UpdatedAt = &updatedAt.Time
	}
	if paymentMethod.Valid {
		txn.Payment = &ResponsePayment{
			Method:     paymentMethod.String,
			AmountPaid: int(amountPaid.Int64),
			Change:     int(changeAmount.Int64),
		}
	}

	// Get items
	itemRows, err := c.db.QueryContext(ctx,
		`SELECT ti.id, ti.product_id, p.name, p.sku, ti.qty, ti.unit_price, ti.subtotal, ti.notes
		FROM transaction_items ti
		JOIN products p ON ti.product_id = p.id
		WHERE ti.transaction_id = $1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer itemRows.Close()

	for itemRows.Next() {
		var item ResponseTransactionItem
		var notes sql.NullString

		item.Product = &ProductRef{}

		err := itemRows.Scan(
			&item.ID,
			&item.Product.ID,
			&item.Product.Name,
			&item.Product.SKU,
			&item.Qty,
			&item.UnitPrice,
			&item.Subtotal,
			&notes,
		)
		if err != nil {
			return nil, err
		}

		if notes.Valid {
			item.Notes = &notes.String
		}

		txn.Items = append(txn.Items, item)
	}
	// Get void logs
	voidRows, err := c.db.QueryContext(ctx,
		`SELECT vl.id, vl.transaction_item_id, vl.cashier_id, uc.name,
			vl.product_id, p.name, p.sku,
			vl.qty_returned, vl.refund_amount, vl.reason, vl.created_at
		FROM void_logs vl
		JOIN users uc ON vl.cashier_id = uc.id
		JOIN products p ON vl.product_id = p.id
		WHERE vl.transaction_id = $1
		ORDER BY vl.created_at DESC`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer voidRows.Close()

	for voidRows.Next() {
		var vl ResponseVoidLog

		vl.Cashier = &CashierRef{}
		vl.Product = &ProductRef{}

		err := voidRows.Scan(
			&vl.ID,
			&vl.TransactionItemID,
			&vl.Cashier.ID,
			&vl.Cashier.Name,
			&vl.Product.ID,
			&vl.Product.Name,
			&vl.Product.SKU,
			&vl.QtyReturned,
			&vl.RefundAmount,
			&vl.Reason,
			&vl.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		txn.VoidLogs = append(txn.VoidLogs, vl)
	}
	return &txn, nil
}
