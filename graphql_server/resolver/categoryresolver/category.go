package categoryresolver

import "github.com/graphql-go/graphql"

type Category struct {
	ID       int
	ParentID int
	Name     string
	Slug     string
	Image    string
	Active   bool
}

type CategoryResolver interface {
	AllCategories() *graphql.Field
}
