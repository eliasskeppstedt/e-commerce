package product

type productService interface {
	getByProductID(id int) (Product, error)
	GetAll() ([]Product, error)
	registerProduct(product Product) error
	deleteProduct(id int) error
}

type ProductServiceImp struct {
	repo ProductRepository
}

func NewProductServiceImp(r ProductRepository) *ProductServiceImp {
	return &ProductServiceImp{repo: r}
}

func (s *ProductServiceImp) getByProductID(id int) (Product, error) {
	return s.repo.GetByProductID(id)
}

func (s *ProductServiceImp) GetAll() ([]Product, error) {
	return s.repo.getAll()
}

func (s *ProductServiceImp) registerProduct(product Product) error {
	return s.repo.registerProduct(product)
}

func (s *ProductServiceImp) deleteProduct(id int) error {
	return s.repo.deleteProduct(id)
}

func (s *ProductServiceImp) updateProduct(id int, stock int, price float64) error {
	return s.repo.updateProduct(id, stock, price)
}
