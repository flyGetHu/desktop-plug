package common

import "go.bug.st/serial"

// ScaleWeight 存放当前电子秤读数
var ScaleWeight = -1.0

// SerialPort 串口
var SerialPort serial.Port = nil

// SerialTypes 电子秤类型
var SerialTypes = []string{"PRIX-3FIT", "TSC-L"}
