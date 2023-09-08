package curl

import (
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

/**
 * 发送GET请求
 *
 */
func Get(requestUrl string, requestBody []byte, requestHeader map[string]string) (httpCode int, body []byte, err error) {
	requestHeader["Method"] = "GET"
	return Request(requestUrl, requestBody, requestHeader)
}

/**
 * 发送POST请求
 *
 */
func Post(requestUrl string, requestBody []byte, requestHeader map[string]string) (httpCode int, body []byte, err error) {
	requestHeader["Method"] = "POST"
	requestHeader["Content-Type"] = "application/json"
	return Request(requestUrl, requestBody, requestHeader)
}

/**
 * 发送请求
 *
 */
func Request(requestUrl string, requestBody []byte, requestHeader map[string]string) (httpCode int, body []byte, err error) {
	// 请求
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	for key, value := range requestHeader {
		req.Header.Set(key, value)
	}

	req.SetRequestURI(requestUrl)
	// 设定body体
	if len(requestBody) > 0 {
		req.SetBody(requestBody)
	}

	if err = fasthttp.Do(req, resp); err != nil {
		logrus.Error("Curl Request err : ", err.Error())
		return
	}

	httpCode = resp.StatusCode()
	if httpCode != fasthttp.StatusOK {
		logrus.Error("Curl Request Status Code : ", httpCode)
	}

	body = resp.Body()
	return
}
