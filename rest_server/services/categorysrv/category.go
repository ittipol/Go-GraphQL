package categorysrv

type Category struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parentId"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Image    string `json:"image"`
	Active   bool   `json:"active"`
}

type CategoryService interface {
	AllCategories() (Categories []Category, err error)
}
