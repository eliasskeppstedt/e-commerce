package category

type categoryService interface {
	getAll() ([]Category, error)
}

type categoryServiceImp struct {
	repo categoryRepository
}

func NewCategoryServiceImp(r categoryRepository) *categoryServiceImp {
	return &categoryServiceImp{repo: r}
}

func (s *categoryServiceImp) getAll() ([]Category, error) {
	return s.repo.getAll()
}

func (s *categoryServiceImp) createCategory(category Category) error {
	return s.repo.createCategory(category)
}

func (s *categoryServiceImp) deleteCategory(id int) error {
	return s.repo.deleteCategory(id)
}
