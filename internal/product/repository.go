package product

import "database/sql"

type productRepository interface {
	getByProductID(product_id int) (product, error)
	getAll() ([]product, error)
	registerProduct(p product) error
	deleteProduct(id int) error
}

type productRepositoryImpl struct {
	products []product
}

type mysqlProductRepository struct {
	db *sql.DB
}

func NewMysqlProductRepository(db *sql.DB) *mysqlProductRepository {
	return &mysqlProductRepository{db: db}
}

func (r *mysqlProductRepository) getByProductID(product_id int) (product, error) {
	var p product
	err := r.db.QueryRow(
		"SELECT id, name, description, price, category_id FROM products WHERE id = ?", product_id).
		Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.CategoryID)
	if err != nil {
		return product{}, err
	}
	return p, nil
}

func (r *mysqlProductRepository) getAll() ([]product, error) {
	rows, err := r.db.Query("SELECT id, name, description, price, category_id FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []product{}
	for rows.Next() {
		var p product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.CategoryID)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil

}

func (r *mysqlProductRepository) registerProduct(product product) error {
	_, err := r.db.Exec(
		"INSERT INTO products (name, description, price, category_id) VALUES (?,?,?,?)",
		product.Name, product.Description, product.Price, product.CategoryID,
	)
	return err
}

func (r *mysqlProductRepository) deleteProduct(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}
