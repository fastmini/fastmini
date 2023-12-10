// Package main
// @Description:
// @Author AN 2023-12-06 23:16:36
package main

import (
	"fiber/cmd"
	"github.com/joho/godotenv"
)

func main() {
	// 配置文件加载
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}
	cmd.Execute()
}
