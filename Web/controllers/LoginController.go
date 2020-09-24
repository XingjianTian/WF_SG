package controllers

import (
	ds "WF_SG/DataStructure"
	"WF_SG/Web/common"
	"WF_SG/Web/models"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"html"
	"strings"
	"sync"
)

type LoginController struct {
	//不能ctx重名
	Ctx     iris.Context
	Session *sessions.Session
	mu      sync.Mutex
}

var channelIdSelected string

func (c *LoginController) Get() {

	if auth := common.SessManager.Start(c.Ctx).Get("user_session"); auth != nil {
		c.Ctx.Redirect("/system/main")
	} else {
		c.Ctx.Redirect("/login/show")
	}
}
func (c *LoginController) GetShow() mvc.View {
	err := c.Ctx.URLParam("err")
	return mvc.View{
		Name:   "index/login.html",
		Layout: "shared/layoutNone.html",
		Data: iris.Map{
			"Title": "Login",
			"err":   err,
		},
	}
}
func (c *LoginController) Post() {
	var userModel models.UserModel
	acc := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("account")))
	pwd := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("password")))
	chl := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("channelId")))
	channelIdSelected = chl
	if userInfo, err := userModel.UserLogin(acc, pwd); err == nil {
		/*
			sessionInfo:=make(map[string]interface{})
			sessionInfo["id"]=userInfo.ID
			sessionInfo["name"]=userInfo.Account
		*/
		//save session

		userOrgName := ds.AiteAfter(userInfo.Account)
		c.Session.Set("user_session", userInfo.Account)
		c.Session.Set("user_type", userInfo.Type)
		c.Session.Set("userOrgName", userOrgName)
		c.Ctx.Redirect("/system/main")
		//c.Session.Set("channel",chl)

	} else {
		c.Ctx.Redirect("/login/show?err=" + err.Error())
	}
}
func (c *LoginController) GetLogout() {
	c.Session.Delete("user_session")
	c.Ctx.Redirect("/login")
}
