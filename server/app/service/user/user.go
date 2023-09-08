package user

import (
	"rango/models/origin"
	"rango/tools"

	"github.com/doug-martin/goqu/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

type UserService struct{}

// 登录
func (the *UserService) Login(phone string) (string, error) {
	data, err := origin.DeUserTable().GetOne(goqu.Ex{"phone": phone}, map[string]string{})
	if err != nil {
		return "", err
	}
	uid := cast.ToInt64(data["id"])
	if uid <= 0 {
		return the.Register(phone)
	}

	// 下发token
	token, err := tools.NewJWT().CreateToken(tools.BaseClaims{
		Uid:      uid,
		Phone:    phone,
		NickName: "大白兔",
	})

	if err != nil {
		logrus.Error("CreateToken :", err.Error())
		return "", err
	}

	logrus.Debug(token)
	return token, nil
}

// 注册
func (the *UserService) Register(phone string) (string, error) {
	uid, rowsAffected, err := origin.DeUserTable().Insert(map[string]interface{}{
		"phone": phone,
	})
	if err != nil {
		return "", err
	}
	if rowsAffected <= 0 {
		return "", err
	}

	// 权限实例
	enforcer, err := tools.GetEnforcer()
	if err != nil {
		return "", err
	}
	// 配置权限
	hasPolicy := enforcer.HasPolicy(cast.ToString(uid), "/user/login", "POST")
	if !hasPolicy {
		enforcer.AddPolicy(cast.ToString(uid), "/user/login", "POST")
	}

	// 下发token
	token, err := tools.NewJWT().CreateToken(tools.BaseClaims{
		Uid:      uid,
		Phone:    phone,
		NickName: "大白兔",
	})

	if err != nil {
		logrus.Error("CreateToken :", err.Error())
		return "", err
	}

	logrus.Debug(token)
	return token, nil
}
