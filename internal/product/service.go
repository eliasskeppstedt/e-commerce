package product

type productService interface {
	getByProductID(id int) (Product, error)
	getAll() ([]Product, error)
	registerProduct(product Product) error
	deleteProduct(id int) error
}

type productServiceImp struct {
	repo productRepository
}

func NewProductServiceImp(r productRepository) *productServiceImp {
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

func (s *productServiceImp) updateProduct(id int, stock int, price float64) error {
	return s.repo.updateProduct(id, stock, price)
}
