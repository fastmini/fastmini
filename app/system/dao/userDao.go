package dao

import (
	"fiber/global"
	"fiber/model"
)

func FindUserByMobile(mobile string) *model.SysUser {
	sysUser := new(model.SysUser)
	global.DB.Where("mobile = ?", mobile).First(sysUser)
	return sysUser
}

func FindUserById(id int64) *model.SysUser {
	sysUser := new(model.SysUser)
	global.DB.Where("id = ?", id).First(sysUser)
	return sysUser
}
