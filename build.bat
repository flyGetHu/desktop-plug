@echo off
::后续命令使用的是：UTF-8编码
chcp 65001
echo 中文
go build -v -o ./桌面插件.exe ./main.go
7z a -tZip 桌面插件.zip ./bin   ./桌面插件.exe
exit