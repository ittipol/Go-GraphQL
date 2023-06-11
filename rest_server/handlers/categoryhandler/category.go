package categoryhandler

import "github.com/gofiber/fiber/v2"

type CategoryHandler interface {
	AllCategories(c *fiber.Ctx) error
}
