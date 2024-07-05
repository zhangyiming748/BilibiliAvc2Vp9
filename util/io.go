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
			if isMkv(path) {
				allFiles = append(allFiles, path)
			}
		}
		return nil
	})
	return allFiles
}

func isMkv(fp string) bool {
	f, err := os.ReadFile(fp)
	if err != nil {
		log.Fatalf("Error opening file:%v\n", err)
	}
	match, err := filetype.Match(f)
	if err != nil {
		return false
	}
	if match.MIME.Value == "video/x-matroska" {
		return true
	}
	return false
}
