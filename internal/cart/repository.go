package cart

import (
	"database/sql"
	"fmt"
)

type CartRepository interface {
	CreateCart(userID int) error

	AddItem(cartID, productID, quantity int) error
	RemoveItem(cartID, productID, quantity int) error

	GetItems(cartID int) ([]CartItem, error)

	GetCartByUserID(userID int) (*Cart, error)
	ClearCart(cartID int) error
	SetCartInactive(cartID int) error
}

type mysqlCartRepository struct {
	db *sql.DB
}

func NewMysqlCartRepository(db *sql.DB) *mysqlCartRepository {
	return &mysqlCartRepository{db}
}

func (r *mysqlCartRepository) CreateCart(userID int) error {
	fmt.Println("test2")
	_, err := r.db.Exec(`
		INSERT INTO carts (user_id, status) 
		VALUES (?, 'active')`,
		userID,
	)
	return err
}

func (r *mysqlCartRepository) GetCartByUserID(userID int) (*Cart, error) {
	var cart Cart

	row := r.db.QueryRow(
		"SELECT cart_id FROM carts WHERE user_id = ? AND status = 'active'",
		userID,
	)
	err := row.Scan(&cart.CartID)
	if err != nil {
		fmt.Println(cart.UserID, cart.CartID)
		return nil, err
	}
	cart.UserID = userID

	return &cart, nil
}

func (r *mysqlCartRepository) AddItem(cartID, productID, quantity int) error {
	res, err := r.db.Exec(`
		UPDATE cart_items
		SET quantity = quantity + ?
		WHERE cart_id = ? AND product_id = ?`,
		quantity,
		cartID,
		productID,
	)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		_, err = r.db.Exec(`
			INSERT INTO cart_items (cart_id, product_id, quantity)
			VALUES (?, ?, ?)`,
			cartID,
			productID,
			quantity,
		)
		return err
	}

	return nil
}

func (r *mysqlCartRepository) RemoveItem(cartID, productID, quantity int) error {
	var curQuantity int

	row := r.db.QueryRow(`
        SELECT quantity 
        FROM cart_items 
        WHERE cart_id = ? AND product_id = ?`,
		cartID,
		productID,
	)

	err := row.Scan(&curQuantity)

	if err != nil {
		return err
	}

	if curQuantity <= quantity {
		_, err = r.db.Exec(`
            DELETE FROM cart_items
            WHERE cart_id = ? AND product_id = ?`,
			cartID,
			productID,
		)
		return err
	}

	_, err = r.db.Exec(`
        UPDATE cart_items
        SET quantity = quantity - ?
        WHERE cart_id = ? AND product_id = ?`,
		quantity,
		cartID,
		productID,
	)

	return err
}

func (r *mysqlCartRepository) GetItems(cartID int) ([]CartItem, error) {
	rows, err := r.db.Query(`
		SELECT product_id, quantity
		FROM cart_items
		WHERE cart_id = ?`,
		cartID,
	)
	if err != nil {
		return nil, err
	}

	var items []CartItem

	for rows.Next() {
		var item CartItem

		err := rows.Scan(&item.ProductID, &item.Quantity)
		if err != nil {
			return nil, err
		}

		item.CartID = cartID
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *mysqlCartRepository) ClearCart(cartID int) error {

	_, err := r.db.Exec(`
		DELETE FROM cart_items
		WHERE cart_id = ?
	`, cartID)

	return err
}

func (r *mysqlCartRepository) SetCartInactive(cartID int) error {

	_, err := r.db.Exec(`
		UPDATE carts
		SET status = 'ordered'
		WHERE cart_id = ?
	`, cartID)

	return err
}
