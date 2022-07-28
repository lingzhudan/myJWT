package middleware

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/rest/httpx"
	"myJWT/jwtUtil"
	"net/http"
)

// Auth 判断是否有token并解析，如果合法通过，过期返回json格式新token
func Auth(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		claims, err := jwtUtil.ParseWithClaims(token)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		if newToken, ok := jwtUtil.Renewal(claims); ok {
			tokenBytes, _ := json.Marshal(newToken)
			httpx.WriteJson(w, http.StatusCreated, struct {
				Token string `json:"token"`
			}{
				string(tokenBytes),
			})
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "userId", claims.UserId)
		next(w, r.WithContext(ctx))
	}
}
