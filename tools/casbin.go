package tools

import (
	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	"github.com/sirupsen/logrus"
)

// 获取实例
func GetEnforcer() (enforcer *casbin.Enforcer, err error) {
	adapter, err := xormadapter.NewAdapter("mysql", "root:123456@tcp(6.tcp.ngrok.io:14527)/casbin", true)
	if err != nil {
		logrus.Fatalf("casbin error, detail: %s", err)
		return
	}

	// 从数据库中加载信息
	enforcer, err = casbin.NewEnforcer("./config/casbin/model.conf", adapter)
	if err != nil {
		logrus.Fatalf("casbin error, detail: %s", err)
		return
	}

	return
}
