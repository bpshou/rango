package models

import (
	"rango/tools/helper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"xorm.io/xorm"
)

// 参见文档 https://gorm.io/zh_CN/docs/
func GetDatabaseGorm(database string) *gorm.DB {
	// dsn配置
	dsn := helper.GetMysqlGormDSN(database)
	// 连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("Mysql Connect error")
	}
	return db
}

// 参见文档 https://xorm.io/zh/docs/chapter-01/1.engine/
func GetDatabaseXorm(database string) *xorm.EngineGroup {
	// dsn配置
	dsn := helper.GetMysqlXormDSN(database)
	// 集群配置
	db, err := xorm.NewEngineGroup("mysql", []string{dsn, dsn})
	if err != nil {
		logrus.Error("Mysql Connect error")
	}
	return db
}
