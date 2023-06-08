package itemsrv

import "restserver/repositories/itemrepository"

type itemService struct {
	itemRepository itemrepository.ItemRepository
}

func NewItemService(itemRepository itemrepository.ItemRepository) ItemService {
	return &itemService{itemRepository}
}

func (obj itemService) AllItems() (items []Item, err error) {
	itemResult, err := obj.itemRepository.AllItems()

	if err != nil {
		return
	}

	for _, item := range itemResult {

		temp := Item{
			ID:            item.ID,
			Title:         item.Title,
			Price:         item.Price,
			OriginalPrice: item.OriginalPrice,
		}

		items = append(items, temp)
	}

	return
}
