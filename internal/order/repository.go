package order

import (
	"database/sql"
)

type OrderRepository interface {
	CheckOut(userID int, orderItems []OrderItem) error
}

type mysqlOrderRepository struct {
	db *sql.DB
}

func NewMysqlOrderRepository(db *sql.DB) *mysqlOrderRepository {
	return &mysqlOrderRepository{db: db}
}

func (r *mysqlOrderRepository) CheckOut(userID int, orderItems []OrderItem) error {
	res, err := r.db.Exec(`
		INSERT INTO orders (user_id)
		VALUES (?)`,
		userID,
	)
	if err != nil {
		return err
	}
	orderID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// for index, value := range slice ...
	for _, orderItem := range orderItems {
		_, err = r.db.Exec(`
			INSERT INTO order_items (order_id, product_id, quantity, price_at_purchase)
			VALUES (?, ?, ?, ?)`,
			orderID,
			orderItem.ProductID,
			orderItem.Quantity,
			orderItem.PriceAtPurchase,
		)

		if err != nil {
			return err
		}
	}
	return nil
}
