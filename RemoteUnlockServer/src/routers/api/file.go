package api

import (
	"RemoteUnlockServer/src/setting"
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/duke-git/lancet/v2/random"
	"github.com/duke-git/lancet/v2/system"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"syscall"
)

func UploadUnlockFile(c *gin.Context) {
	err := c.Request.ParseMultipartForm(setting.FileMaxSize << 20)
	if err != nil {
		log.Println("上传文件异常", err.Error())
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("上传文件异常", err.Error())
		return
	}
	fileName, _ := random.UUIdV4()
	filePath := path.Join(setting.FileSavePath, fileName)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		log.Println("保存文件异常", err.Error())
		return
	}
	//todo test
	//filePath = fmt.Sprintf("E:\\MyProject\\YST-Remote\\RemoteUnlockServer\\bin\\files\\%v", fileName)
	absPath, _ := filepath.Abs(filePath)
	UnlockFile(absPath)

	//filePath := filepath.Join(setting.FileSavePath, fileName)
	filePath, _ = filepath.Abs(filePath)
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.File(filePath)
	if fileutil.IsExist(filePath) {
		fileutil.RemoveFile(filePath)
	}
}

func UnlockFile(pathTemp string) {
	docPath := pathTemp + ".docx"
	os.Rename(pathTemp, docPath)
	cmd := fmt.Sprintf(`& "%v"  "%v"`, setting.UnlockExePath, docPath)
	_, _, err := system.ExecCommand(cmd, func(cmd *exec.Cmd) {
		cmd.SysProcAttr = &syscall.SysProcAttr{
			HideWindow:    true,
			CreationFlags: 0x08000000,
		}
	})
	if err != nil {
		log.Println("Failed to run command:", cmd, err.Error())
		fmt.Scanln()
	} else {
		dstFilePath := docPath + ".temp"
		os.Rename(dstFilePath, pathTemp)
	}
}
