package util

import (
	"BilibiliAvc2Vp9/constant"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	p := &constant.Param{
		Root: "D:\\bilibili\\张栋梁",
	}
	fps := GetAllfiles(p)
	log.Printf("根据mime找到mkv文件%+v\n", fps)
}
