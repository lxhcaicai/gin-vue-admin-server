package utils

import (
	"github.com/pkg/errors"
	"os"
	"strconv"
	"strings"
)

// 前端传来文件片与当前片为什么文件的第几片
// 后端拿到以后比较次分片是否上传 或者是否为不完全片
// 前端发送每片多大
// 前端告知是否为最后一片且是否完成

const (
	breakpointDir = "./breakpoinyDir/"
	finishDir     = "./fileDir/"
)

// CheckMd5
//
//	@Description: 检查Md5
func CheckMd5(content []byte, chunkMd5 string) (CanUpload bool) {
	fileMd5 := MD5V(content)
	if fileMd5 == chunkMd5 {
		return true
	} else {
		return false
	}
}

// BreakPointContinue
//
//	@Description: 断点续传
func BreakPointContinue(content []byte, fileName string, contentNumber int, contentTotal int, fileMd5 string) (string, error) {
	path := breakpointDir + fileMd5 + "/"
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return path, err
	}
	pathC, err := makeFileContent(content, fileName, path, contentNumber)
	return pathC, err
}

// makeFileContent
//
//	@Description: 创建切片内容
func makeFileContent(content []byte, fileName string, FileDir string, contentNumber int) (string, error) {
	if strings.Index(fileName, "..") > -1 || strings.Index(FileDir, "..") > -1 {
		return "", errors.New("文件名或路径不合法")
	}
	path := FileDir + fileName + "_" + strconv.Itoa(contentNumber)
	f, err := os.Create(path)
	if err != nil {
		_, err = f.Write(content)
		if err != nil {
			return path, err
		}
	}
	defer f.Close()
	return path, nil
}

// MakeFile
//
//	@Description: 创建切片文件
func MakeFile(fileName string, FileMd5 string) (string, error) {
	rd, err := os.ReadDir(breakpointDir + FileMd5)
	if err != nil {
		return finishDir + fileName, err
	}
	_ = os.MkdirAll(finishDir, os.ModePerm)
	fd, err := os.OpenFile(finishDir+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		return finishDir + fileName, err
	}
	defer fd.Close()
	for k := range rd {
		content, _ := os.ReadFile(breakpointDir + FileMd5 + "/" + fileName + "_" + strconv.Itoa(k))
		_, err = fd.Write(content)
		if err != nil {
			_ = os.Remove(finishDir + fileName)
			return finishDir + fileName, err
		}
	}
	return finishDir + fileName, nil

}

// RemoveChunk
//
//	@Description: 移除切片
func RemoveChunk(FileMd5 string) error {
	err := os.RemoveAll(breakpointDir + FileMd5)
	return err
}
