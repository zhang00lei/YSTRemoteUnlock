package util

import (
	"bufio"
	"bytes"
	"os"
)

var lockedByte = []byte{20, 35, 101}

func FileIsLocked(filePath string) bool {
	data, err := readBlock(filePath, 4)
	if err != nil {
		return false
	}
	return bytes.Equal(data[1:], lockedByte)
}

func readBlock(filePth string, bufSize int) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf := make([]byte, bufSize)
	bfRd := bufio.NewReader(f)
	_, err = bfRd.Read(buf)
	return buf, err
}
