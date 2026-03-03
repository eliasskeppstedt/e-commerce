package category

import (
	"database/sql"
	"fmt"
)

type categoryRepository interface {
	getAll() ([]Category, error)
	createCategory(category Category) error
	deleteCategory(id int) error
}

type mysqlCategoryRepository struct {
	db *sql.DB
}

func NewMysqlCategoryRepository(db *sql.DB) *mysqlCategoryRepository {
	return &mysqlCategoryRepository{db: db}
}

// HÄMTAR ALLA KATEGORIER
func (r *mysqlCategoryRepository) getAll() ([]Category, error) {
	rows, err := r.db.Query(`SELECT category_id, category_name FROM categories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var c Category
		if err := rows.Scan(&c.CategoryID, &c.CategoryName); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

func (r *mysqlCategoryRepository) createCategory(c Category) error {
	_, err := r.db.Exec(`
		INSERT INTO categories (category_name)
		Values (?)`,
		c.CategoryName,
	)
	return err
}

// TA BORT KATEGORI
func (r *mysqlCategoryRepository) deleteCategory(id int) error {
	res, err := r.db.Exec("DELETE FROM categories WHERE category_id = ?", id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("category not found")
	}
	return nil
}
