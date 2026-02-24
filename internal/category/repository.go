package category

import "database/sql"

// Interface
type categoryRepository interface {
	getAll() ([]Category, error)
	createCategory(category Category) error
	deleteCategory(id int) error
}

// MySQL implementation
type mysqlCategoryRepository struct {
	db *sql.DB
}

func NewMysqlCategoryRepository(db *sql.DB) *mysqlCategoryRepository {
	return &mysqlCategoryRepository{db: db}
}

// HÄMTAR ALLA KATEGORIER
func (r *mysqlCategoryRepository) getAll() ([]Category, error) {
	rows, err := r.db.Query("SELECT category_id, category_name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var c Category
		if err := rows.Scan(
			&c.CategoryID,
			&c.CategoryName,
		); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

func (r *mysqlCategoryRepository) createCategory(c Category) error {
	_, err := r.db.Exec(`
		INSERT INTO categories (category_id, category_name)
		Values (?, ?)`,
		c.CategoryID,
		c.CategoryName,
	)
	return err
}

// TA BORT KATEGORI
func (r *mysqlCategoryRepository) deleteCategory(id int) error {
	_, err := r.db.Exec("DELETE FROM categories WHERE category_id = ?", id)
	return err
}
