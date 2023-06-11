package categoryhandler

import (
	"restserver/services/categorysrv"

	"github.com/gofiber/fiber/v2"
)

type categoryhandler struct {
	categoryService categorysrv.CategoryService
}

func NewCategoryHandler(categoryService categorysrv.CategoryService) CategoryHandler {
	return &categoryhandler{categoryService}
}

func (obj categoryhandler) AllCategories(c *fiber.Ctx) error {

	categories, err := obj.categoryService.AllCategories()

	if err != nil {
		return fiber.ErrInternalServerError
	}

	c.Status(fiber.StatusOK)
	return c.JSON(categories)
}
