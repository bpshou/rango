package models

import (
	"fmt"

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
	mysqlConfig := viper.GetStringMap("mysql." + database)
	dsnFormat := viper.GetString("model_gorm.dsn_format")

	// 组合dsn配置
	dsn := fmt.Sprintf(dsnFormat, mysqlConfig["user"], mysqlConfig["pass"], mysqlConfig["host"], mysqlConfig["port"], mysqlConfig["database"], mysqlConfig["charset"])

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Error("Mysql Connect error")
	}
	return db
}

// 参见文档 https://xorm.io/zh/docs/chapter-01/1.engine/
func GetDatabaseXorm(database string) *xorm.EngineGroup {
	// 读取数据库配置
	mysqlConfig := viper.GetStringMap("mysql." + database)
	dsnFormat := viper.GetString("model_xorm.dsn_format")

	// 组合dsn配置
	dsn := fmt.Sprintf(dsnFormat, mysqlConfig["user"], mysqlConfig["pass"], mysqlConfig["host"], mysqlConfig["port"], mysqlConfig["database"], mysqlConfig["charset"])

	// 集群配置
	conns := []string{
		dsn,
		dsn,
	}

	db, err := xorm.NewEngineGroup("mysql", conns)

	if err != nil {
		logrus.Error("Mysql Connect error")
	}
	return db
}
