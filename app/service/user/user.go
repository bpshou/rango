package user

import (
	"rango/tools"

	"github.com/sirupsen/logrus"
)

type UserService struct{}

// 登录
func (the *UserService) Login() (string, error) {
	jwtTools := tools.NewJWT()
	token, err := jwtTools.CreateToken(tools.BaseClaims{
		Uid:      1,
		Phone:    "13199996666",
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
func (the *UserService) Register() (string, error) {
	jwtTools := tools.NewJWT()
	token, err := jwtTools.CreateToken(tools.BaseClaims{
		Uid:      1,
		Phone:    "13199996666",
		NickName: "大白兔",
	})

	if err != nil {
		logrus.Error("jwtTools.CreateToken :", err.Error())
		return "", err
	}

	logrus.Debug(token)
	return token, nil
}
