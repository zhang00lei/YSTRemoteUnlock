# 设置server程序的源文件路径
$serverSourceFile = "G:\YSTRemoteUnlock\RemoteUnlockClient\main.go"
cd (Split-Path -Parent $serverSourceFile)
go build -o "G:\YSTRemoteUnlock\ClientOutput\RemoteUnlock.exe"

$watchFileSourceFile = "G:\YSTRemoteUnlock\WatchFileAndUnlock\main.go"
cd (Split-Path -Parent $watchFileSourceFile)
#go build -o E:\MyProject\YST-Remote\ClientOutput\WatchFileAndUnlock.exe
go build -ldflags "-H windowsgui" -o "G:\YSTRemoteUnlock\ClientOutput\WatchFileAndUnlock.exe"

# 设置Inno Setup脚本文件路径
$issFile = "G:\YSTRemoteUnlock\ClientOutput\build.iss"
# 编译Inno Setup脚本
& 'D:\Program Files (x86)\Inno Setup 5\ISCC.exe' $issFile
