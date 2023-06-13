package itemsrv

import "restserver/repositories/itemrepository"

type itemService struct {
	itemRepository itemrepository.ItemRepository
}

func NewItemService(itemRepository itemrepository.ItemRepository) ItemService {
	return &itemService{itemRepository}
}

func (obj itemService) AllItems() (items []Item, err error) {
	itemData, err := obj.itemRepository.AllItems()

	if err != nil {
		return
	}

	for _, item := range itemData {

		temp := Item{
			ID:            item.ID,
			Title:         item.Title,
			Slug:          item.Slug,
			Price:         item.Price,
			OriginalPrice: item.OriginalPrice,
		}

		items = append(items, temp)
	}

	return
}

func (obj itemService) GetItemBySlug(slug string) (item Item, err error) {

	itemData, err := obj.itemRepository.GetItemBySlug(slug)

	if err != nil {
		return
	}

	item.ID = itemData.ID
	item.Title = itemData.Title
	item.Slug = itemData.Slug
	item.Price = itemData.Price
	item.OriginalPrice = itemData.OriginalPrice

	return
}
