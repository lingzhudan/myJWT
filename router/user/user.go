package user

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io/ioutil"
	"myJWT/jwtUtil"
	"myJWT/middleware"
	"net/http"
)

func Router(srv *rest.Server) {
	srv.AddRoutes([]rest.Route{
		{
			Path:    "/login",
			Method:  http.MethodPost,
			Handler: Login,
		},
		{
			Path:    "/login",
			Method:  http.MethodGet,
			Handler: LoginPage,
		},
		{
			Path:    "/register",
			Method:  http.MethodPost,
			Handler: Register,
		},
		{
			Path:    "/info",
			Method:  http.MethodGet,
			Handler: middleware.Auth(Info),
		},
	}, rest.WithPrefix("/user"))
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data, err := ioutil.ReadFile("./view/login.html")
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	var resp LoginResponse
	if err := httpx.Parse(r, &req); err != nil { // 解析参数
		httpx.Error(w, err)
		return
	}
	token, err := jwtUtil.NewJWTByUserId(req.UserId)
	if err != nil {
		httpx.Error(w, err)
		return
	}
	tokenBytes, err := json.Marshal(token)
	resp.Token = string(tokenBytes)
	httpx.WriteJson(w, http.StatusOK, resp)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	var resp RegisterResponse
	if err := httpx.Parse(r, &req); err != nil { // 解析参数
		httpx.Error(w, err)
		return
	}
	//调用rpc服务存储注册信息
	//if resp, err := rpc.Register(); err != nil {
	//	httpx.Error(w, err)
	//	return
	//}
	httpx.WriteJson(w, http.StatusOK, resp)
}

func Info(w http.ResponseWriter, r *http.Request) {

	var resp InfoResponse
	//for test
	resp = InfoResponse{
		UserId: r.Context().Value("userId").(int64),
	}
	//调用rpc服务获取用户信息
	//if resp, err := rpc.Info(); err != nil {
	//	httpx.Error(w, err)
	//	return
	//}
	httpx.WriteJson(w, http.StatusOK, resp)
}
