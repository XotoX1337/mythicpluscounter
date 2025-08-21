package index

import "github.com/gofiber/fiber/v2"

func Index(c *fiber.Ctx) error {

	return c.Render("app/views/index", fiber.Map{})
}
