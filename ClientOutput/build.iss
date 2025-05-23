; 脚本由 Inno Setup 脚本向导 生成！
; 有关创建 Inno Setup 脚本文件的详细资料请查阅帮助文档！

#define MyAppName "RemoteUnlock"
#define MyAppVersion "1.1"
#define MyAppPublisher "JohnRey"

[Setup]
; 注: AppId的值为单独标识该应用程序。
; 不要为其他安装程序使用相同的AppId值。
; (生成新的GUID，点击 工具|在IDE中生成GUID。)
AppId={{9B2DDDA2-0130-469E-9A9C-F9F33F358625}
AppName={#MyAppName}
AppVersion={#MyAppVersion}
;AppVerName={#MyAppName} {#MyAppVersion}
AppPublisher={#MyAppPublisher}
DefaultDirName=D:\Program Files\{#MyAppName}
DefaultGroupName={#MyAppName}
OutputBaseFilename=RemoteUnlockClient
Compression=lzma
SolidCompression=yes

[Languages]
Name: "chinesesimp"; MessagesFile: "compiler:Default.isl"

[Tasks]
; 添加用户可选任务
Name: "AutoStartup"; Description: "开机自动启动监控程序"; GroupDescription: "启动设置:"

[Files]
Source: "G:\YSTRemoteUnlock\ClientOutput\RemoteUnlock.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "G:\YSTRemoteUnlock\ClientOutput\WatchFileAndUnlock.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "G:\YSTRemoteUnlock\ClientOutput\conf\*"; DestDir: "{app}\conf"; Flags: ignoreversion
; 注意: 不要在任何共享系统文件上使用“Flags: ignoreversion”

[Icons]
; 核心实现：启动文件夹快捷方式
Name: "{commonstartup}\文件监控服务"; Filename: "{app}\WatchFileAndUnlock.exe"; Tasks: AutoStartup

[Registry]
Root: HKCR; Subkey: "Directory\shell\RemoteUnlock"; ValueType: string; ValueName: ""; ValueData: "RemoteUnlock"; Flags: uninsdeletekey
Root: HKCR; Subkey: "Directory\shell\RemoteUnlock\command"; ValueType: string; ValueName: ""; ValueData: """{app}\RemoteUnlock.exe"" ""%1"""; Flags: uninsdeletekey
Root: HKCR; Subkey: "*\shell\RemoteUnlock"; ValueType: string; ValueName: ""; ValueData: "RemoteUnlock"; Flags: uninsdeletekey
Root: HKCR; Subkey: "*\shell\RemoteUnlock\command"; ValueType: string; ValueName: ""; ValueData: """{app}\RemoteUnlock.exe"" ""%1"""; Flags: uninsdeletekey