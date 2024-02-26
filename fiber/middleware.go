package fiber

import (
	"backend/settings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	appSettings = settings.Get()
)

func middleware(app *fiber.App) {
	rcc := recover.Config{
		Next:             nil,
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			log.Error(e)
		},
	}

	rc := recover.New(rcc)
	app.Use(rc)
	app.Use(cors.New())

}
