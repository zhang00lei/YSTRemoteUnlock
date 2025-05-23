package watch_file

import (
	"WatchFileAndUnlock/src/unlock_file"
	"fmt"
	getFolderSize "github.com/markthree/go-get-folder-size/src"
	"slices"
)

var fileSizeMap map[string]int64
var dirUnlockMap []string

func WatchDir(dirs []string) {
	if fileSizeMap == nil {
		fileSizeMap = make(map[string]int64)
	}
	for _, dir := range dirs {
		size, _ := dirSize(dir)
		if size != fileSizeMap[dir] {
			fmt.Println(dir)
			dirUnlockMap = append(dirUnlockMap, dir)
			//unlock_file.AddUnlockFile(dir)
		} else {
			if slices.Contains(dirUnlockMap, dir) {
				unlock_file.AddUnlockFile(dir)
				dirUnlockMap = slices.DeleteFunc(dirUnlockMap, func(s string) bool {
					return s == dir
				})
			}
		}
		fileSizeMap[dir] = size
	}
}

func dirSize(path string) (int64, error) {
	return getFolderSize.Invoke(path)
}
