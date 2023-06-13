package itemsrv

type Item struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Slug          string  `json:"slug"`
	Price         float32 `json:"price"`
	OriginalPrice float32 `json:"originalPrice"`
}

type ItemService interface {
	AllItems() (items []Item, err error)
	GetItemBySlug(slug string) (item Item, err error)
}
