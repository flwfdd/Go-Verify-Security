package service

import (
	"awesomeProject/dao"
	"awesomeProject/model"
	"awesomeProject/parameter"
	"awesomeProject/tool"
	"fmt"
	"math/rand"
	"time"
)

// UserService 和用户相关的服务
type UserService struct {
}

// SendCode 发送验证码
func (s *UserService) SendCode(phone string) (bool, string) {
	// 生成随机验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	// 数据库记录验证码
	smsCode := model.SmsCode{Phone: phone, Code: code, CreateTime: time.Now().Unix()}
	userDao := dao.UserDao{Orm: tool.DbEngine}
	result := userDao.InsertCode(&smsCode)

	return result > 0, code
}

// LoginByPhone 手机登录
func (s *UserService) LoginByPhone(loginParam parameter.LoginByPhoneRequest) *model.User {
	// 查询验证码表里是否有此记录
	userDao := dao.UserDao{Orm: tool.DbEngine}
	smsCode := userDao.ValidateSmsCode(loginParam.PhoneNumber, loginParam.VerifyCode)
	if smsCode.Id == 0 {
		return nil
	}
	// 查询此用户是否已经注册，因为有没注册但是有验证码的情况
	result1 := userDao.QueryByPhone(loginParam.PhoneNumber)
	if result1 == nil {
		return nil
	}
	result2 := userDao.UpdateLoggedDeviceNumber(1, result1.Username)
	if result2 == nil {
		return nil
	}
	return result1
}

// LoginByPassword 密码登录
func (s *UserService) LoginByPassword(loginParam parameter.LoginByPasswordRequest) *model.User {
	userDao := dao.UserDao{Orm: tool.DbEngine}
	result1 := userDao.ValidatePassword(loginParam.Username, loginParam.Password)
	if result1 == nil {
		return nil
	}
	result2 := userDao.UpdateLoggedDeviceNumber(1, result1.Username)
	if result2 == nil {
		return nil
	}
	return result1
}

// Register 注册
func (s *UserService) Register(registerParam parameter.RegisterRequest) *model.User {
	userDao := dao.UserDao{Orm: tool.DbEngine}
	result, user := userDao.InsertUser(registerParam.Username, registerParam.PhoneNumber, registerParam.Password)
	if result == 0 {
		return nil
	}
	return user
}

//func (s *UserService) LogoutUser(logoutParam parameter.LogoutRequest) *model.User {
//	userDao := dao.UserDao{}
//	userDao.UpdateLoggedDeviceNumber(-1, logoutParam.)
//}
