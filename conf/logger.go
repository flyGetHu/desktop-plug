package conf

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"bailidaming.com/anjun/destop/common"
	"bailidaming.com/anjun/destop/utils"
)

// InitLog 初始化程序一些基础配置
func initLog() {
	//初始化日志
	var err error
	b, _ := utils.PathExists("./logs/")
	if !b {
		err = os.Mkdir("./logs/", 0666)
		if err != nil {
			log.Fatalln("创建日志目录失败:", err)
		}
	}
	file, err := os.OpenFile("./logs/errors.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	common.TraceLog = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	common.InfoLog = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	common.WarnLog = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	common.ErrorLog = log.New(io.MultiWriter(file, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
