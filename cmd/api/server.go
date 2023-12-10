package api

import (
	"fiber/config"
	businessError "fiber/error"
	"fiber/global"
	"fiber/libraray/core"
	"fiber/libraray/database"
	"fiber/libraray/elastic"
	"fiber/libraray/grtm"
	"fiber/libraray/redis"
	"fiber/libraray/xxljob"
	"fiber/middleware"
	"fiber/resultVo"
	"fiber/router"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"log"
)

var (
	StartCmd = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "goapp server",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func setup() {
}

func run() error {
	defer global.CloseGlobal()
	core.Tips()
	// 启动配置
	addr := fmt.Sprintf(":%s", config.Config("PORT"))
	initConfig := fiber.Config{
		ServerHeader:          "GoApp",
		ReduceMemoryUsage:     true,
		ColorScheme:           fiber.DefaultColors,
		DisableStartupMessage: true,
		EnablePrintRoutes:     false,
		AppName:               config.Config("APP_NAME"),
		// 业务异常返回
		ErrorHandler: func(ctx *fiber.Ctx, e error) error {
			if err, ok := e.(*businessError.Err); ok {
				// 业务异常
				return ctx.JSON(resultVo.Fail(err, ctx), fiber.MIMEApplicationJSONCharsetUTF8)
			} else {
				// 系统异常
				return ctx.JSON(resultVo.Fail(businessError.New(businessError.SERVER_ERROR), ctx), fiber.MIMEApplicationJSONCharsetUTF8)
			}
		},
	}
	// 配置初始化
	app := fiber.New(initConfig)
	// 中间件初始化
	middleware.InitMiddleware(app)
	// redis
	redis.InitRedis()
	// 连接ds数据库
	database.ConnectDB()
	// elastic
	elastic.ConnectES()
	// xxl
	xxljob.ConnectXxlJob(app)
	// 初始化路由
	router.AppRouter(app)
	// 初始化线程池
	grtm.InitCoPool()
	// 启动消息
	core.StartupMessage(addr)
	// 启动服务
	log.Fatal(app.Listen(addr))
	return nil
}
