package controllers

import (
	_ "embed"
	"github.com/gofiber/fiber/v2"
)

//go:embed pages/register.html
var registerPageHtml string

func ShowRegister(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")
	return c.SendString(registerPageHtml)
}
