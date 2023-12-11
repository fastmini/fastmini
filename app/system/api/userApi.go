package api

import (
	"fiber/app/system/service"
	"fiber/resultVo"
	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) {

}

func RefreshToken(c *fiber.Ctx) {

}

func Profile(c *fiber.Ctx) error {
	user := service.Profile()
	return c.JSON(resultVo.Success(user, c))
}
