package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"

	"bailidaming.com/anjun/destop/common"
	"bailidaming.com/anjun/destop/utils"
)

func registeredPrinterRouter() {
	http.HandleFunc("/v1/printer", printHandler)
}

func printHandler(writer http.ResponseWriter, request *http.Request) {
	var err error
	printerType := request.FormValue("type")
	switch printerType {
	case "url":
		file := request.FormValue("file")
		if len(file) != 0 {
			err = printByUrl(file)
		} else {
			err = fmt.Errorf("传入文件链接不能为空")
		}
	case "file":
		if _, file, fileErr := request.FormFile("file"); fileErr == nil {
			err = printByFile(file)
		} else {
			common.ErrorLog.Println("读取上传文件失败:", fileErr)
			err = fmt.Errorf("读取上传文件失败")
		}
	default:
		common.Result{Msg: "无效打印方式!"}.Error(writer)
		return
	}
	if err != nil {
		common.ErrorLog.Println("打印失败:", err)
		common.Result{Msg: err.Error()}.Error(writer)
	} else {
		common.InfoLog.Println("打印成功!")
		common.Result{}.Success(writer)
	}
}

// 打印文件处理器 传入参数为一个网络文件链接
func printByUrl(file string) error {
	response, err := http.Get(file)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		common.ErrorLog.Printf("%s 获取面单异常:%v\n", file, err)
		return fmt.Errorf("获取面单异常")
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if len(bytes) == 0 {
		common.ErrorLog.Printf("%s 面单不存在:%v\n", file, err)
		return fmt.Errorf("面单不存在")
	}
	filePath, err := utils.GetTempFilePath("pdf")
	if err != nil {
		return err
	}
	if err = fileToPrinter(bytes, filePath); err != nil {
		return err
	}
	return nil
}

//printByFile 根据上传的文件输出打印机
func printByFile(formFile *multipart.FileHeader) error {
	filePath, err := utils.GetTempFilePath("jpg")
	if err != nil {
		return err
	}
	file, err := formFile.Open()
	if err != nil {
		return err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Println("关闭打印文件失败:", err)
		}
	}(file)
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	if err = fileToPrinter(bytes, filePath); err != nil {
		return err
	}
	return nil
}

// 将文件数据投送到打印机 bytes 文件数据  filePath 文件保存的本地地址
func fileToPrinter(bytes []byte, filePath string) error {
	if err := ioutil.WriteFile(filePath, bytes, 0644); err != nil {
		common.ErrorLog.Println("下载文件失败:", err)
		return err
	}
	if err := utils.FileToPrinter(filePath); err != nil {
		return err
	}
	return nil
}
