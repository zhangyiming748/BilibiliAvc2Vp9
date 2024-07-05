package util

import (
	"BilibiliAvc2Vp9/constant"
	"github.com/h2non/filetype"
	"log"
	"os"
	"path/filepath"
)

/*
获取当前文件夹和子文件夹下的全部文件的绝对路径 并放在字符串切片下
*/
func GetAllfiles(p *constant.Param) []string {
	log.Println("开始获取文件路径...")
	var allFiles []string
	filepath.Walk(p.GetRoot(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			allFiles = append(allFiles, path)
		}
		return nil
	})
	currentFiles := getFilesByMate(allFiles)
	return currentFiles
}

/*
根据扩展名判断路径切片中符合的文件
*/
func getFilesByExt(fps []string) (allFiles []string) {
	mkv := []string{
		".mkv",
	}
	for _, fp := range fps {
		for _, item := range mkv {
			if filepath.Ext(fp) == item {
				allFiles = append(allFiles, fp)
			}
		}
	}
	return allFiles
}

/*
根据元数据判断路径切片中符合的文件
*/

func getFilesByMate(fps []string) (allFiles []string) {
	for _, fp := range fps {
		f, err := os.ReadFile(fp)
		if err != nil {
			log.Fatalf("Error opening file:%v\n", err)
		}
		match, err := filetype.Match(f)
		if err != nil {
			return nil
		}
		if match.MIME.Value == "video/x-matroska" {
			allFiles = append(allFiles, fp)
		}
	}
	return allFiles
}
