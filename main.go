package main

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"log"
	"myJWT/router/user"
)

var opts []func(*rest.Server)

func Include(router ...func(server *rest.Server)) {
	opts = append(opts, router...)
}

func main() {
	srv, err := rest.NewServer(rest.RestConf{
		Port: 9090, // 侦听端口
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{Path: "./logs"}, // 日志路径
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	Include(user.Router)
	for _, opt := range opts {
		opt(srv)
	}

	srv.Start()
	defer srv.Stop()
}
