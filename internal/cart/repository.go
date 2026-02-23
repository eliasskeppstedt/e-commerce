package cart

import (
	"database/sql"
)

type CartRepository interface {
	cartItemAdd(carttID, productID int) error
	cartItemBulkAdd(cartID, productID int) error
	cartItemRemove(cartID, productID int) error
	cartItemBulkRemove(cartID, productID int) error
}

type mysqlCartRepository struct {
	db *sql.DB
}

func NewMysqlCartRepository(db *sql.DB) *mysqlCartRepository {
	return &mysqlCartRepository{db}
}

func (r *mysqlCartRepository) cartItemAdd(cartID, productID int) error {
	_, err := r.db.Exec(`
		INSERT INTO cartItems (cart_id, product_id)
		VALUES (?, ?)`,
		cartID,
		productID,
	)

	if err != nil {
		return err
	}

	_, err = r.db.Exec(`
		UPDATE cartItems
		SET quantity = quantity + 1
		WHERE cart_id = ? AND product_id = ?;`,
		cartID,
		productID,
	)

	return err
}

func (r *mysqlCartRepository) cartItemBulkAdd(cartID, productID int) error {
	return nil
}

func (r *mysqlCartRepository) cartItemRemove(cartID, productID int) error {
	_, err := r.db.Exec(`
		UPDATE cartItems
		SET quantity = quantity - 1
		WHERE cart_id = ? AND product_id = ?;`,
		cartID,
		productID,
	)
	return err
}

func (r *mysqlCartRepository) cartItemBulkRemove(cartID, productID int) error {
	return nil
}
