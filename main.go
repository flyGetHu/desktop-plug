package main

import (
	"net/http"

	"bailidaming.com/anjun/destop/common"
	"bailidaming.com/anjun/destop/middleware"

	_ "bailidaming.com/anjun/destop/conf"
	_ "bailidaming.com/anjun/destop/controller"
)

func main() {
	server := http.Server{
		Addr:    ":3000",
		Handler: new(middleware.CustomizeMiddleware),
	}
	//启动web服务
	if err := server.ListenAndServe(); err != nil {
		common.ErrorLog.Fatalln("启动项目失败:", err)
	}
}
