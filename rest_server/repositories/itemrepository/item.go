package itemrepository

// https://gorm.io/docs/models.html
// VAR_NAME string `gorm:"primarykey;size:16"`
type item struct {
	ID            int
	Title         string `gorm:"size:255"`
	Price         float32
	OriginalPrice float32 `gorm:"column:original_price"`
}

type ItemRepository interface {
	AllItems() (items []item, err error)
}
