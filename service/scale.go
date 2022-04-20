package service

import (
	"bailidaming.com/anjun/destop/common"
	"bailidaming.com/anjun/destop/utils"
	"go.bug.st/serial"
	"net/http"
	"strconv"
	"strings"
)

func CheckScaleParams(writer http.ResponseWriter, request *http.Request) (string, bool) {
	comPort := request.FormValue("comPort")
	if len(comPort) == 0 {
		common.Result{Msg: "串口号必传:comPort!"}.Error(writer)
		return "", true
	}
	portList, err := serial.GetPortsList()
	if err != nil {
		common.ErrorLog.Println("获取本机存在的串口失败:", err)
		common.Result{Msg: "打开电子秤串口失败!"}.Error(writer)
		return "", true
	}
	if !utils.ContainsInSlice(portList, comPort) {
		common.Result{Msg: "传递串口参数与本地可用串口不匹配,请重新选择!"}.Error(writer)
		return "", true
	}
	baudRateForm := request.FormValue("baudRate")
	if baudRateForm == "" {
		baudRateForm = "9600"
	}
	baudRate, err := strconv.Atoi(baudRateForm)
	if err != nil {
		common.Result{Msg: "串口比特率传递错误,范围300-460800!"}.Error(writer)
		return "", true
	}
	dataBitsForm := request.FormValue("dataBits")
	if dataBitsForm == "" {
		dataBitsForm = "8"
	}
	dataBits, err := strconv.Atoi(dataBitsForm)
	if err != nil {
		common.Result{Msg: "串口数据位传递错误,范围[5,6,7,8]中一个!"}.Error(writer)
		return "", true
	}
	//电子称类型
	var scaleType = request.FormValue("scaleType")
	if len(scaleType) == 0 {
		common.Result{Msg: "电子秤类型必传,切必须为:" + strings.Join(common.SerialTypes, ",") + "中的一个!"}.Error(writer)
		return "", true
	}
	if !utils.ContainsInSlice(common.SerialTypes, scaleType) {
		common.Result{Msg: "电子秤类型为:" + strings.Join(common.SerialTypes, ",") + "中的一个!"}.Error(writer)
		return "", true
	}
	//打开串口前,先将重量归零
	common.ScaleWeight = -1
	if err = utils.OpenSerial(comPort, baudRate, dataBits); err != nil {
		common.Result{Msg: err.Error()}.Error(writer)
		return "", true
	}
	return scaleType, false
}

// DecodingScale 切分字符串获取校验位中间读数 65.000=675.000=675.000=675
//此处为一种电子秤读数解码方式,后续还会添加多种电子秤解码方式
//校验读数精度是否正常,防止粘包 丢包
//适配电子称解码
func DecodingScale(buff []byte, scaleType string) {
	if len(scaleType) == 0 {
		scaleType = "TSC-L"
	}
	data := string(buff)
	switch scaleType {
	case "TSC-L":
		//默认电子秤类型
		split := strings.Split(data, "=")
		for _, num := range split {
			if len(num) == 7 {
				num = utils.ReverseStr(num)
				if utils.StrIsNum(num) {
					if weight, err := strconv.ParseFloat(num, 64); err == nil {
						common.ScaleWeight = weight
						break
					}
				}
			}
		}
	case "PRIX-3FIT":
		//巴西仓库电子秤:00154002260015400226
		split := strings.Split(data, "")
		var tempStr = ""
		for _, num := range split {
			if len(tempStr) == 5 && utils.StrIsNum(tempStr) {
				break
			}
			if len(tempStr) != 0 && !utils.StrIsNum(num) {
				return
			}
			if utils.StrIsNum(num) {
				tempStr += num
			}
		}
		var num = ""
		if utils.StrIsNum(tempStr) && len(tempStr) == 5 {
			for i, item := range tempStr {
				if i == 3 {
					num += "."
				}
				num += string(item)
			}
		}
		if weight, err := strconv.ParseFloat(num, 64); err == nil {
			common.ScaleWeight = weight
		}
	default:
		common.ScaleWeight = -1
		common.ErrorLog.Println("错误的电子称类型:", scaleType)
	}
}
