package sdk

import (
	"errors"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (client *dysmsapi.Client, err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	// 实例化客户端
	client, err = dysmsapi.NewClient(config)
	return
}

func SendSms(phone string, templateCode string, templateParam string) (err error) {
	if phone == "" {
		err = errors.New("手机号为空")
		return
	}
	if templateCode == "" {
		err = errors.New("模板code为空")
		return
	}
	if templateParam == "" {
		err = errors.New("模板内容为空")
		return
	}

	// 密钥
	SignName := viper.GetString("AiliyunSms.SignName")
	accessKeyId := viper.GetString("AiliyunSms.AccessId")
	accessKeySecret := viper.GetString("AiliyunSms.AccessSecret")

	// 客户端
	client, err := CreateClient(&accessKeyId, &accessKeySecret)
	if err != nil {
		logrus.Errorf("aliyun sms CreateClient err:%s", err.Error())
		return
	}

	sendSmsRequest := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  &phone,
		SignName:      &SignName,
		TemplateCode:  &templateCode,
		TemplateParam: &templateParam,
	}

	response, err := client.SendSms(sendSmsRequest)
	if err != nil {
		logrus.Errorf("aliyun sms err:%s", err.Error())
		return
	}
	logrus.Infof("aliyun sms response:%s", response)
	return
}
