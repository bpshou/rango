package api

import (
	"encoding/base64"
	"rango/app/controller"
	"rango/app/service"

	"github.com/gin-gonic/gin"
)

type Secret struct {
	controller.Base
}

type SecretParams struct {
	Input  string `form:"input" json:"input" xml:"input"`
	Secret string `form:"secret" json:"secret" xml:"secret"`
}

func (the Secret) Create(c *gin.Context) {
	// 解析参数
	var params SecretParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(200, gin.H{
			"code":    4001,
			"message": err.Error(),
		})
		return
	}

	// 数据
	input := params.Input
	secret := params.Secret

	if input == "" {
		c.JSON(200, gin.H{
			"code":    4001,
			"message": "参数为空",
		})
		return
	}
	if secret == "" {
		secret = "123456"
	}

	outputByte, err := service.Encrypt(input, secret)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    4004,
			"message": err.Error(),
			"output":  string(outputByte),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "secret",
		"output":  string(outputByte),
		"input":   input,
	})
}

func (the Secret) Decrypt(c *gin.Context) {
	// 解析参数
	var params SecretParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(200, gin.H{
			"code":    4001,
			"message": err.Error(),
		})
		return
	}

	// 数据
	input := params.Input
	secret := params.Secret

	if input == "" {
		c.JSON(200, gin.H{
			"code":    4001,
			"message": "参数为空",
		})
		return
	}
	if secret != "" {
		secretByte, _ := base64.StdEncoding.DecodeString(secret)
		secret = string(secretByte)
	}
	if secret == "" {
		secret = "123456"
	}

	outputByte, err := service.Decrypt(input, secret)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    4004,
			"message": err.Error(),
			"output":  string(outputByte),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "secret",
		"output":  string(outputByte),
		"input":   input,
	})
}
