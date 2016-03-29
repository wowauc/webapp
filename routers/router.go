package routers

import (
	"github.com/astaxie/beego"
	"github.com/wowauc/webapp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
