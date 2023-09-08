package tools

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type JWT struct {
	SigningKey string
}

func NewJWT() *JWT {
	return &JWT{
		SigningKey: viper.GetString("jwt.SigningKey"),
	}
}

type (
	BaseClaims struct {
		Uid      int64  `json:"uid"`
		Phone    string `json:"phone"`
		NickName string `json:"nick"`
	}
	// 通用
	CustomClaims struct {
		BaseClaims
		ResetTime int64 `json:"rt"`
		jwt.RegisteredClaims
	}
)

// 创建claims
func (the *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	claims := CustomClaims{
		BaseClaims: baseClaims,
		ResetTime:  time.Now().Add(20 * time.Hour * time.Duration(1)).Unix(), // 20个小时后
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    viper.GetString("jwt.Issuer"),                                         // 签名的发行者
			Audience:  jwt.ClaimStrings{"GVA"},                                               // 受众
			NotBefore: jwt.NewNumericDate(time.Now()),                                        // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * time.Duration(1))), // 过期时间 1 天  配置文件
		},
	}
	return claims
}

// 签发token
func (the *JWT) CreateToken(baseClaims BaseClaims) (string, error) {
	claims := the.CreateClaims(baseClaims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(the.SigningKey))
}

// 解析token
func (the *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(the.SigningKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, errors.New("couldn't handle this token")
}
