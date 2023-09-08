package sms

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"rango/tools/rds"
	"rango/tools/sdk"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type SmsService struct{}

type Template struct { // 具体格式需要具体设计
	Code string `json:"code"` // name
}

// 发送短信
func (the *SmsService) Send(phone string) (string, error) {
	// 随机数
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	jsonTpl, err := json.Marshal(&Template{
		Code: code,
	})
	if err != nil {
		return "", err
	}

	// 线上真实发送
	if viper.GetString("env.name") == "online" {
		// 设置缓存
		redisKey := "phone:code:" + phone
		err = rds.GetRedis().SetEX(context.Background(), redisKey, code, time.Duration(180)*time.Second).Err()
		if err != nil {
			logrus.Error("rdb.SetEX :", err.Error())
			return "", err
		}

		// 发送短信
		sdk.SendSms(phone, viper.GetString("AiliyunSms.LoginCode"), string(jsonTpl))
		if err != nil {
			logrus.Error("sdk.SendSms :", err.Error())
			return "", err
		}
	} else {
		code = "123456"
	}

	logrus.Debug("send sms code :", code)
	return code, nil
}

// 检查短信验证码是否正确
func (the *SmsService) CheckCode(phone string, code string) bool {
	codeString := "123456"
	if viper.GetString("env.name") == "online" {
		// 获取缓存
		redisKey := "phone:code:" + phone
		codeString, err := rds.GetRedis().Get(context.Background(), redisKey).Result()
		if err != nil {
			logrus.Error("rdb.Get :", err.Error())
			return false
		}
		if code == codeString {
			return true
		}
	}

	if code == codeString {
		return true
	}
	return false
}
