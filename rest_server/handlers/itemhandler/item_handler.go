package itemhandler

import (
	"restserver/services/itemsrv"

	"github.com/gofiber/fiber/v2"
)

type itemHandler struct {
	itemService itemsrv.ItemService
}

func NewItemHandler(itemService itemsrv.ItemService) ItemHandler {
	return &itemHandler{itemService}
}

func (obj itemHandler) AllItems(c *fiber.Ctx) error {

	items, err := obj.itemService.AllItems()

	if err != nil {
		return fiber.ErrInternalServerError
	}

	c.Status(fiber.StatusOK)
	return c.JSON(items)
}

func (obj itemHandler) GetItemBySlug(c *fiber.Ctx) error {

	slug := c.Params("slug")

	item, err := obj.itemService.GetItemBySlug(slug)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	c.Status(fiber.StatusOK)
	return c.JSON(item)
}
