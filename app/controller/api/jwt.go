package api

import (
	"fmt"
	"rango/app/controller"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type Jwt struct {
	controller.Base
}

func (the Jwt) Index(c *gin.Context) {

	signKey := viper.GetString("jwt.token")
	signKeyByte := []byte(signKey)
	// Create the Claims
	// claims := &jwt.RegisteredClaims{
	// 	ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
	// 	Issuer:    "test",
	// }

	claimsCustom := &jwt.MapClaims{
		"username": "test",
		"id":       1,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsCustom)
	authorization, err := token.SignedString(signKeyByte)

	if err != nil {
		c.JSON(200, gin.H{
			"code":  400,
			"error": err,
		})
		return
	}

	the.ParseToken(c, authorization)

	c.JSON(200, gin.H{
		"code":          200,
		"token":         token,
		"authorization": authorization,
	})
}

func (the Jwt) ParseToken(c *gin.Context, tokenString string) {
	// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0IiwiZXhwIjoxNTE2MjM5MDIyfQ.4vrcTWoa1Fc4DMl-yqcplGX5zQVhQxaV4jkRrbyHFxQ"
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return viper.GetString("jwt.token"), nil
	})

	if err != nil {
		c.JSON(200, gin.H{
			"code":  400,
			"error": err,
		})
	}

	claims := token.Claims.(*jwt.MapClaims)
	fmt.Println(claims)
}
