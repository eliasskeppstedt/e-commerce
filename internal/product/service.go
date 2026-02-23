package product

type productService interface {
	getByProductID(id int) (Product, error)
	getAll() ([]Product, error)
	registerProduct(product Product) error
	deleteProduct(id int) error
}

type productServiceImp struct {
	repo ProductRepository
}

func NewProductServiceImp(r ProductRepository) *productServiceImp {
	return &productServiceImp{repo: r}
}

func (s *productServiceImp) getByProductID(id int) (Product, error) {
	return s.repo.getByProductID(id)
}

func (s *productServiceImp) getAll() ([]Product, error) {
	return s.repo.getAll()
}

func (s *productServiceImp) registerProduct(product Product) error {
	return s.repo.registerProduct(product)
}

func (s *productServiceImp) deleteProduct(id int) error {
	return s.repo.deleteProduct(id)
}
