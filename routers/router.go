package routers

import (
	"faceRecognition/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/play", &controllers.Playcontroller{})
}
