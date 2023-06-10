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
func GetDatabaseGorm(dbName string) *gorm.DB {
	// 读取数据库配置
	dbKey := "mysql." + dbName
	dsnFormat := viper.GetString("mysql.format.gorm")
	user := viper.GetString(dbKey + ".user")
	pass := viper.GetString(dbKey + ".pass")
	host := viper.GetString(dbKey + ".host")
	port := viper.GetString(dbKey + ".port")
	database := viper.GetString(dbKey + ".database")
	charset := viper.GetString(dbKey + ".charset")

	// 组合dsn配置
	dsn := fmt.Sprintf(dsnFormat, user, pass, host, port, database, charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Error("Mysql Connect error")
	}
	return db
}

// 参见文档 https://xorm.io/zh/docs/chapter-01/1.engine/
func GetDatabaseXorm(dbName string) *xorm.EngineGroup {
	// 读取数据库配置
	dbKey := "mysql." + dbName
	dsnFormat := viper.GetString("mysql.format.xorm")
	user := viper.GetString(dbKey + ".user")
	pass := viper.GetString(dbKey + ".pass")
	host := viper.GetString(dbKey + ".host")
	port := viper.GetString(dbKey + ".port")
	database := viper.GetString(dbKey + ".database")
	charset := viper.GetString(dbKey + ".charset")

	// 组合dsn配置
	dsn := fmt.Sprintf(dsnFormat, user, pass, host, port, database, charset)

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
