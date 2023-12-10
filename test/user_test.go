package test

import (
	"fiber/app/system/service"
	"fmt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	if pwd, err := service.HashPassword("123456"); err == nil {
		fmt.Printf("生成的密码: %s\n", pwd)
	}
}
