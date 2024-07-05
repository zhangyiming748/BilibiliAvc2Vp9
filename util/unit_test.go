package util

import (
	"BilibiliAvc2Vp9/constant"
	"fmt"
	"log"
	"testing"
)

func TestAllFiles(t *testing.T) {
	p := &constant.Param{
		Root: "D:\\bilibili\\张栋梁",
	}
	fps := GetAllfiles(p)
	for _, fp := range fps {
		log.Printf("根据mime找到mkv文件%+v\n", fp)
		for _, c := range fp {
			fmt.Printf("%U\t%c\n", c, c)
		}
	}

}
