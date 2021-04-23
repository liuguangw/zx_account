package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/zx_account/controllers"
)

func LoadAPIRoutes(app *fiber.App) {
	apiGroup := app.Group("/api")
	apiGroup.Post("/auth/register", controllers.Register)
}
