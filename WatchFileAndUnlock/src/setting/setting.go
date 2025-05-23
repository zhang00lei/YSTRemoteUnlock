package setting

import (
	"WatchFileAndUnlock/src"
	"fmt"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/go-ini/ini"
	"log"
	"os"
	"path/filepath"
)

var (
	Cfg             *ini.File
	MonitorDir      []string
	MonitorInterval float64
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
	pathTemp := sec.Key("MONITOR_DIR").MustString("")

	if pathTemp != "" {
		MonitorDir = strutil.SplitEx(pathTemp, "|", true)
	}
	MonitorInterval = sec.Key("MONITOR_INTERVAL").MustFloat64(1)
}
