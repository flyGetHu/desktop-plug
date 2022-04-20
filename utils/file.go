package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"

	uuid "github.com/satori/go.uuid"
)

// PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetTempFilePath 获取一个临时文件目录
// fileSuffix 文件后缀
func GetTempFilePath(fileSuffix string) (string, error) {
	pathPre := "./temp/"
	exists, _ := PathExists(pathPre)
	if !exists {
		if err := os.Mkdir(pathPre, os.ModePerm); err != nil {
			return "", err
		}
	}
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	filePath := pathPre + id.String() + "." + fileSuffix
	return filePath, nil
}

// GetAllFiles 遍历出指定目录下的所有文件
func GetAllFiles(dirPth string) (files []string, err error) {
	fis, err := ioutil.ReadDir(filepath.Clean(filepath.ToSlash(dirPth)))
	if err != nil {
		return nil, err
	}

	for _, f := range fis {
		_path := filepath.Join(dirPth, f.Name())

		if f.IsDir() {
			fs, _ := GetAllFiles(_path)
			files = append(files, fs...)
			continue
		}

		// 指定格式
		switch filepath.Ext(f.Name()) {
		case ".png", ".jpg", ".jpeg":
			files = append(files, _path)
		}
	}

	return files, nil
}
