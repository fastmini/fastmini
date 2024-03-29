// Package router
// @Description:
// @Author AN 2023-12-06 23:16:19
package router

import (
	"fiber/app/system/api"
	taskApi "fiber/app/task/api"
	businessError "fiber/error"
	"fiber/global"
	"fiber/middleware"
	"fiber/resultVo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"time"
)

func AppRouter(app *fiber.App) {
	app.Get("/metrics", monitor.New(monitor.Config{Title: "GoApp Monitor", Refresh: 2 * time.Second}))
	app.Get("/_startup", func(ctx *fiber.Ctx) error {
		global.BLog.Infof("hello")
		return ctx.JSON(resultVo.Success("ok", ctx), fiber.MIMEApplicationJSONCharsetUTF8)
	})
	app.Get("/_health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(resultVo.Success("success", ctx), fiber.MIMEApplicationJSONCharsetUTF8)
	})
	app.Post("/task/testCheck", taskApi.TestCheck)
	app.Post("/user/login", api.Login)
	app.Get("/system/getBingBackgroundImage", api.GetBingBackgroundImage)
	// 需要登录鉴权的路由
	apiRoute := app.Group("", middleware.AuthMiddleware())
	apiRoute.Get("/user/profile", api.Profile)
	// 其他
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(resultVo.Success(nil, ctx), fiber.MIMEApplicationJSONCharsetUTF8)
	})
	// 404返回
	app.Use(func(c *fiber.Ctx) error {
		return c.JSON(resultVo.Fail(businessError.New(businessError.NOT_FOUND), c), fiber.MIMEApplicationJSONCharsetUTF8)
	})
	// 这个后面不要写
}
