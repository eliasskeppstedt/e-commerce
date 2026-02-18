package products

import ( 
	"errors"
)

type productRepository interface {
	getByProductID(productID int) (Product, error)
	getAll() ([]Product, error)
	add(p Product) error
	delete(id int) error
}

type productRepositoryImpl struct  {
	products []Product
}
