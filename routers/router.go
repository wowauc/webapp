package routers

import (
	"github.com/wowauc/gowowuction/apps/webapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
