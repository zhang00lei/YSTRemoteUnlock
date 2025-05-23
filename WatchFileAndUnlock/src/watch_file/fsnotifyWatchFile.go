package watch_file

import (
	"WatchFileAndUnlock/src/unlock_file"
	"WatchFileAndUnlock/src/util"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"path/filepath"
	"time"
)

type Watch struct {
	Watch *fsnotify.Watcher
}

func (w *Watch) WatchDir(dir []string) {
	//通过Walk来遍历目录下的所有子目录
	for _, dirTemp := range dir {
		//w.Watch.Add(dirTemp)
		filepath.Walk(dirTemp, func(path string, info os.FileInfo, err error) error {
			//这里判断是否为目录，只需监控目录即可
			//目录下的文件也在监控范围内，不需要我们一个一个加
			if info != nil && info.IsDir() {
				path, err := filepath.Abs(path)
				if err != nil {
					return err
				}
				err = w.Watch.Add(path)
				if err != nil {
					return err
				}
			}
			return nil
		})
	}

	go func() {
		for {
			select {
			case ev := <-w.Watch.Events:
				{
					fmt.Println(ev.Op, ev.Name)
					if ev.Op&fsnotify.Create == fsnotify.Create {
						//这里获取新创建文件的信息，如果是目录，则加入监控中
						fi, err := os.Stat(ev.Name)
						if err == nil {
							if fi.IsDir() {
								//time.Sleep(20 * time.Millisecond)
								//addUnlockFile(ev.Name)
							} else {
								time.Sleep(20 * time.Millisecond)
								if util.FileIsLocked(ev.Name) {
									unlock_file.AddUnlockFile(ev.Name)
								}
							}
						}
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						fi, err := os.Stat(ev.Name)
						if err == nil {
							if !fi.IsDir() {
								//time.Sleep(20 * time.Millisecond)
								//if util.FileIsLocked(ev.Name) {
								//	addUnlockFile(ev.Name)
								//}
							} else {
								time.Sleep(20 * time.Millisecond)
								unlock_file.AddUnlockFile(ev.Name)
							}
						}
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						fi, err := os.Stat(ev.Name)
						if err == nil && fi.IsDir() {
							w.Watch.Remove(ev.Name)
						}
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						//fmt.Println("重命名文件 : ", ev.Name)
						//如果重命名文件是目录，则移除监控
						//注意这里无法使用os.Stat来判断是否是目录了
						//因为重命名后，go已经无法找到原文件来获取信息了
						//所以这里就简单粗爆的直接remove好了
						fmt.Println("移出监控 : ", ev.Name)
						w.Watch.Remove(ev.Name)
					}
				}
			case err := <-w.Watch.Errors:
				{
					fmt.Println("error : ", err)
					return
				}
			}
		}
	}()
}
