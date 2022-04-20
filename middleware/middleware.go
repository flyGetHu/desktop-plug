package middleware

import (
	"log"
	"net/http"

	"bailidaming.com/anjun/destop/common"
)

// CustomizeMiddleware 自定义中间件
type CustomizeMiddleware struct {
	Next http.Handler
}

func (middleware CustomizeMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			common.ErrorLog.Println("捕获到异常:", err)
			common.Result{Msg: "系统异常,请稍后再试"}.Error(w)
		}
	}()
	if middleware.Next == nil {
		middleware.Next = http.DefaultServeMux
	}
	//值可以设为星号,也可以指定具体主机地址,可设置多个地址用逗号隔开,设为指定主机地址第三项才有效
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//允许请求头修改的类容
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	//允许使用cookie
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	middleware.Next.ServeHTTP(w, r)
	log.Printf("  %v  %v  %v  %v \n", r.Method, r.RemoteAddr, r.RequestURI, r.Proto)
}
