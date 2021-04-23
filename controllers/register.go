package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/zx_account/service"
	"github.com/liuguangw/zx_account/service/request"
)

func Register(c *fiber.Ctx) error {
	requestInfo := new(request.RegisterAccount)
	if err := c.BodyParser(requestInfo); err != nil {
		return err
	}
	if err := service.RegisterAccount(requestInfo.Username, requestInfo.Password, requestInfo.Email); err != nil {
		return c.JSON(fiber.Map{
			"code":    500,
			"message": err.Error(),
			"data":    nil,
		})
	}
	return c.JSON(fiber.Map{
		"code":    0,
		"message": "",
		"data":    nil,
	})
}
