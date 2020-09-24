package controllers

import (
	"WF_SG/Web/common"
	"WF_SG/Web/models"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"html"
	"strconv"
	"strings"
)

type UserController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *UserController) GetListUser() mvc.View {
	userModel := models.UserModel{}
	page, err := strconv.Atoi(c.Ctx.URLParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	list, total, totalPages := userModel.UserList(page)
	return mvc.View{
		Name: "user/listUser.html",
		Data: iris.Map{
			"Title":    "List of Users",
			"list":     list,
			"PageHtml": common.GetPageHtml(totalPages, page, total, c.Ctx.Path()),
		},
	}

}

func (c *UserController) PostUpdateUser() {
	_, filePath := common.UploadFile("headico", c.Ctx)
	/*
		if err == false {
			common.DefaultErrorShow(filePath, c.Ctx)
			return
		}
	*/

	var postValues map[string][]string
	postValues = c.Ctx.FormValues()
	acc := postValues["account"][0]
	if err := userModel.UserUpdate(postValues, acc, filePath); err == nil {
		c.Ctx.Redirect("/system/main")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}
func (c *UserController) GetUpdateUser() mvc.View {

	currentAcc := c.Session.GetString("user_session")
	userInfo, err := userModel.UserInfo(currentAcc)
	if err != nil {
		return common.MvcError(err.Error(), c.Ctx)
	}
	return mvc.View{
		Name: "user/updateUser.html",
		Data: iris.Map{
			"Title":    "Update User Info",
			"userInfo": userInfo,
		},
	}
}
func (c *UserController) GetUpdateUserBy(acc string) mvc.View {

	currentUserType := c.Session.GetString("user_type")
	userToUpdateInfo, err := userModel.UserInfo(acc)
	if err != nil {
		return common.MvcError(err.Error(), c.Ctx)
	}
	if currentUserType != "admin" {
		return common.MvcError(err.Error(), c.Ctx)
	} else if userToUpdateInfo.Type == "admin" && userToUpdateInfo.Account != acc {
		return common.MvcError(err.Error(), c.Ctx)
	}
	return mvc.View{
		Name: "user/updateUser.html",
		Data: iris.Map{
			"Title":    "Update User Info",
			"userInfo": userToUpdateInfo,
		},
	}
}

func (c *UserController) GetAddUser() mvc.View {
	return mvc.View{
		Name: "user/addUser.html",
		Data: iris.Map{
			"Title": "Add User",
		},
	}
}
func (c *UserController) PostAddUser() {
	err, filePath := common.UploadFile("headico", c.Ctx)
	if err == false {
		common.DefaultErrorShow(filePath, c.Ctx)
		return
	}
	if err := userModel.UserAdd(c.Ctx.FormValues(), filePath); err == nil {
		c.Ctx.Redirect("/user/list/user")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *UserController) GetUpdatePasswordBy(acc string) mvc.View {
	currentAcc := c.Session.GetString("user_session")
	if acc != currentAcc {
		return common.MvcError("can not update others' password", c.Ctx)
	}
	userInfo, err := userModel.UserInfo(acc)
	if err != nil {
		return common.MvcError(err.Error(), c.Ctx)
	}
	return mvc.View{
		Name: "user/updatePassword.html",
		Data: iris.Map{
			"Title":   "Update Password",
			"Account": userInfo.Account,
		},
	}
}
func (c *UserController) PostUpdatePassword() {
	acc := c.Session.GetString("user_session")
	//id := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("id")))
	password := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("password")))
	Repassword := html.EscapeString(strings.TrimSpace(c.Ctx.FormValue("Repassword")))
	//int_admin_id, _ := strconv.Atoi(id)
	if err := userModel.UserPasswdUpdate(acc, password, Repassword); err == nil {
		c.Ctx.Redirect("/system/main")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}
func (c *UserController) GetUpdatePassword() mvc.View {
	return mvc.View{
		Name: "user/updatePassword.html",
		Data: iris.Map{
			"Title": "update password",
		},
	}
}

func (c *UserController) GetDelUserBy(acc string) {
	currentUserType := c.Session.GetString("user_type")
	userToDelInfo, err := userModel.UserInfo(acc)
	if err != nil {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
	if currentUserType != "admin" {
		common.DefaultErrorShow("only admin can delete", c.Ctx)
		return
	} else if userToDelInfo.Type == "admin" {
		common.DefaultErrorShow("admin can not be deleted", c.Ctx)
		return
	}
	if err := userModel.UserDel(acc); err == nil {
		c.Ctx.Redirect("/user/list/user")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *UserController) GetQueryUserBy(acc string) mvc.View {
	userInfo, err := userModel.UserInfo(acc)
	if err != nil {
		return common.MvcError(err.Error(), c.Ctx)
	}
	return mvc.View{
		Name: "user/queryUser.html",
		Data: iris.Map{
			"Title":    "User Details",
			"userInfo": userInfo,
		},
	}
}
func (c *UserController) GetQueryUser() mvc.View {
	currentAcc := c.Session.GetString("user_session")
	userInfo, err := userModel.UserInfo(currentAcc)
	if err != nil {
		return common.MvcError(err.Error(), c.Ctx)
	}
	return mvc.View{
		Name: "user/queryUser.html",
		Data: iris.Map{
			"Title":    "User Details",
			"userInfo": userInfo,
		},
	}
}
