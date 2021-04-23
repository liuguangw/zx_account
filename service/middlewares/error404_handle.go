package middlewares

import (
	_ "embed"
	"github.com/gofiber/fiber/v2"
)

//go:embed pages/404.html
var error404PageHtml string

//Error404Handle 处理HTTP404
func Error404Handle(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")
	return c.Status(404).SendString(error404PageHtml)
}
