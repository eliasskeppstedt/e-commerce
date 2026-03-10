package order

import (
	"context"
	"database/sql"
	"fmt"
)

type OrderRepository interface {
	CheckOut(ctx context.Context, tx *sql.Tx, userID int, subtotal float64, orderItems []OrderItem) error
	GetOrdersByUser(ctx context.Context, tx *sql.Tx, userID int) ([]OrderWithItems, error)
	GetAllOrders(ctx context.Context, tx *sql.Tx) ([]UserOrders, error)
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

	orderID, err := res.LastInsertId()
	if err != nil {
		return err
	}

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

func (r *mysqlOrderRepository) GetOrdersByUser(ctx context.Context, tx *sql.Tx, userID int) ([]OrderWithItems, error) {

	rows, err := tx.QueryContext(
		ctx,
		`SELECT 
			o.order_id,
			o.status,
			o.date,
			p.product_id,
			p.product_name,
			oi.quantity,
			oi.price_at_purchase
		FROM orders o
		JOIN order_items oi ON o.order_id = oi.order_id
		JOIN products p ON oi.product_id = p.product_id
		WHERE o.user_id = ?
		ORDER BY o.date DESC`,
		userID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orderMap := map[int]*OrderWithItems{}

	for rows.Next() {

		var (
			orderID int
			status  string
			date    string
			item    OrderItemView
		)

		err := rows.Scan(
			&orderID,
			&status,
			&date,
			&item.ProductID,
			&item.ProductName,
			&item.Quantity,
			&item.Price,
		)

		if err != nil {
			return nil, err
		}

		if _, exists := orderMap[orderID]; !exists {
			orderMap[orderID] = &OrderWithItems{
				OrderID: orderID,
				Status:  status,
				Date:    date,
				Items:   []OrderItemView{},
			}
		}

		orderMap[orderID].Items = append(orderMap[orderID].Items, item)
	}

	var orders []OrderWithItems
	for _, o := range orderMap {
		orders = append(orders, *o)
	}

	return orders, nil
}

func (r *mysqlOrderRepository) GetAllOrders(ctx context.Context, tx *sql.Tx) ([]UserOrders, error) {

	rows, err := tx.QueryContext(
		ctx,
		`SELECT
			u.user_id,
			u.username,
			o.order_id,
			o.status,
			o.date,
			p.product_id,
			p.product_name,
			oi.quantity,
			oi.price_at_purchase
		FROM users u
		JOIN orders o ON u.user_id = o.user_id
		JOIN order_items oi ON o.order_id = oi.order_id
		JOIN products p ON oi.product_id = p.product_id
		ORDER BY u.user_id, o.date DESC`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	userMap := map[int]*UserOrders{}
	orderMap := map[int]*OrderWithItems{}

	for rows.Next() {
		var (
			userID   int
			username string
			orderID  int
			status   string
			date     string
			item     OrderItemView
		)

		err := rows.Scan(
			&userID,
			&username,
			&orderID,
			&status,
			&date,
			&item.ProductID,
			&item.ProductName,
			&item.Quantity,
			&item.Price,
		)

		if err != nil {
			return nil, err
		}

		if _, exists := userMap[userID]; !exists {
			userMap[userID] = &UserOrders{
				UserID:   userID,
				Username: username,
				Orders:   []OrderWithItems{},
			}
		}

		if _, exists := orderMap[orderID]; !exists {

			order := OrderWithItems{
				OrderID: orderID,
				Status:  status,
				Date:    date,
				Items:   []OrderItemView{},
			}

			userMap[userID].Orders = append(userMap[userID].Orders, order)
			orderMap[orderID] = &userMap[userID].Orders[len(userMap[userID].Orders)-1]
		}

		orderMap[orderID].Items = append(orderMap[orderID].Items, item)
	}

	var result []UserOrders
	for _, u := range userMap {
		result = append(result, *u)
	}

	return result, nil
}
