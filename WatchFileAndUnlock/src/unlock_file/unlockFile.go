package unlock_file

import (
	"WatchFileAndUnlock/src"
	"fmt"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/duke-git/lancet/v2/system"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"syscall"
)

var unlockFileList []string
var mutex sync.Mutex

func AddUnlockFile(filePath string) {
	mutex.Lock()
	if !slice.Contain(unlockFileList, filePath) {
		unlockFileList = append(unlockFileList, filePath)
	}
	mutex.Unlock()
}

func UnlockFiles() {
	unlockFilePath := ""
	for {
		if len(unlockFileList) > 0 {
			mutex.Lock()
			unlockFilePath = unlockFileList[0]
			unlockFileList = slice.DeleteAt(unlockFileList, 0)
			mutex.Unlock()
		}
		if unlockFilePath != "" {
			//fmt.Println("解密", unlockFilePath)
			go unlockFile(unlockFilePath)
			unlockFilePath = ""
		}
	}
}

func unlockFile(filePath string) {
	slfExePath, _ := os.Executable()
	exePath := fmt.Sprintf("%s/%s", filepath.Dir(slfExePath), "RemoteUnlock.exe")
	if src.IS_EDITOR {
		exePath = "D:\\Program Files (x86)\\RemoteUnlock\\RemoteUnlock.exe"
	}
	cmd := fmt.Sprintf(`& "%v"  "%v"`, exePath, filePath)
	fmt.Println(cmd)
	system.ExecCommand(cmd, func(cmd *exec.Cmd) {
		cmd.SysProcAttr = &syscall.SysProcAttr{
			HideWindow:    true,
			CreationFlags: 0x08000000,
		}
	})
}
