package itemsrv

type Item struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Price         float32 `json:"price"`
	OriginalPrice float32 `json:"originalPrice"`
}

type ItemService interface {
	AllItems() (items []Item, err error)
}
