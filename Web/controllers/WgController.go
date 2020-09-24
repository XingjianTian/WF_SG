package controllers

import (
	"WF_SG/Web/common"
	"WF_SG/Web/models"
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"strconv"
)

type WgController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *WgController) GetListWg() mvc.View {
	/*
		wgModel := models.WgModel{}
		page, err := strconv.Atoi(c.Ctx.URLParam("page"))
		if err != nil || page < 1 {
			page = 1
		}
		list, total, totalPages := wgModel.WgList(page)
	*/
	return mvc.View{
		Name: "wg/listWg.html",
		Data: iris.Map{
			"Title": "List of WireGuard Clients",
			//"clientDataList":     list,
			//"PageHtml": common.GetPageHtml(totalPages, page, total, c.Ctx.Path()),
		},
	}

}
func (c *WgController) GetListWgajax() error {
	wgModel := models.WgModel{}
	page, err := strconv.Atoi(c.Ctx.URLParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	list, _, _ := wgModel.WgList(page)
	listJson, err := json.Marshal(list)
	if err != nil {
		return err
	}
	_, err = c.Ctx.ResponseWriter().Write(listJson)
	if err != nil {
		return err
	}
	return nil

}

func (c *WgController) PostAddWg() {
	wgJson := c.Ctx.FormValue("jsonvalues")
	if err := wgModel.WgAdd([]byte(wgJson)); err == nil {
		c.Ctx.Redirect("/wg/list/wgajax")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *WgController) PostUpdateStatus() error {
	wgStatusJson := c.Ctx.FormValue("jsonvalues")
	var wgStatus models.WgModel
	err := json.Unmarshal([]byte(wgStatusJson), &wgStatus)
	if err != nil {
		return err
	}

	acc := wgStatus.Account
	if err := wgModel.WgUpdateStatus(acc, wgStatus.Enabled); err == nil {
		c.Ctx.Redirect("/wg/list/wg")
		return nil
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}

	return nil
}

func (c *WgController) GetQueryInfoBy(acc string) error {
	wg, err := wgModel.WgInfo(acc)
	if err != nil {
		return err
	}
	wgJson, err := json.Marshal(wg)
	if err != nil {
		return err
	}
	_, err = c.Ctx.ResponseWriter().Write(wgJson)
	if err != nil {
		return err
	}

	return nil
}

func (c *WgController) GetQueryMy() error {
	acc := c.Session.GetString("user_session")
	wg, err := wgModel.WgInfo(acc)
	if err != nil {
		return err
	}
	wgJson, err := json.Marshal(wg)
	if err != nil {
		return err
	}
	_, err = c.Ctx.ResponseWriter().Write(wgJson)
	if err != nil {
		return err
	}

	return nil
}

func (c *WgController) PostUpdateInfoBy() error {
	return nil
}

func (c *WgController) PostDelWg() error {
	acc := c.Ctx.FormValue("client-acc")
	currentUserType := c.Session.GetString("user_type")
	if currentUserType != "admin" {
		common.DefaultErrorShow("only admin can delete wg client", c.Ctx)
	}
	if err := wgModel.WgDel(acc); err == nil {
		c.Ctx.Redirect("/wg/list/wg")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}

	return nil
}

func (c *WgController) GetSuggestedAllocated() error {

	res, err := wgModel.IpList("allocated", 2)
	if err != nil {
		return err
	}
	_, err = c.Ctx.ResponseWriter().Write(res)
	if err != nil {
		return err
	}
	return nil
}
