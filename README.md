# 亿赛通远程解密
在Win11系统中，安装亿赛通应用程序后，资源管理器会频繁的崩溃重启，以及其他各种莫名奇妙的问题，为解决这些问题，需将本机的亿赛通应用程序卸载（卸载方法可参考互联网教程）  
**确保在当前的公司环境下，使用[此程序](https://github.com/zhang00lei/YiSaiTongUnlock)可以正常使用**  
**注：此工具用于在能够正常打开加密文件的情境下使用，简而言之，就是省去了解密申请的步骤。各公司亿赛通加密策略不同，本程序可能不适用于所有策略，可根据实际情况进行修改**
## 服务端部署
服务端需windows操作系统，并且安装了亿赛通应用程序。解压RemoteUnlockServer.rar，将remote-unlock目录放置在服务端，并配置conf/app.ini中相关信息(如端口，解密文件路径)，如下：
![alt text](image.png)  
然后双击启动即可。
## 客户端  
在卸载完成的机器中，双击安装RemoteUnlockClient.exe，完成后在安装位置修改conf/app.ini文件，修改对应信息  
```powershell
HOST：服务端ip、端口  
MONITOR_DIR：需要监控的目录  
MONITOR_INTERVAL：监控时间间隔  
IGNORE_NAME：忽略的解密目录
```
![alt text](image-1.png)
## 使用
在客户端，选中文件或目录，右键RemoteUnlock即可完成解密。
## 已知问题  
1. 大文件解密可能存在服务端内存溢出问题
2. 监控程序CPU占用过高
# 免责声明  
本软件的任何使用仅用于非营利性的教育和测试目的。
