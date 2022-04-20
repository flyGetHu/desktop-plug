package conf

import (
	"log"
	"os"

	"bailidaming.com/anjun/destop/common"
)

func initBanner() {
	//读取 banner
	banner, err := os.OpenFile("./bin/banner.txt", os.O_RDONLY, 0644)
	if err != nil {
		return
	}
	defer func(banner *os.File) {
		err := banner.Close()
		if err != nil {
			common.ErrorLog.Println("关闭banner文件失败:", err)
		}
	}(banner)
	data := make([]byte, 1024)
	_, err = banner.Read(data)
	if err == nil {
		log.Println(string(data))
		log.Println("v1.0.0 Anjun 桌面插件启动")
	}
}
