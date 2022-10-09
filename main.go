package main

import (
	"net/http"

	"bailidaming.com/anjun/destop/common"
	"bailidaming.com/anjun/destop/middleware"
	"bailidaming.com/anjun/destop/service"
	"fyne.io/systray"

	_ "bailidaming.com/anjun/destop/conf"
	_ "bailidaming.com/anjun/destop/controller"
)

func main() {
	//启动web服务
	go func() {
		server := http.Server{
			Addr:    ":3000",
			Handler: new(middleware.CustomizeMiddleware),
		}
		if err := server.ListenAndServe(); err != nil {
			common.ErrorLog.Fatalln("启动项目失败:", err)
		}
	}()
	// 托盘程序逻辑
	systray.Run(service.OnReady, service.OnExit)
}
