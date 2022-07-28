package jwtUtil

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

// jwtClaims jwt密钥,非公开
var jwtSecret = []byte("lingzhudan")

// NewJWTByUserId 根据id和secret创建一个token
func NewJWTByUserId(userId int64) (string, error) {
	claim := MyClaims{
		UserId:    userId,
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}

// ParseWithClaims 封装jwt-go的操作，返回合法的claims
func ParseWithClaims(token string) (*MyClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return new(MyClaims), err
	}

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		if claims, ok := tokenClaims.Claims.(*MyClaims); ok {
			return claims, nil
		}
	}

	return new(MyClaims), errors.New("token or claims is invalid")
}

// Renewal 过期返回新token，否则ok为false
func Renewal(claims *MyClaims) (string, bool) {
	if claims.ExpiresAt <= time.Now().Unix() {
		if newToken, err := NewJWTByUserId(claims.UserId); err == nil {
			return newToken, true
		}
	}
	return "", false
}
