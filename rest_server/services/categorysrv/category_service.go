package categorysrv

import "restserver/repositories/categoryrepository"

type categoryService struct {
	categoryRepository categoryrepository.CategoryRepository
}

func NewCategoryService(categoryRepository categoryrepository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository}
}

func (obj categoryService) AllCategories() (Categories []Category, err error) {

	categoryData, err := obj.categoryRepository.AllCategories()

	if err != nil {
		return
	}

	for _, category := range categoryData {

		category := Category{
			ID:       category.ID,
			ParentID: category.ParentID,
			Name:     category.Name,
			Slug:     category.Slug,
			Image:    category.Image,
			Active:   category.Active,
		}

		Categories = append(Categories, category)
	}

	return
}
