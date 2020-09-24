package controllers

import (
	"WF_SG/SDKInit"
	"WF_SG/Web/models"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

var userModel models.UserModel
var companyModel models.CompanyModel
var wgModel models.WgModel

type SystemController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *SystemController) GetMain() mvc.View {
	err := SDKInit.SetupInitInfo(channelIdSelected)
	if err != nil {
		fmt.Println("Failed in sdk setup" + err.Error())
	}

	return mvc.View{
		Name: "system/main.html",
	}
}
