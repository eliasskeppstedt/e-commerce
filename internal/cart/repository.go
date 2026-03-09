package cart

import (
	"context"
	"database/sql"
)

type CartRepository interface {
	CreateCart(ctx context.Context, tx *sql.Tx, userID int) (*Cart, error)
	GetCartByUserID(ctx context.Context, tx *sql.Tx, userID int) (*Cart, error)
	GetCartItemQuantity(ctx context.Context, tx *sql.Tx, cartID, productID int) (int, error)

	AddItem(ctx context.Context, tx *sql.Tx, cartID, productID, quantity int) error
	RemoveItem(ctx context.Context, tx *sql.Tx, cartID, productID, quantity int) error
	DeleteItem(ctx context.Context, tx *sql.Tx, cartID, productID int) error

	RequestCartItems(ctx context.Context, tx *sql.Tx, cartID int) ([]CartItemRequest, error)
	MarkCartAsCheckedOut(ctx context.Context, tx *sql.Tx, cartID int) error
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
		`SELECT cart_id 
		FROM carts 
		WHERE user_id = ? AND status = 'active'`,
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
			quantity = quantity + VALUES(quantity)`,
		cartID,
		productID,
		quantity,
	)

	return err
}

func (r *mysqlCartRepository) RemoveItem(ctx context.Context, tx *sql.Tx, cartID, productID, quantity int) error {
	_, err := tx.ExecContext(
		ctx,
		`UPDATE cart_items
		SET quantity = quantity - ?
		WHERE cart_id = ? AND product_id = ?`,
		quantity,
		cartID,
		productID,
	)
	return err
}

func (r *mysqlCartRepository) DeleteItem(ctx context.Context, tx *sql.Tx, cartID, productID int) error {
	_, err := tx.ExecContext(
		ctx,
		`DELETE FROM cart_items
		WHERE cart_id = ? and product_id = ?`,
		cartID,
		productID,
	)
	return err
}

func (r *mysqlCartRepository) RequestCartItems(ctx context.Context, tx *sql.Tx, cartID int) ([]CartItemRequest, error) {
	rows, err := tx.QueryContext(
		ctx,
		`SELECT 
			p.product_id, 
			p.product_name, 
			p.price, 
			ci.quantity, 
			(p.price * ci.quantity) AS subtotal
		FROM cart_items ci
		JOIN products p ON ci.product_id = p.product_id
		WHERE ci.cart_id = ?`,
		cartID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []CartItemRequest

	for rows.Next() {
		var item CartItemRequest

		err := rows.Scan(&item.ProductID, &item.ProductName, &item.Price, &item.Quantity, &item.Subtotal)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *mysqlCartRepository) MarkCartAsCheckedOut(ctx context.Context, tx *sql.Tx, cartID int) error {
	_, err := tx.ExecContext(
		ctx,
		`UPDATE carts
		SET status = 'ordered'
		WHERE cart_id = ?`,
		cartID,
	)
	// hmm nu vid närmare eftertanke borde vi kanske göra en ny cart här, och endast göra första karten vid registrering
	return err
}

func (r *mysqlCartRepository) GetCartItemQuantity(ctx context.Context, tx *sql.Tx, cartID, productID int) (int, error) {
	row := tx.QueryRowContext(
		ctx,
		`SELECT quantity
		FROM cart_items 
		WHERE cart_id = ? AND product_id = ?`,
		cartID, productID,
	)

	var quantity int
	err := row.Scan(&quantity)

	if err == sql.ErrNoRows {
		return 0, nil
	} else if err != nil {
		return -1, err
	}
	return quantity, nil
}
