package main

import (
	"WatchFileAndUnlock/src"
	"WatchFileAndUnlock/src/lockfile"
	"WatchFileAndUnlock/src/setting"
	"WatchFileAndUnlock/src/unlock_file"
	"WatchFileAndUnlock/src/watch_file"
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
)

func main() {
	if !src.IS_EDITOR {
		lockFile, err := lockfile.LockFile()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer lockfile.UnlockFile(lockFile)
	}

	// fsnotify模式监控目录变化
	//watch, _ := fsnotify.NewWatcher()
	//w := watch_file.Watch{
	//	Watch: watch,
	//}
	//w.WatchDir(setting.MonitorDir)

	spec := fmt.Sprintf("@every %fs", setting.MonitorInterval)
	c := cron.New()
	c.AddFunc(spec, func() {
		watch_file.WatchDir(setting.MonitorDir)
	})
	c.Start()
	defer c.Stop()

	go unlock_file.UnlockFiles()
	select {}
}
