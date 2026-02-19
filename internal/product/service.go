package product

type productService interface {
	getByProductID(product_id string) (product, error)
	registerProduct(product product) error
	deleteProduct(product product) error
}

type productServiceImp struct {
	repo productRepository
}

func NewProductServiceImp(r productRepository) *productServiceImp {
	return &productServiceImp{repo: r}
}

func (s *productServiceImp) getByProductID(product_id int) (product, error) {
	return s.repo.getByProductID(product_id)
}

func (s *productServiceImp) registerProduct(product product) error {
	return s.repo.registerProduct(product)
}

func (s *productServiceImp) deleteProduct(product_id int) error {
	return s.repo.deleteProduct(product_id)
}
