package user

import (
	"rango/tools"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

type UserService struct{}

// 登录
func (the *UserService) Login(phone string) (string, error) {
	jwtTools := tools.NewJWT()

	// 下发token
	token, err := jwtTools.CreateToken(tools.BaseClaims{
		Uid:      cast.ToInt64(phone),
		Phone:    phone,
		NickName: "大白兔",
	})

	if err != nil {
		logrus.Error("jwtTools.CreateToken :", err.Error())
		return "", err
	}

	logrus.Debug(token)
	return token, nil
}

// 注册
func (the *UserService) Register(phone string) (string, error) {
	jwtTools := tools.NewJWT()
	enforcer, err := tools.GetEnforcer()
	if err != nil {
		return "", err
	}

	// 配置权限
	hasPolicy := enforcer.HasPolicy(phone, "/user/login", "GET")
	if !hasPolicy {
		enforcer.AddPolicy(phone, "/user/login", "GET")
	}

	// 下发token
	token, err := jwtTools.CreateToken(tools.BaseClaims{
		Uid:      cast.ToInt64(phone),
		Phone:    phone,
		NickName: "大白兔",
	})

	if err != nil {
		logrus.Error("jwtTools.CreateToken :", err.Error())
		return "", err
	}

	logrus.Debug(token)
	return token, nil
}
