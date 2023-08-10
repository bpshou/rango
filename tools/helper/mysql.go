package helper

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetMysqlGormDSN(database string) string {
	// 读取数据库配置
	dsnFormat := viper.GetString("mysql.format.gorm")
	user := viper.GetString("mysql.origin.user")
	pass := viper.GetString("mysql.origin.pass")
	host := viper.GetString("mysql.origin.host")
	port := viper.GetString("mysql.origin.port")
	charset := viper.GetString("mysql.origin.charset")

	// 组合dsn配置
	return fmt.Sprintf(dsnFormat, user, pass, host, port, database, charset)
}

func GetMysqlXormDSN(database string) string {
	// 读取数据库配置
	dsnFormat := viper.GetString("mysql.format.xorm")
	user := viper.GetString("mysql.origin.user")
	pass := viper.GetString("mysql.origin.pass")
	host := viper.GetString("mysql.origin.host")
	port := viper.GetString("mysql.origin.port")
	charset := viper.GetString("mysql.origin.charset")

	// 组合dsn配置
	return fmt.Sprintf(dsnFormat, user, pass, host, port, database, charset)
}
