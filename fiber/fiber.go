package fiber

import (
	"backend/core"

	"github.com/gofiber/fiber/v2"
)

var (
	log = core.GetLogger()
)

func Start() {
	fiberConfig := fiber.Config{}
	app := fiber.New(fiberConfig)
	middleware(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("node backend..")
	})

	app.Listen(":3000")
}
