package lockfile

import (
	"os"
	"path/filepath"
)

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func LockFile() (*os.File, error) {
	exePath, err := os.Executable()
	if err != nil {
		return nil, err
	}

	lockFilePath := filepath.Join(filepath.Dir(exePath), "lockfile.lock")
	if fileExists(lockFilePath) {
		err = os.Remove(lockFilePath)
		if err != nil {
			return nil, err
		}
	}
	lockFile, err := os.Create(lockFilePath)
	if err != nil {
		return nil, err
	}

	return lockFile, nil
}

func UnlockFile(lockFile *os.File) {
	lockFile.Close()
	os.Remove(lockFile.Name())
}
