package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"xorm.io/xorm"
)

// 参见文档 https://gorm.io/zh_CN/docs/
func GetDatabaseGorm(database string) *gorm.DB {
	// 读取数据库配置
	viper.GetString("mysql.dsn." + database)
	// 集群配置
	dsn := "root:secret20@(192.168.5.5:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Error("Mysql Connect error")
	}
	return db
}

// 参见文档 https://xorm.io/zh/docs/chapter-01/1.engine/
func GetDatabaseXorm(database string) *xorm.EngineGroup {
	// 读取数据库配置
	viper.GetString("mysql.dsn." + database)
	// 集群配置
	conns := []string{
		// viper.GetString("mysql.dsn"),
		"root:secret20@tcp(192.168.5.5:3306)/golang?charset=utf8mb4",
	}

	db, err := xorm.NewEngineGroup("mysql", conns)

	if err != nil {
		logrus.Error("Mysql Connect error")
	}
	return db
}
