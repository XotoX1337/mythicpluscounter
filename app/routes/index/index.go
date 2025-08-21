package index

import (
	"github.com/XotoX1337/mythicpluscounter/app/raiderio"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	rio, err := raiderio.NewClient()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	runs, err := rio.Runs.List(&raiderio.RunsListOptions{
		Season:      "season-tww-3",
		CharacterId: "236673948",
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return c.Render("app/views/index", fiber.Map{
		"Runs": runs,
	})
}
