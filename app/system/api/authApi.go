package api

import (
	"fiber/app/system/request"
	"fiber/app/system/service"
	businessError "fiber/error"
	"fiber/resultVo"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	p := new(request.LoginRequest)
	if err := c.BodyParser(p); err != nil {
		return c.JSON(resultVo.Fail(businessError.New(businessError.BAD_REQUEST), c))
	}
	validate := validator.New()
	if err := validate.Struct(p); err != nil {
		return c.JSON(resultVo.Fail(businessError.New(businessError.BAD_REQUEST, err.Error()), c))
	}
	user := service.Login(p)
	return c.JSON(resultVo.Success(user, c))
}

func GetBingBackgroundImage(c *fiber.Ctx) error {
	data := make(map[string]interface{})
	data["url"] = ""
	return c.JSON(resultVo.Success(data, c))
}
