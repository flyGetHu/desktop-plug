package utils

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// FileToPrinter 向打印机输出文件
func FileToPrinter(filePath string) error {
	cmd := exec.Command("./bin/printer/SumatraPDF.exe", "-print-to-default", filePath)
	stdout, err := cmd.StderrPipe()
	defer func(stdout io.ReadCloser) {
		err := stdout.Close()
		if err != nil {
			log.Println("关闭打印命令失败:", err)
			return
		}
	}(stdout)
	if err = cmd.Start(); err != nil {
		return err
	}
	if _, err = ioutil.ReadAll(stdout); err != nil {
		return err
	}
	if err = os.Remove(filePath); err != nil {
		log.Println("删除文件失败:", filePath)
		return err
	}
	return nil
}
