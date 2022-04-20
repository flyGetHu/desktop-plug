package conf

import (
	"log"
	"os/exec"
	"time"
)

//初始化一些配置
func init() {
	initLog()
	initBanner()
	//打开默认的管理端
	// openDefSite()
}

func openDefSite() {
	go func() {
		err := exec.Command(`cmd`, `/c`, `start`, `https://www.baidu.com/`).Start()
		if err != nil {
			log.Println("使用浏览器打开指定网址错误:", err)
			time.Sleep(time.Second * 2)
			log.Println("请手动打开!")
			time.Sleep(time.Second * 2)
			log.Println("请手动打开!")
			time.Sleep(time.Second * 2)
			log.Println("请手动打开!")
		}
	}()
}
