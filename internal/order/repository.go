package order

import (
	"context"
	"database/sql"
	"fmt"
)

type OrderRepository interface {
	CheckOut(ctx context.Context, tx *sql.Tx, userID int, subtotal float64, orderItems []OrderItem) error
}

type mysqlOrderRepository struct {
	db *sql.DB
}

func NewMysqlOrderRepository(db *sql.DB) *mysqlOrderRepository {
	return &mysqlOrderRepository{db: db}
}

func (r *mysqlOrderRepository) CheckOut(ctx context.Context, tx *sql.Tx, userID int, subtotal float64, orderItems []OrderItem) error {
	fmt.Println("userID:", userID)
	fmt.Println("subtotal:", subtotal)
	res, err := tx.ExecContext(
		ctx,
		`INSERT INTO orders (user_id, subtotal_at_purchase)
		VALUES (?, ?)`,
		userID, subtotal,
	)
	if err != nil {
		return err
	}
	fmt.Println("hej2")
	orderID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	fmt.Println("hej1")

	// for index, value := range slice ...
	for _, orderItem := range orderItems {
		_, err = tx.ExecContext(
			ctx,
			`INSERT INTO order_items (order_id, product_id, quantity, price_at_purchase)
			VALUES (?, ?, ?, ?)`,
			int(orderID),
			orderItem.ProductID,
			orderItem.Quantity,
			orderItem.PriceAtPurchase,
		)

		if err != nil {
			return err
		}
		fmt.Println("1")
	}
	return nil
}
