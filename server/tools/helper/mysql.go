package helper

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
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

// mysql helper
// 查询出来的byte数据转int64
func ToInt64(data map[string]interface{}, key string) int64 {
	byteArr, ok := data[key].([]byte)
	if !ok {
		return 0
	}

	var value int64
	err := json.Unmarshal(byteArr, &value)
	if err != nil {
		logrus.Error("json Unmarshal err", err.Error())
	}
	return value
}
