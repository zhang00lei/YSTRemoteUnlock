package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg     *ini.File
	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	FileSavePath  string
	FileMaxSize   int64
	UnlockExePath string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini':%v", err)
	}
	loadBase()
	loadServer()
	loadFile()
}

func loadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server':%v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func loadFile() {
	sec, err := Cfg.GetSection("unlock_info")
	if err != nil {
		log.Fatalf("Fail to get section 'app':%v", err)
	}
	FileSavePath = sec.Key("PATH").MustString("./bin/files/share_file")
	FileMaxSize = sec.Key("MAX_SIZE").MustInt64(1024 * 5)
	UnlockExePath = sec.Key("UNLOCK_EXE_PATH").MustString("E:\\MyProject\\YST-Remote\\UnlockAll\\wps.exe")
}
