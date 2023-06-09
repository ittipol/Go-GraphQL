package itemresolver

import "github.com/graphql-go/graphql"

type itemRequest struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

type item struct {
	Id            int
	Title         string
	Price         float32
	OriginalPrice float32
}

type ItemResolver interface {
	AllItems() *graphql.Field
	AddItem() *graphql.Field
}
