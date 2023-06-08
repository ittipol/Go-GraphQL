package itemhandler

import "github.com/gofiber/fiber/v2"

type ItemHandler interface {
	AllItems(c *fiber.Ctx) error
}
