package itemrepository

import "gorm.io/gorm"

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db}
}

func (obj itemRepository) AllItems() (items []Item, err error) {

	tx := obj.db.Limit(10).Order("updated_at DESC").Find(&items)

	if tx.Error != nil {
		return items, tx.Error
	}

	return items, err
}
