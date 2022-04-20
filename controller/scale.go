package controller

import (
	"bailidaming.com/anjun/destop/service"
	"net/http"
	"sync"
	"time"

	"bailidaming.com/anjun/destop/common"
	"go.bug.st/serial"
)

//关闭串口 同步锁 防止在读取数据的同时,关闭串口导致 panic
var closePortMutex sync.Mutex

//跳出循环获取电子秤读数
var portErr = true

func registeredScaleRouter() {
	//获取本机所有串口
	http.HandleFunc("/v1/scale/ports", getPortsList)
	//打开指定串口
	http.HandleFunc("/v1/scale/open", scaleOpen)
	//关闭打开的串口
	http.HandleFunc("/v1/scale/close", scaleClose)
	//获取当前电子秤稳定读数
	http.HandleFunc("/v1/scale/weight", scaleWeight)
}

//GetPortsList 获取本机存在的串口
func getPortsList(writer http.ResponseWriter, _ *http.Request) {
	if ports, err := serial.GetPortsList(); err != nil {
		common.Result{Msg: "获取串口失败"}.Error(writer)
	} else {
		common.InfoLog.Println("获取串口列表成功:", ports)
		common.Result{Data: ports}.Success(writer)
	}
}

// ScaleClose 关闭当前打开的串口
func scaleClose(writer http.ResponseWriter, _ *http.Request) {
	closePortMutex.Lock()
	defer closePortMutex.Unlock()
	if common.SerialPort != nil {
		err := common.SerialPort.Close()
		if err != nil {
			common.ErrorLog.Println("关闭串口失败:", err)
			common.Result{Msg: "电子秤关闭失败!"}.Error(writer)
			return
		}
		common.SerialPort = nil
	} else {
		common.Result{Msg: "没有串口被打开!"}.Error(writer)
		return
	}
	//读取电子秤数据 循环关闭
	portErr = false
	after := time.After(time.Millisecond * 200)
	<-after
	common.ScaleWeight = -1
	common.InfoLog.Println("串口关闭成功")
	common.Result{}.Success(writer)
}

// ScaleWeight 获取当前电子秤读数
func scaleWeight(writer http.ResponseWriter, _ *http.Request) {
	if common.SerialPort == nil {
		common.Result{Msg: "未打开相应电子秤串口,无法读取数据!"}.Error(writer)
		return
	}
	var weightBegin = -1.0
	for {
		weightBegin = common.ScaleWeight
		if weightBegin == -1 {
			common.Result{Msg: "读取电子秤错误,请检查电子秤是否正确链接!"}.Error(writer)
			return
		}
		//延迟指定毫秒数再获取一次读数,两次读数一致再返回
		//可以做到电子秤读数稳定
		after := time.After(time.Millisecond * 100)
		<-after
		weightEnd := common.ScaleWeight
		if weightEnd == weightBegin {
			break
		}
	}
	common.InfoLog.Println("获取重量成功:", weightBegin)
	common.Result{Data: weightBegin}.Success(writer)
}

//电子秤读数读取临时存放
//局部变量,重复利用缓存
var buff = make([]byte, 150)

// ScaleOpen 打开指定串口电子秤并读取数据
func scaleOpen(writer http.ResponseWriter, request *http.Request) {
	closePortMutex.Lock()
	defer closePortMutex.Unlock()
	if common.SerialPort != nil {
		if err := common.SerialPort.Close(); err != nil {
			common.ErrorLog.Println("关闭旧串口失败:", err)
			common.Result{Msg: "旧串口关闭失败"}.Error(writer)
			return
		}
		//读取电子秤数据 循环关闭
		portErr = false
		after := time.After(time.Millisecond * 100)
		<-after
		common.InfoLog.Println("关闭旧串口成功")
		common.SerialPort = nil
	}
	scaleType, done := service.CheckScaleParams(writer, request)
	if done {
		return
	}
	common.InfoLog.Println("电子秤串口打开成功!")
	common.Result{}.Success(writer)
	//异步存放电子秤读数 不阻塞当前线程
	go func() {
		portErr = true
		//判断读取串口数据是否异常,中断无限循环
		for portErr {
			func() {
				closePortMutex.Lock()
				defer closePortMutex.Unlock()
				if common.SerialPort == nil {
					common.ScaleWeight = -1
					common.ErrorLog.Println("串口为空,无法读取数据!")
					portErr = false
					return
				}
				if _, err := common.SerialPort.Read(buff); err != nil {
					common.ErrorLog.Println("读取串口失败:", err)
					common.ScaleWeight = -1
					portErr = false
					return
				}
				//解码串口数据
				service.DecodingScale(buff, scaleType)
				after := time.After(time.Millisecond * 50)
				<-after
			}()
		}
	}()
}
