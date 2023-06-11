package categoryrepository

type category struct {
	ID       int
	ParentID int `gorm:"column:parent_id"`
	Name     string
	Slug     string
	Image    string
	Active   bool
}

type CategoryRepository interface {
	AllCategories() (categories []category, err error)
}
