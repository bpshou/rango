package service

import (
	"encoding/base64"
	"fmt"
	"rango/tools/cipher"

	"github.com/gin-gonic/gin"
)

func Aes(c *gin.Context) {
	text := "123" // 你要加密的数据
	// AesKey := []byte("#HvL%$o0oNNoOZnk#o2qbqCeQB1iXeIR") // 对称秘钥长度必须是16的倍数
	AesKey := []byte("aaaaaaaaaaaaaaaa") // 对称秘钥长度必须是16的倍数

	fmt.Printf("明文: %s\n秘钥: %s\n", text, string(AesKey))
	encrypted, err := cipher.AesEncrypt([]byte(text), AesKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密后: %s\n", base64.StdEncoding.EncodeToString(encrypted))

	origin, err := cipher.AesDecrypt(encrypted, AesKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("解密后明文: %s\n", string(origin))
	c.JSON(200, gin.H{
		"code":    200,
		"message": "index",
	})
}
