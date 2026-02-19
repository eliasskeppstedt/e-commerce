package product

type productService interface {
	getByProductID(id int) (product, error)
	registerProduct(product product) error
	deleteProduct(id int) error
}

type productServiceImp struct {
	repo productRepository
}

func NewProductServiceImp(r productRepository) *productServiceImp {
	return &productServiceImp{repo: r}
}

func (s *productServiceImp) getByProductID(id int) (product, error) {
	return s.repo.getByProductID(id)
}

func (s *productServiceImp) registerProduct(product product) error {
	return s.repo.registerProduct(product)
}

func (s *productServiceImp) deleteProduct(id int) error {
	return s.repo.deleteProduct(id)
}
