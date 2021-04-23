package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/zx_account/controllers"
)

func LoadCommonRoutes(app *fiber.App) {
	app.Get("/", controllers.ShowRegister)
}
