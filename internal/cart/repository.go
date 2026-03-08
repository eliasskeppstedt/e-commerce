package cart

import (
	"context"
	"database/sql"
)

type CartRepository interface {
	CreateCart(ctx context.Context, tx *sql.Tx, userID int) (*Cart, error)
	GetCartByUserID(ctx context.Context, tx *sql.Tx, userID int) (*Cart, error)

	AddItem(ctx context.Context, tx *sql.Tx, cartID, productID, quantity int) error
	RemoveItem(ctx context.Context, tx *sql.Tx, cartID, productID, quantity int) error

	GetItems(ctx context.Context, tx *sql.Tx, cartID int) ([]CartItem, error)
}

type mysqlCartRepository struct {
	db *sql.DB
}

func NewMysqlCartRepository(db *sql.DB) *mysqlCartRepository {
	return &mysqlCartRepository{db}
}

func (r *mysqlCartRepository) CreateCart(ctx context.Context, tx *sql.Tx, userID int) (*Cart, error) {
	res, err := tx.ExecContext(
		ctx,
		`INSERT INTO carts (user_id) 
		VALUES (?)`,
		userID,
	)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &Cart{CartID: int(id), UserID: userID, Status: "active"}, nil
}

func (r *mysqlCartRepository) GetCartByUserID(ctx context.Context, tx *sql.Tx, userID int) (*Cart, error) {
	var cart Cart

	row := tx.QueryRowContext(
		ctx,
		"SELECT cart_id FROM carts WHERE user_id = ? AND status = 'active'",
		userID,
	)
	err := row.Scan(&cart.CartID)

	if err == sql.ErrNoRows {
		return r.CreateCart(ctx, tx, userID)
	}

	if err != nil {
		return nil, err
	}
	cart.UserID = userID
	return &cart, nil
}

func (r *mysqlCartRepository) AddItem(ctx context.Context, tx *sql.Tx, cartID, productID, quantity int) error {
	_, err := tx.ExecContext(
		ctx,
		`INSERT INTO cart_items (cart_id, product_id, quantity)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE 
			quantity = quantity + VALUES(quantity)
		`,
		cartID,
		productID,
		quantity,
	)

	return err
}

func (r *mysqlCartRepository) RemoveItem(ctx context.Context, tx *sql.Tx, cartID, productID, quantity int) error {
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

func (r *mysqlCartRepository) GetItems(ctx context.Context, tx *sql.Tx, cartID int) ([]CartItem, error) {
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
