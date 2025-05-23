package main

import (
	"RemoteUnlockClient/src/setting"
	"RemoteUnlockClient/src/util"
	"bytes"
	_ "embed"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/bytedance/gopkg/util/gopool"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/schollz/progressbar/v3"
	"io"
	"io/fs"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

//go:embed assets/logo.ico
var logoData []byte
var wg sync.WaitGroup

func onReady() {
	systray.SetIcon(logoData)
}

func onExit() {}

func main() {
	//if src.IS_EDITOR {
	//	watchAndUnlock()
	//	fmt.Println("over")
	//	return
	//}

	//pathTemp := "C:\\Users\\pc\\Desktop\\新建文件夹 (2)\\test1.doc"

	pathTemp := os.Args[1]
	info, err := os.Stat(pathTemp)
	if err != nil {
		fmt.Println("无法获取文件或目录信息：", err)
		fmt.Scanln()
		return
	}

	url := fmt.Sprintf(`http://%v/UploadUnlockFile`, setting.Host)
	if info.IsDir() {
		//目录
		lock := sync.Mutex{}
		s := spinner.New(spinner.CharSets[59], 500*time.Millisecond)
		s.Prefix = "搜索加密文件中 "
		s.Start()
		allFile, _ := getAllFileIncludeSubFolder(pathTemp)
		needFile := getNeedUnlockFile(allFile)
		s.Stop()
		bar := progressbar.Default(int64(len(needFile)))
		unlockCount := 0
		poolTemp := gopool.NewPool("Unlock", 200, gopool.NewConfig())
		if len(needFile) > 0 {
			go systray.Run(onReady, onExit)
		}
		for _, filePath := range needFile {
			wg.Add(1)
			temp := filePath
			poolTemp.Go(func() {
				uploadFile(temp, url)
				lock.Lock()
				unlockCount++
				bar.Add(1)
				lock.Unlock()
				wg.Done()
			})
		}
		wg.Wait()
		fmt.Scanln()
	} else if info.Mode().IsRegular() {
		//解密当前文件
		if !util.FileIsLocked(pathTemp) {
			return
		}
		go systray.Run(onReady, onExit)
		uploadFile(pathTemp, url)
	}
	fmt.Println("准备退出")
}

func uploadFile(filename string, targetURL string) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		fmt.Printf("创建form文件字段错误: %v", err)
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("打开文件错误: %v", err)
		return
	}

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		fmt.Printf("复制文件内容错误: %v", err)
		return
	}
	bodyWriter.Close()

	file.Close()

	req, err := http.NewRequest("POST", targetURL, bodyBuf)
	if err != nil {
		fmt.Printf("创建HTTP请求错误: %v", err)
		return
	}
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("发送请求错误: %v", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("准备下载")
	downloadAndSaveFile(resp, filename)
}

func downloadAndSaveFile(resp *http.Response, filename string) error {
	if fileutil.IsExist(filename) {
		fileutil.RemoveFile(filename)
	}
	// 创建文件并写入内容
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("创建文件错误: %v", err)
	}
	defer file.Close()

	// 将响应内容写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("保存文件错误: %v", err)
	}
	beeep.Notify("解密成功", fmt.Sprintf("解密成功：%v", filename), "")
	return nil
}

//
// getAllFileIncludeSubFolder
//  @Description: 获取目录下所有文件（包含子目录）
//  @param folder
//  @return []string
//  @return error
//
func getAllFileIncludeSubFolder(folder string) ([]string, error) {
	var result []string
	filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Println(err.Error())
			return err
		}
		for _, s := range setting.IgnoreName {
			if strings.Contains(path, s) {
				return err
			}
		}

		if info != nil && !info.IsDir() {
			result = append(result, path)
		}
		return nil
	})
	return result, nil
}

func getNeedUnlockFile(allFiles []string) []string {
	var result []string
	var lock sync.Mutex
	poolTemp := gopool.NewPool("Unlock", 8888, gopool.NewConfig())
	for _, pathTemp := range allFiles {
		filePath := pathTemp
		wg.Add(1)
		poolTemp.Go(func() {
			defer wg.Done()
			isLocked := util.FileIsLocked(filePath)
			if isLocked {
				lock.Lock()
				result = append(result, filePath)
				lock.Unlock()
			}
		})
	}
	wg.Wait()
	return result
}
