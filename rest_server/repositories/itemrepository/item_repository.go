package itemrepository

import (
	"database/sql"

	"gorm.io/gorm"
)

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db}
}

func (obj itemRepository) AllItems() (items []item, err error) {

	tx := obj.db.Limit(10).Order("updated_at DESC").Find(&items)

	if tx.Error != nil {
		return items, tx.Error
	}

	return items, err
}

func (obj itemRepository) GetItemBySlug(slug string) (item item, err error) {

	tx := obj.db.Where("slug = @slug", sql.Named("slug", slug)).Find(&item)

	if tx.Error != nil {
		return item, tx.Error
	}

	return
}
