package controllers

import (
	"github.com/astaxie/beego"
	"os/exec"
	"strconv"
)

type Playcontroller struct {
	beego.Controller
}

var videoStringUrl  = "rtsp://admin:123qweasd@192.168.1.233:554/video1"
func (c *Playcontroller) Get() {
	//拉流
	ffmeg := "localhost:8080"
	params := []string{"-fflags", "genpts", "-rtsp_transport", "tcp", "-i", videoStringUrl, "-hls_time", strconv.Itoa(123), "-hls_list_size", "0", "123"}
	cmd :=exec.Command(ffmeg,params...)

}
