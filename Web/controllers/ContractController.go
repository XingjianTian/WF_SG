package controllers

import (
	"WF_SG/Web/common"
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"strconv"
)

type ContractController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *ContractController) GetListContract() mvc.View {
	page, err := strconv.Atoi(c.Ctx.URLParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	list, total, totalPages, err := contractModel.ContractList(page)
	if err != nil {
		return common.MvcError(err.Error(), c.Ctx)
	}
	return mvc.View{
		Name: "contract/listContract.html",
		Data: iris.Map{
			"Title":    "List of Contracts",
			"list":     list,
			"PageHtml": common.GetPageHtml(totalPages, page, total, c.Ctx.Path()),
		},
	}

}
func (c *ContractController) GetAddContract() mvc.View {
	return mvc.View{
		Name: "contract/addContract.html",
		Data: iris.Map{
			"Title": "Add Contract",
		},
	}
}
func (c *ContractController) PostAddContract() {
	currentAcc := c.Session.GetString("user_session")
	if err := contractModel.ContractAdd(c.Ctx.FormValues(), "filePath", currentAcc); err == nil {
		c.Ctx.Redirect("/contract/list/contract")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *ContractController) GetQueryInfoBy(contratId string) error {

	contract, _, err := contractModel.ContractInfo(contratId)
	if err != nil {
		return err
	}
	contractJson, err := json.Marshal(contract)
	if err != nil {
		return err
	}
	_, err = c.Ctx.ResponseWriter().Write(contractJson)
	if err != nil {
		return err
	}

	return nil
}
func (c *ContractController) GetQueryMy() error {
	//acc->contractID,through ied
	acc := c.Session.GetString("user_session")
	contract, _, err := contractModel.ContractInfo(acc)
	if err != nil {
		return err
	}
	contractJson, err := json.Marshal(contract)
	if err != nil {
		return err
	}
	_, err = c.Ctx.ResponseWriter().Write(contractJson)
	if err != nil {
		return err
	}

	return nil
}
