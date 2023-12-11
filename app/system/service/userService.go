package service

import (
	"fiber/app/system/dao"
	"fiber/app/system/request"
	"fiber/app/system/response"
	"fiber/config"
	businessError "fiber/error"
	"fiber/global"
	"fiber/utils"
	"golang.org/x/crypto/bcrypt"
)

func Login(request *request.LoginRequest) response.LoginResponse {
	// 查询用户的数据
	user := dao.FindUserByMobile(request.Mobile)
	if user.Id == 0 {
		panic(businessError.New(businessError.BAD_REQUEST, "手机号或密码不正确[nf]"))
	}
	// 验证用户密码是否正确
	if ComparePassword(user.Password, request.Password) {
		token := utils.CreateToken(user.Id, user.PasswordVersion)
		return response.LoginResponse{Token: token, ExpireTime: int64(config.TTL)}
	}
	panic(businessError.New(businessError.BAD_REQUEST, "手机号或密码不正确"))
}

func Profile() response.UserResponse {
	userId := global.GetAuthUser()
	user := dao.FindUserById(userId)
	if user.Id != 0 {
		return response.UserResponse{
			ID:              user.Id,
			Mobile:          user.Mobile,
			PasswordVersion: user.PasswordVersion,
			NickName:        user.NickName,
		}
	}
	panic(businessError.New(businessError.BAD_REQUEST, "用户不存在"))
}

func HashPassword(password string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPwd), nil
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
