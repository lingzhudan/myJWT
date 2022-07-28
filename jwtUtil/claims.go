package jwtUtil

import (
	"errors"
	"time"
)

// MyClaims jwt中的载荷,不能存放敏感信息
type MyClaims struct {
	UserId    int64
	ExpiresAt int64
}

// Valid in MyClaims
func (mc MyClaims) Valid() error {
	now := time.Now().Add(-time.Hour * 24).Unix()
	//当token中的过期时间超过设定
	if mc.ExpiresAt <= now {
		return errors.New("token is invalid")
	}
	return nil
}
