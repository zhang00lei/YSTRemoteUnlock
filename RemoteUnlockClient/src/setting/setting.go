package setting

import (
	"RemoteUnlockClient/src"
	"fmt"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/go-ini/ini"
	"log"
	"os"
	"path/filepath"
)

var (
	Cfg        *ini.File
	Host       string
	MonitorDir []string
	IgnoreName []string
)

func init() {
	var err error
	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)
	confPath := fmt.Sprintf("%s/conf/app.ini", exeDir)
	if src.IS_EDITOR {
		confPath = "conf/app.ini"
	}
	Cfg, err = ini.Load(confPath)
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini':%v", err)
	}

	loadServer()
}

func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server':%v", err)
	}
	Host = sec.Key("HOST").MustString("127.0.0.1:8090")
	pathTemp := sec.Key("MONITOR_DIR").MustString("")
	if pathTemp != "" {
		MonitorDir = strutil.SplitEx(pathTemp, "|", true)
	}

	pathTemp = sec.Key("IGNORE_NAME").MustString("")

	if pathTemp != "" {
		IgnoreName = strutil.SplitEx(pathTemp, "|", true)
	}
}
