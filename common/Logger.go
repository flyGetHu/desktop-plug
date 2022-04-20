package common

import (
	"log"
)

var (
	TraceLog *log.Logger // 记录所有日志
	InfoLog  *log.Logger // 重要的信息
	WarnLog  *log.Logger // 需要注意的信息
	ErrorLog *log.Logger // 非常严重的问题
)
