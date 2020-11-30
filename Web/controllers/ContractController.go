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
	bidId := c.Ctx.FormValue("bidId")
	currentAcc := c.Session.GetString("user_session")
	if err := contractModel.ContractAdd(currentAcc, "filePath", bidId); err == nil {
		c.Ctx.Redirect("/contract/list/contract")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *ContractController) GetQueryContractBy(contratKey string) error {

	contract, _, err := contractModel.ContractInfo(contratKey)
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

//bid
func (c *ContractController) GetListBid() mvc.View {
	page, err := strconv.Atoi(c.Ctx.URLParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	list, total, totalPages := bidModel.BidList(page)
	return mvc.View{
		Name: "contract/listBid.html",
		Data: iris.Map{
			"Title":    "List of Bids",
			"list":     list,
			"PageHtml": common.GetPageHtml(totalPages, page, total, c.Ctx.Path()),
		},
	}

}
func (c *ContractController) GetAddBid() mvc.View {
	return mvc.View{
		Name: "contract/addBid.html",
		Data: iris.Map{
			"Title": "Add Bid",
		},
	}
}

func (c *ContractController) PostAddBid() {
	currentAcc := c.Session.GetString("user_session")
	if err := bidModel.BidAdd(c.Ctx.FormValues(), currentAcc); err == nil {
		c.Ctx.Redirect("/contract/list/bid")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *ContractController) GetUpdateBidBy(bidId string) mvc.View {
	currentAcc := c.Session.GetString("user_session")
	bid, err := bidModel.BidInfo(bidId)
	if err != nil {
		return common.MvcError(err.Error(), c.Ctx)
	}

	if currentAcc != bid.ContractCompanyOwnerAccount {
		return common.MvcError("Authentication Failed", c.Ctx)
	}
	return mvc.View{
		Name: "contract/updateBid.html",
		Data: iris.Map{
			"Title":   "Update Bid Info",
			"bidInfo": bid,
		},
	}
}
func (c *ContractController) PostUpdateBid() {
	var postValues map[string][]string
	postValues = c.Ctx.FormValues()
	if err := bidModel.BidUpdate(postValues, postValues["contractId"][0]); err == nil {
		c.Ctx.Redirect("/system/main")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}

func (c *ContractController) GetQueryBidBy(bidId string) error {

	bidInfo, err := bidModel.BidInfo(bidId)
	if err != nil {
		return err
	}

	bidJson, err := json.Marshal(bidInfo)
	if err != nil {
		return err
	}

	_, err = c.Ctx.ResponseWriter().Write(bidJson)
	if err != nil {
		return err
	}
	return nil
}
func (c *ContractController) GetDelBidBy(bidId string) {
	currentAcc := c.Session.GetString("user_session")
	bid, err := bidModel.BidInfo(bidId)
	if err != nil {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
	if currentAcc != bid.ContractCompanyOwnerAccount {
		common.DefaultErrorShow("Authentication Failed", c.Ctx)
		return
	}
	if err := bidModel.BidDel(bidId); err == nil {
		c.Ctx.Redirect("/contract/list/bid")
	} else {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}
}
