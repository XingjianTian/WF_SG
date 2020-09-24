package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type IndexController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *IndexController) Get() mvc.View {
	return mvc.View{
		Name:   "index/index.html",
		Layout: "shared/layoutFront.html",
		Data:   iris.Map{"Title": "index"},
	}
}
