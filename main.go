/*
 * @Author: Ali2vu <751815097@qq.com>
 * @Date: 2021-08-09 22:31:58
 * @LastEditors: Ali2vu
 * @LastEditTime: 2021-08-10 01:34:30
 */
package main

import (
	"fastmini/config"
	"fastmini/global"
	ZAP_LOGGER "fastmini/logger"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	prefork, _ := strconv.ParseBool(config.Config("APP_PREFORK", "false"))
	app := fiber.New(fiber.Config{
		Prefork:       prefork,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fastmini",
		AppName:       "Fastmini v1.0.0",
	})

	ZAP_LOGGER.InitLogger()
	defer global.FM_LOG.Sync()

	app.Use(requestid.New())
	app.Use(requestid.New(requestid.Config{
		Header:    fiber.HeaderXRequestID,
		Generator: func() string {
			//global.FM_LOG.Debugf("uuid: %v", utils.UUIDv4())
			return fmt.Sprintf("req.%v", time.Now().UnixNano())
		},
	}))
	app.Use(logResponseBody)

	format := "[${time}] ${pid} ${locals:requestid} ${status} - ${latency} - ${method} ${path}​"
	timeFormat := "2006-01-02 15:04:05"
	timeZone := "Asia/Shanghai"

	if config.Config("APP_LOGGER", "file") == "file" {
		file, err := os.OpenFile("./request.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer file.Close()
		app.Use(logger.New(logger.Config{
			Format:     format,
			TimeFormat: timeFormat,
			TimeZone:   timeZone,
			Output:     file,
		}))
	}

	if config.Config("APP_LOGGER", "file") == "stdout" {
		app.Use(logger.New(logger.Config{
			Format:     format,
			TimeFormat: timeFormat,
			TimeZone:   timeZone,
			Output:     os.Stdout,
		}))
	}

	app.Get("/", func(c *fiber.Ctx) error {
		c.Response().Header.Set("Server", "123")
		global.FM_LOG.Debugf("时间: %d", time.Now().Unix())
		global.FM_LOG.Debugf("请求URL: %s", c.BaseURL())
		global.FM_LOG.Debugf("请求requesttid: %s", c.Locals("requestid"))
		return c.SendString("Hello, World 👋!")
	})

	app.Static("/public", "./public", fiber.Static{
		Compress: true,
		Browse: true,
	})

	app.Get("/download", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/plain; charset=utf-8")
		return c.SendFile("./public/update.json", false)
	})

	log.Fatal(app.Listen(":3002"))
}

func logResponseBody(c *fiber.Ctx) error {
	c.Next()
	//c.Response().SetBodyString("1111")
	global.FM_LOG.Debugf("body->%s", c.Response().Header.String())
	return nil
}