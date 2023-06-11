package categoryrepository

import "gorm.io/gorm"

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (obj categoryRepository) AllCategories() (categories []category, err error) {

	if err = obj.db.Find(&categories).Error; err != nil {
		return
	}

	return
}
