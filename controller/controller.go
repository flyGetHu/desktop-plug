package controller

import (
	"log"
	"net/http"

	"bailidaming.com/anjun/destop/common"
)

func init() {
	registeredRouters()
}

//初始化路由
func registeredRouters() {
	defaultRegisteredRouter()
	registeredPrinterRouter()
	registeredScaleRouter()
	log.Println("路由版本V1,已挂载!")
}

//默认加载路由
func defaultRegisteredRouter() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		common.Result{Msg: "桌面插件已开启!"}.Success(writer)
	})
}
