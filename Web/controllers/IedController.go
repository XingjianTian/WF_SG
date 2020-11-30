package controllers

import (
	"WF_SG/Services"
	"WF_SG/Web/common"
	"encoding/base64"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"strconv"
)

type IedController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *IedController) GetListIed() mvc.View {

	page, err := strconv.Atoi(c.Ctx.URLParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	list, total, totalPages, err := iedModel.IedList(page)
	if err != nil {
		return common.MvcError(err.Error(), c.Ctx)
	}
	return mvc.View{
		Name: "ied/listIED.html",
		Data: iris.Map{
			"Title":    "List of IEDs",
			"list":     list,
			"PageHtml": common.GetPageHtml(totalPages, page, total, c.Ctx.Path()),
		},
	}

}

func (c *IedController) GetAddIed() mvc.View {
	return mvc.View{
		Name: "ied/addIED.html",
		Data: iris.Map{
			"Title": "Add ied",
		},
	}
}
func (c *IedController) PostAddIed() {

	currentAcc := c.Session.GetString("user_session")
	if err := iedModel.IedAdd(c.Ctx.FormValues(), currentAcc); err == nil {
		c.Ctx.Redirect("/ied/list/ied")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *IedController) GetQueryIedBy(id string) mvc.View {
	//userOrgName := c.Session.GetString("userOrgName")
	//currentAcc := c.Session.GetString("user_session")

	iedInfo, txID, err := iedModel.IedInfo(id)
	if err != nil {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}

	block, err := Services.HLservice.QueryBlockByTx(txID)

	if err != nil {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}

	return mvc.View{
		Name: "ied/queryIED.html",
		Data: iris.Map{
			"Title": "ied Info",
			//json
			"iedInfo":      iedInfo,
			"blockNumber":  block.GetHeader().Number,
			"blockCurHash": base64.StdEncoding.EncodeToString(block.Header.DataHash),
			"blockPreHash": base64.StdEncoding.EncodeToString(block.Header.PreviousHash),
		},
	}
}

func (c *IedController) GetUpdateIedBy(id string) mvc.View {
	//currentAcc := c.Session.GetString("user_session")
	iedInfo, _, err := iedModel.IedInfo(id)
	if err != nil {
		return common.MvcError(err.Error(), c.Ctx)
	}

	currentUserType := c.Session.GetString("user_type")
	if currentUserType != "admin" {
		return common.MvcError("only admin can update", c.Ctx)
	}

	return mvc.View{
		Name: "ied/updateIed.html",
		Data: iris.Map{
			"Title":   "Update Ied Info",
			"iedInfo": iedInfo,
		},
	}
}
func (c *IedController) PostUpdateIed() {
	var postValues map[string][]string
	postValues = c.Ctx.FormValues()
	if err := iedModel.IedUpdate(postValues); err == nil {
		c.Ctx.Redirect("/system/main")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}
