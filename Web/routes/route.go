package routes

import (
	"WF_SG/Web/common"
	"WF_SG/Web/controllers"
	"WF_SG/Web/middlewares"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func Routes(app *iris.Application) {

	//index
	mvc.New(app.Party("/")).
		Register(common.SessManager.Start).
		Handle(new(controllers.IndexController))

	//login
	mvc.New(app.Party("/login")).
		Register(common.SessManager.Start).
		Handle(new(controllers.LoginController))

	//system
	mvc.New(app.Party("/system", middlewares.SessionLoginAuth)).
		Register(common.SessManager.Start).
		Handle(new(controllers.SystemController))

	//user management
	mvc.New(app.Party("/user", middlewares.SessionLoginAuth)).
		Register(common.SessManager.Start).
		Handle(new(controllers.UserController))

	//表单管理
	mvc.New(app.Party("/Table", middlewares.SessionLoginAuth)).
		Register(common.SessManager.Start).
		Handle(new(controllers.TableController))

	//ied management
	mvc.New(app.Party("/ied", middlewares.SessionLoginAuth)).
		Register(common.SessManager.Start).
		Handle(new(controllers.IedController))

	//company management
	mvc.New(app.Party("/company", middlewares.SessionLoginAuth)).
		Register(common.SessManager.Start).
		Handle(new(controllers.CompanyController))

	//wireguard management
	mvc.New(app.Party("/wg", middlewares.SessionLoginAuth)).
		Register(common.SessManager.Start).
		Handle(new(controllers.WgController))

	//contract management
	mvc.New(app.Party("/contract", middlewares.SessionLoginAuth)).
		Register(common.SessManager.Start).
		Handle(new(controllers.ContractController))
}
