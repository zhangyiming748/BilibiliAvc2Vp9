package main

import (
	"BilibiliAvc2Vp9/constant"
	mylog "BilibiliAvc2Vp9/log"
	"BilibiliAvc2Vp9/util"
	"BilibiliAvc2Vp9/vp9"
	"os"
)

func main() {
	p := new(constant.Param)
	p.SetRoot("D:\\bilibili\\星光明007")
	if root := os.Getenv("root"); root != "" {
		p.SetRoot(root)
	}
	mylog.SetLog(p)
	files := util.GetAllfiles(p)
	for _, f := range files {
		vp9.ToVP9(f)
	}
}
