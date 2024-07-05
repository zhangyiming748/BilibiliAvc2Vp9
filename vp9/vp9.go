package vp9

import (
	"BilibiliAvc2Vp9/util"
	"github.com/zhangyiming748/FastMediaInfo"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

/*
* 将mp3文件转为aac文件
接收源文件的绝对路径
*/
func ToVP9(fp string) {
	ext := path.Ext(fp)
	out := strings.Replace(fp, ext, "vp9.mkv", 1)
	mi := FastMediaInfo.GetStandMediaInfo(fp)
	width, _ := strconv.Atoi(mi.Image.Width)
	height, _ := strconv.Atoi(mi.Image.Height)
	frameCount := mi.Video.FrameCount
	log.Printf("文件%v的分辨率是%vx%v,帧数是%v\n", fp, width, height, frameCount)
	crf := FastMediaInfo.GetCRF("vp9", width, height)
	if crf == "" {
		crf = "31"
	}

	if mi.Video.Format == "AVC" {
		log.Printf("文件%v是h264格式,转换\n", fp)
	} else {
		log.Printf("文件%v不是h264格式,无需转换\n", fp)
		return
	}
	cmd := exec.Command("ffmpeg", "-y", "-i", fp, "-map", "0:v:0", "-map", "0:a:0", "-map", "0:s:0", "-c:v", "libvpx-vp9", "-crf", crf, "-c:a", "libopus", "-vbr", "on", "-c:s", "copy", out)
	log.Printf("开始执行命令%v\n", cmd.String())
	err := util.ExecCommand(cmd, frameCount)
	if err != nil {
		log.Printf("命令%v执行失败%v有可能不包括字幕\n", cmd.String(), err)
		cmd = exec.Command("ffmpeg", "-y", "-i", fp, "-map", "0:v:0", "-map", "0:a:0", "-c:v", "libvpx-vp9", "-crf", crf, "-c:a", "libopus", "-vbr", "on", out)
		log.Printf("开始执行命令%v\n", cmd.String())
		if err = util.ExecCommand(cmd, frameCount); err != nil {
			log.Fatalf("命令不带字幕%v执行失败%v\n", cmd.String(), err)
		} else {
			log.Printf("命令成功执行%v\n", cmd.String())
			os.Remove(fp)
		}
	} else {
		log.Printf("命令成功执行%v\n", cmd.String())
		os.Remove(fp)
	}
}
