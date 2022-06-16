package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"xorm.io/xorm"
)

var Db *xorm.EngineGroup

func init() {
	var err error
	// 集群配置
	conns := []string{
		"root:secret20@tcp(192.168.5.5:3306)/golang?charset=utf8mb4",
		"root:secret20@tcp(192.168.5.5:3306)/golang?charset=utf8mb4",
	}

	Db, err = xorm.NewEngineGroup("mysql", conns)

	if err != nil {
		logrus.Error("Mysql Connect error")
	}
}

// 参见文档 https://gorm.io/zh_CN/docs/
func Gorm() *gorm.DB {
	var db *gorm.DB
	var err error
	// 集群配置
	dsn := "root:secret20@(192.168.5.5:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Error("Mysql Connect error")
	}
	return db
}

// 参见文档 https://xorm.io/zh/docs/chapter-01/1.engine/
func Xorm() *xorm.EngineGroup {
	var db *xorm.EngineGroup
	var err error
	// 集群配置
	conns := []string{
		viper.GetString("mysql.dsn"),
		viper.GetString("mysql.dsn"),
	}

	db, err = xorm.NewEngineGroup("mysql", conns)

	if err != nil {
		logrus.Error("Mysql Connect error")
	}
	return db
}
