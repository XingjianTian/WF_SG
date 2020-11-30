package controllers

import (
	"C"
	ds "WF_SG/Chaincode/DataStructure"
	"WF_SG/Web/common"
	"WF_SG/Web/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

//Controllers call from Models and transfer the data to view

type TableController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

var tableFormatStrBuilderRead string
var tableFormatStrBuilderWrite string
var tableFormatStrConstructorRead string
var tableFormatStrConstructorWrite string
var tableFormatStrSupervisorRead string
var tableFormatStrSupervisorWrite string
var tableFormatStrBuilderManualWrite string

func init() {

	if tableFormatStrBuilderManualWrite == "" {
		tableFormatJson, err := ioutil.ReadFile("jsonformsTables/tableFormat_builder_manual_write.json")
		if err != nil {
			fmt.Print(err.Error())
		}
		tableFormatStrBuilderManualWrite = string(tableFormatJson)
	}

	if tableFormatStrBuilderRead == "" {
		tableFormatJson, err := ioutil.ReadFile("jsonformsTables/tableFormat_builder_read.json")
		if err != nil {
			fmt.Print(err.Error())
		}
		tableFormatStrBuilderRead = string(tableFormatJson)
	}

	if tableFormatStrBuilderWrite == "" {
		tableFormatJson, err := ioutil.ReadFile("jsonformsTables/tableFormat_builder_write.json")
		if err != nil {
			fmt.Print(err.Error())
		}
		tableFormatStrBuilderWrite = string(tableFormatJson)
	}

	if tableFormatStrConstructorRead == "" {
		tableFormatJson, err := ioutil.ReadFile("jsonformsTables/tableFormat_constructor_read.json")
		if err != nil {
			fmt.Print(err.Error())
		}
		tableFormatStrConstructorRead = string(tableFormatJson)
	}

	if tableFormatStrConstructorWrite == "" {
		tableFormatJson, err := ioutil.ReadFile("jsonformsTables/tableFormat_constructor_write.json")
		if err != nil {
			fmt.Print(err.Error())
		}
		tableFormatStrConstructorWrite = string(tableFormatJson)
	}

	if tableFormatStrSupervisorRead == "" {
		tableFormatJson, err := ioutil.ReadFile("jsonformsTables/tableFormat_supervisor_read.json")
		if err != nil {
			fmt.Print(err.Error())
		}
		tableFormatStrSupervisorRead = string(tableFormatJson)
	}

	if tableFormatStrSupervisorWrite == "" {
		tableFormatJson, err := ioutil.ReadFile("jsonformsTables/tableFormat_supervisor_write.json")
		if err != nil {
			fmt.Print(err.Error())
		}
		tableFormatStrSupervisorWrite = string(tableFormatJson)
	}
}

func (c *TableController) Get() mvc.View {

	userOrgName := c.Session.GetString("userOrgName")

	table := models.TableForWeb{}

	list := table.List(userOrgName)

	//models.ListTree = []models.TableForWeb{}
	return mvc.View{
		Name: "Table/list.html",
		Data: iris.Map{
			"Title": "表单列表",
			"list":  list,
		},
	}
}
func (c *TableController) GetSearchTable() mvc.View {

	//models.ListTree = []models.TableForWeb{}
	return mvc.View{
		Name: "Table/searchTable.html",
		Data: iris.Map{
			"Title": "搜索表单",
		},
	}
}

func (c *TableController) GetAddTablefile() mvc.View {
	currentAcc := c.Session.GetString("userOrgName")
	if !strings.Contains(currentAcc, "builder") {
		return common.MvcError("只有建设方可以添加新的表单", c.Ctx)
	}
	return mvc.View{
		Name: "Table/addTablefile.html",
		Data: iris.Map{
			"Title": "从文件新增表单",
		},
	}
}
func (c *TableController) GetAddTablemanually() mvc.View {
	currentAcc := c.Session.GetString("userOrgName")
	if !strings.Contains(currentAcc, "builder") {
		return common.MvcError("只有建设方可以添加新的表单", c.Ctx)
	}

	return mvc.View{
		Name: "Table/addTableManually.html",
		Data: iris.Map{
			"Title":          "手动新增表单",
			"tableFormatStr": tableFormatStrBuilderWrite,
		},
	}
}
func (c *TableController) PostAddTablemanually() {
	table := models.TableForWeb{}
	t := c.Ctx.FormValue("jsonvalues")
	if t == "" {
		fmt.Println("error receiving table manually")
	}
	userid := c.Session.GetString("user_session")
	userOrgName := c.Session.GetString("userOrgName")
	if err := table.AddTable([]byte(t), userid, userOrgName); err == nil {
		c.Ctx.Header("Content-Type", "text/plain")
		_, err := c.Ctx.ResponseWriter().Write([]byte("/Table"))
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		c.Ctx.Header("Content-Type", "text/plain")
		_, err := c.Ctx.ResponseWriter().Write([]byte(err.Error()))
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
func (c *TableController) PostAddTable() mvc.View {
	//table := models.TableForWeb{}
	file, info, err := c.Ctx.FormFile("uploadfile")
	if err != nil {
		c.Ctx.StatusCode(iris.StatusInternalServerError)
		c.Ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return common.MvcError("文件有误", c.Ctx)
	}
	defer file.Close()
	// 创建保存文件
	out, err := os.Create("Web/uploads/" + info.Filename)
	if err != nil {
		c.Ctx.StatusCode(iris.StatusInternalServerError)
		c.Ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return common.MvcError("文件保存失败", c.Ctx)
	}
	defer out.Close()
	io.Copy(out, file)

	tableAsJson, err := ioutil.ReadFile("Web/uploads/" + info.Filename)

	return mvc.View{
		Name: "Table/addTableManually.html",
		Data: iris.Map{
			"Title":          "您选择的表单文件如下所示",
			"tableFormatStr": tableFormatStrBuilderManualWrite,
			"tableValuesStr": string(tableAsJson),
		},
	}

	/*

		userid := c.Session.GetString("admin_user")
		userOrgName := c.Session.GetString("userOrgName")
		if err := table.AddTable(tableAsJson, userid, userOrgName); err == nil {
			c.Ctx.Redirect("/Table")
		} else {
			if strings.Contains(err.Error(), "already exists") {
				common.DefaultErrorShow("该表单已经存在，无法重复添加", c.Ctx)
			} else if strings.Contains(err.Error(), "dose not exist") {
				common.DefaultErrorShow("该表单的前提条件表不存在，无法添加该表单", c.Ctx)
			}
		}

	*/

}

func (c *TableController) GetQueryTableBy(id string) mvc.View {
	table := models.TableForWeb{}

	userOrgName := c.Session.GetString("userOrgName")
	//response is either tableAsJson or error
	response, txID, err := table.QueryTable(id, userOrgName)
	//test := strings.Contains(response,"UserModel@WH-zhijianju.supervisor.com")
	//fmt.Println(test)
	if response == "error" {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}

	block, err := table.QueryBlockByTxID(id, userOrgName, txID)
	ls := table.ListWithoutExclude(userOrgName, id)

	//change id from xxx@yyy.com to xxx
	var tmptable ds.Table
	json.Unmarshal([]byte(response), &tmptable)

	tmptable.TId = ds.AiteBefore(tmptable.TId)

	r, err := json.Marshal(tmptable)
	if err != nil {
		common.DefaultErrorShow(err.Error(), c.Ctx)
	}

	var tableFormatDisabledStr string

	if strings.Contains(userOrgName, "builder") {
		tableFormatDisabledStr = tableFormatStrBuilderRead
	}

	if strings.Contains(userOrgName, "constructor") {
		if strings.Contains(response, "UserModel@WH-zhijianju.supervisor.com") {
			tableFormatDisabledStr = tableFormatStrConstructorWrite
		} else if strings.Contains(response, "UserModel@HUST.builder.com") || strings.Contains(response, "UserModel@zhongjian-1-ju.constructor.com") || strings.Contains(response, "UserModel@zhongjian-2-ju.constructor.com") {
			tableFormatDisabledStr = tableFormatStrConstructorRead
		}
	}

	if strings.Contains(userOrgName, "supervisor") {
		if strings.Contains(response, "UserModel@HUST.builder.com") {
			tableFormatDisabledStr = tableFormatStrSupervisorWrite
		} else if strings.Contains(response, "UserModel@WH-zhijianju.supervisor.com") || strings.Contains(response, "UserModel@zhongjian-1-ju.constructor.com") {
			tableFormatDisabledStr = tableFormatStrSupervisorRead
		}
	}

	return mvc.View{
		Name: "Table/queryTable.html",
		Data: iris.Map{
			"Title": "查询表单",
			//json
			"tableFormatDisabledStr": tableFormatDisabledStr,
			"tableValuesStr":         string(r),
			"tableHistory":           ls,
			"blockNumber":            block.GetHeader().Number,
			"blockCurHash":           base64.StdEncoding.EncodeToString(block.Header.DataHash),
			"blockPreHash":           base64.StdEncoding.EncodeToString(block.Header.PreviousHash),
		},
	}
}
func (c *TableController) PostQueryTable() mvc.View {
	table := models.TableForWeb{}
	tableIDContains := c.Ctx.FormValue("tableIDContains")
	userOrgName := c.Session.GetString("userOrgName")

	allList := table.List(userOrgName)
	var tableWithContainsList []ds.TableForWebinCC

	for _, t := range allList {
		if ds.AiteBefore(t.TID) == tableIDContains {
			tableWithContainsList = append(tableWithContainsList, t)
		}
	}

	return mvc.View{
		Name: "Table/list.html",
		Data: iris.Map{
			"Title": "包含 " + tableIDContains + " 的表单列表",
			"list":  tableWithContainsList,
		},
	}

}

func (c *TableController) PostQueryTablecomplex() mvc.View {

	table := models.TableForWeb{}
	//programNum := c.Ctx.FormValue("programNUm")
	manFill := c.Ctx.FormValue("manFill")
	orgFill := c.Ctx.FormValue("orgFill")
	orgName := c.Ctx.FormValue("orgName")
	depName := c.Ctx.FormValue("depName")
	subName := c.Ctx.FormValue("subName")
	timeFrom := c.Ctx.FormValue("timeFrom")
	timeEnd := c.Ctx.FormValue("timeEnd")

	userOrgName := c.Session.GetString("userOrgName")

	allList := table.List(userOrgName)
	var tableWithContainsList []ds.TableForWebinCC
	for _, t := range allList {
		if manFill != "" && strings.Contains(t.TID, manFill) {
			tableWithContainsList = append(tableWithContainsList, t)
			continue
		}

		if orgFill != "" && strings.Contains(t.OrgEngineeringName, orgFill) {
			tableWithContainsList = append(tableWithContainsList, t)
			continue
		}
		if orgName != "" && strings.Contains(t.OrgEngineeringName, orgName) {
			tableWithContainsList = append(tableWithContainsList, t)
			continue
		}
		if depName != "" && strings.Contains(t.DepEngineeringName, depName) {
			tableWithContainsList = append(tableWithContainsList, t)
			continue
		}
		if subName != "" && strings.Contains(t.SubEngineeringName, subName) {
			tableWithContainsList = append(tableWithContainsList, t)
			continue
		}

		if timeFrom != "" && timeEnd != "" {

			tf, _ := time.Parse("2006-01-02 15:04:05", timeFrom)
			te, _ := time.Parse("2006-01-02 15:04:05", timeEnd)
			if t.CreatedAt.After(tf) && t.CreatedAt.Before(te) {
				tableWithContainsList = append(tableWithContainsList, t)
				continue
			}

		}

		if timeFrom == "" && timeEnd != "" {

			te, _ := time.Parse("2006-01-02 15:04:05", timeEnd)
			if t.CreatedAt.Before(te) {
				tableWithContainsList = append(tableWithContainsList, t)
				continue
			}

		}

		if timeFrom != "" && timeEnd == "" {

			tf, _ := time.Parse("2006-01-02 15:04:05", timeFrom)
			if t.CreatedAt.After(tf) {
				tableWithContainsList = append(tableWithContainsList, t)
				continue
			}

		}

	}

	return mvc.View{
		Name: "Table/list.html",
		Data: iris.Map{
			"Title": "高级查询得到的的表单列表",
			"list":  tableWithContainsList,
		},
	}

}

func bFitSignerOrder(lastSigner, currentSigner string) bool {
	var bSignedOrderRight = false
	if strings.Contains(lastSigner, "builder") && strings.Contains(currentSigner, "supervisor") {
		bSignedOrderRight = true
	}

	if strings.Contains(lastSigner, "supervisor") && strings.Contains(currentSigner, "constructor") {
		bSignedOrderRight = true
	}
	return bSignedOrderRight
}
func (c *TableController) PostAddSigntable() {
	table := models.TableForWeb{}
	values := c.Ctx.FormValue("jsonvalues")
	if values == "" {
		fmt.Println("error receiving table manually")
	}
	userid := c.Session.GetString("user_session")
	userOrgName := c.Session.GetString("userOrgName")

	var t ds.Table
	err := json.Unmarshal([]byte(values), &t)
	if err != nil {
		fmt.Println(err.Error())
	}
	//检查签名顺序
	bSignedOrderRight := bFitSignerOrder(t.LastSigner, userOrgName)
	if !bSignedOrderRight {
		c.Ctx.Header("Content-Type", "text/plain")
		_, err := c.Ctx.ResponseWriter().Write([]byte("你没有权利签名该表"))
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		if err := table.AddTable([]byte(values), userid, userOrgName); err == nil {
			c.Ctx.Header("Content-Type", "text/plain")
			_, err := c.Ctx.ResponseWriter().Write([]byte("/Table"))
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			c.Ctx.Header("Content-Type", "text/plain")
			_, err := c.Ctx.ResponseWriter().Write([]byte(err.Error()))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
