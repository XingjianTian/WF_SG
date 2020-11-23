package controllers

import (
	"WF_SG/Web/common"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"strconv"
)

type CompanyController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *CompanyController) GetListCompany() mvc.View {
	page, err := strconv.Atoi(c.Ctx.URLParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	list, total, totalPages := companyModel.CompanyList(page)
	return mvc.View{
		Name: "company/listCompany.html",
		Data: iris.Map{
			"Title":    "List of Companies",
			"list":     list,
			"PageHtml": common.GetPageHtml(totalPages, page, total, c.Ctx.Path()),
		},
	}

}
func (c *CompanyController) GetUpdateCompanyBy(companyName string) mvc.View {
	currentAcc := c.Session.GetString("user_session")
	companyInfo, err := companyModel.CompanyInfo(companyName)
	if err != nil {
		return common.MvcError(err.Error(), c.Ctx)
	}

	if currentAcc != companyInfo.CompanyOwnerAccount {
		return common.MvcError("can not update others' company info", c.Ctx)
	}
	return mvc.View{
		Name: "company/updateCompany.html",
		Data: iris.Map{
			"Title":       "Update Company Info",
			"companyInfo": companyInfo,
		},
	}
}
func (c *CompanyController) PostUpdateCompany() {
	_, filePath := common.UploadFile("headico", c.Ctx)
	/*
		if err == false {
			common.DefaultErrorShow(filePath, c.Ctx)
			return
		}
	*/

	var postValues map[string][]string
	postValues = c.Ctx.FormValues()
	if err := companyModel.CompanyUpdate(postValues, postValues["companyName"][0], filePath); err == nil {
		c.Ctx.Redirect("/system/main")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}
func (c *CompanyController) GetAddCompany() mvc.View {
	return mvc.View{
		Name: "company/addCompany.html",
		Data: iris.Map{
			"Title": "Add Company",
		},
	}
}
func (c *CompanyController) PostAddCompany() {
	currentAcc := c.Session.GetString("user_session")
	err, filePath := common.UploadFile("headico", c.Ctx)
	if err == false {
		common.DefaultErrorShow(filePath, c.Ctx)
		return
	}
	if err := companyModel.CompanyAdd(c.Ctx.FormValues(), filePath, currentAcc); err == nil {
		c.Ctx.Redirect("/company/list/company")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *CompanyController) GetDelCompanyBy(companyName string) {
	currentUserType := c.Session.GetString("user_type")
	if currentUserType != "admin" {
		common.DefaultErrorShow("only admin can delete", c.Ctx)
		return
	}
	if err := companyModel.CompanyDel(companyName); err == nil {
		c.Ctx.Redirect("/company/list/company")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *CompanyController) GetQueryCompany() mvc.View {

	currentAcc := c.Session.GetString("user_session")
	page, err := strconv.Atoi(c.Ctx.URLParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	list, total, totalPages := companyModel.CompanyInfoByOwnerAcc(currentAcc, page)
	return mvc.View{
		Name: "company/listCompany.html",
		Data: iris.Map{
			"Title":    "List of Companies You own",
			"list":     list,
			"PageHtml": common.GetPageHtml(totalPages, page, total, c.Ctx.Path()),
		},
	}

}
func (c *CompanyController) GetQueryCompanyBy(companyName string) mvc.View {

	companyInfo, err := companyModel.CompanyInfo(companyName)
	if err != nil {
		return common.MvcError(err.Error(), c.Ctx)
	}
	return mvc.View{
		Name: "company/queryCompany.html",
		Data: iris.Map{
			"Title":       "Company Details",
			"companyInfo": companyInfo,
		},
	}
}
