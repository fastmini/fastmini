// Package global
// @Description:
// @Author AN 2023-12-06 23:21:03
package global

import (
	"fiber/model"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

const Version = "1.0.0"

// Redis 默认redis连接池
var Redis redis.Conn

// DB 数据库
var DB *gorm.DB

// SLog 系统日志
var SLog *log.Logger

// BLog 系统日志
var BLog *log.Entry

// LogFile 日志文件
var LogFile *os.File

// ES ES客户端
var ES *elasticsearch.Client

type AuthUserPayload struct {
	UserId int64
}

var AuthUser *AuthUserPayload

func SetAuthUser(user *model.SysUser) {
	AuthUser = &AuthUserPayload{UserId: user.Id}
}

func GetAuthUser() int64 {
	return AuthUser.UserId
}
