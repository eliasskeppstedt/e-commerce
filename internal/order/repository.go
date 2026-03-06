package order

import (
	"database/sql"
)

type OrderRepository interface {
	CheckOut(userID int, orderItems OrderItem) error
	CreateOrder(userID int) (int, error)
	AddOrderItem(item OrderItem) error
	GetOrdersByUser(userID int) ([]OrderWithItems, error)
}

type mysqlOrderRepository struct {
	db *sql.DB
}

func NewMysqlOrderRepository(db *sql.DB) *mysqlOrderRepository {
	return &mysqlOrderRepository{db: db}
}

func (r *mysqlOrderRepository) CheckOut(userID int, orderItem OrderItem) error {
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
	//for _, orderItem := range orderItems {
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
	//}
	return nil
}
func (r *mysqlOrderRepository) CreateOrder(userID int) (int, error) {

	res, err := r.db.Exec(`
		INSERT INTO orders (user_id)
		VALUES (?)
	`, userID)

	if err != nil {
		return 0, err
	}

	orderID, err := res.LastInsertId()

	return int(orderID), err
}

func (r *mysqlOrderRepository) AddOrderItem(item OrderItem) error {

	_, err := r.db.Exec(`
		INSERT INTO order_items
		(order_id, product_id, quantity, price_at_purchase)
		VALUES (?, ?, ?, ?)
	`,
		item.OrderID,
		item.ProductID,
		item.Quantity,
		item.PriceAtPurchase,
	)

	return err
}

func (r *mysqlOrderRepository) GetOrdersByUser(userID int) ([]OrderWithItems, error) {

	rows, err := r.db.Query(`
	SELECT 
		o.order_id,
		o.date,
		o.status,
		p.product_name,
		oi.quantity,
		oi.price_at_purchase
	FROM orders o
	JOIN order_items oi ON o.order_id = oi.order_id
	JOIN products p ON oi.product_id = p.product_id
	WHERE o.user_id = ?
	ORDER BY o.order_id DESC
	`, userID)

	if err != nil {
		return nil, err
	}

	orderMap := map[int]*OrderWithItems{}

	for rows.Next() {

		var orderID int
		var date, status, productName string
		var quantity int
		var price float64

		rows.Scan(&orderID, &date, &status, &productName, &quantity, &price)

		if _, exists := orderMap[orderID]; !exists {

			orderMap[orderID] = &OrderWithItems{
				OrderID: orderID,
				Date:    date,
				Status:  status,
			}
		}

		orderMap[orderID].Items = append(orderMap[orderID].Items,
			OrderItemWithProduct{
				ProductName: productName,
				Quantity:    quantity,
				Price:       price,
			})
	}

	var orders []OrderWithItems

	for _, o := range orderMap {
		orders = append(orders, *o)
	}

	return orders, nil
}
