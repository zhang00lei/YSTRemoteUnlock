; �ű��� Inno Setup �ű��� ���ɣ�
; �йش��� Inno Setup �ű��ļ�����ϸ��������İ����ĵ���

#define MyAppName "RemoteUnlock"
#define MyAppVersion "1.1"
#define MyAppPublisher "JohnRey"

[Setup]
; ע: AppId��ֵΪ������ʶ��Ӧ�ó���
; ��ҪΪ������װ����ʹ����ͬ��AppIdֵ��
; (�����µ�GUID����� ����|��IDE������GUID��)
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
; ����û���ѡ����
Name: "AutoStartup"; Description: "�����Զ�������س���"; GroupDescription: "��������:"

[Files]
Source: "G:\YSTRemoteUnlock\ClientOutput\RemoteUnlock.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "G:\YSTRemoteUnlock\ClientOutput\WatchFileAndUnlock.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "G:\YSTRemoteUnlock\ClientOutput\conf\*"; DestDir: "{app}\conf"; Flags: ignoreversion
; ע��: ��Ҫ���κι���ϵͳ�ļ���ʹ�á�Flags: ignoreversion��

[Icons]
; ����ʵ�֣������ļ��п�ݷ�ʽ
Name: "{commonstartup}\�ļ���ط���"; Filename: "{app}\WatchFileAndUnlock.exe"; Tasks: AutoStartup

[Registry]
Root: HKCR; Subkey: "Directory\shell\RemoteUnlock"; ValueType: string; ValueName: ""; ValueData: "RemoteUnlock"; Flags: uninsdeletekey
Root: HKCR; Subkey: "Directory\shell\RemoteUnlock\command"; ValueType: string; ValueName: ""; ValueData: """{app}\RemoteUnlock.exe"" ""%1"""; Flags: uninsdeletekey
Root: HKCR; Subkey: "*\shell\RemoteUnlock"; ValueType: string; ValueName: ""; ValueData: "RemoteUnlock"; Flags: uninsdeletekey
Root: HKCR; Subkey: "*\shell\RemoteUnlock\command"; ValueType: string; ValueName: ""; ValueData: """{app}\RemoteUnlock.exe"" ""%1"""; Flags: uninsdeletekey