package utils

import (
	"log"
	"time"

	"bailidaming.com/anjun/destop/common"
	"go.bug.st/serial"
)

// OpenSerial 打开指定串口并返回串口 baudRate 波特率 300-460800 默认9600
// dataBits 校验位 [5,6,7,8]
func OpenSerial(comPort string, baudRate int, dataBits int) error {
	mode := &serial.Mode{
		BaudRate: baudRate,
		DataBits: dataBits,
	}
	port, err := serial.Open(comPort, mode)
	if err != nil {
		log.Println("打开端口失败", err)
		return err
	}
	err = port.SetReadTimeout(time.Second * 2)
	if err != nil {
		common.ErrorLog.Println("打开本机串口失败:", err)
		return err
	}
	common.SerialPort = port
	return nil
}
