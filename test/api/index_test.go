package api

import (
	"encoding/json"
	"os"
	"rango/tools/load"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestIndex(t *testing.T) {
	// 当前文件路径
	path, _ := os.Getwd()
	load.LoadViper(path + "/../../")

	host := viper.GetString("rango.host")
	url := host + "/api/ping"

	// 发送请求
	status, response, err := fasthttp.Get(nil, url)
	if err != nil {
		t.Errorf("Request API [%v] error :%v", url, err.Error())
	}

	// 使用日志打印
	logrus.Debug("==response==", string(response))
	// 使用官方test打印
	t.Log(string(response))

	// 转换结果
	var raw map[string]interface{}
	if err := json.Unmarshal(response, &raw); err != nil {
		t.Error("json parse fail")
	}

	// assert := assert.New(t)
	// 断言请求码等于200
	assert.Equal(t, 200, status, "Http code error")
	assert.Equal(t, float64(200), raw["code"], "response code error")
	assert.NotEmpty(t, raw["message"], "response code error")
}
