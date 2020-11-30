package models

import (
	ds "WF_SG/Chaincode/DataStructure"
	"WF_SG/Services"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/jinzhu/gorm"
)

//Models call from HLservices and regularize the data into what controller needs

type TableForWeb struct {
	gorm.Model
	TID   string `gorm:"column:tid;not null;VARCHAR(100);"validate:"required"`
	TName string `gorm:"column:tname;not null;VARCHAR(100);"validate:"required"`
}
type TableID struct {
	TId   string `json:"table_id"`
	TName string `json:"table_name"`
}

func (this *TableForWeb) ListWithoutExclude(userOrgName, id string) []ds.TableForWebinCC {

	respJson, err := Services.HLservice.QueryAllTablesWithoutExclude(userOrgName)
	//transform
	var ts []ds.TableForWebinCC
	err = json.Unmarshal([]byte(respJson), &ts)
	if err != nil {
		fmt.Print(err.Error())
	}

	var tableWithContainsList []ds.TableForWebinCC

	for _, t := range ts {
		if t.RealID == ds.AiteBefore(id) {
			tableWithContainsList = append(tableWithContainsList, t)
			continue
		}
	}

	for index := range tableWithContainsList { //获取索引
		tableWithContainsList[index].CreatedTimeWeb =
			tableWithContainsList[index].CreatedAt.Format("2006-01-02 15:04:05")
	}

	/*
		var data []TableForWeb
		db :=libs.DB
		err := db.Order("id").Find(&data).Error
		if err != nil {
			//log.Fatalln(err)
		}

	*/

	return tableWithContainsList
}

func (this *TableForWeb) List(userOrgName string) []ds.TableForWebinCC {

	respJson, err := Services.HLservice.QueryAllTables(userOrgName)
	//transform
	var ts []ds.TableForWebinCC
	err = json.Unmarshal([]byte(respJson), &ts)
	if err != nil {
		fmt.Print(err.Error())
	}
	for index := range ts { //获取索引
		ts[index].CreatedTimeWeb = ts[index].CreatedAt.Format("2006-01-02 15:04:05") //通过下标获取元素进行修改
	}

	return ts
}

func (this *TableForWeb) AddTable(tableAsJson []byte, userid string, userOrgName string) error {

	//var err error
	res, err := Services.HLservice.AddTableService(tableAsJson, userid, userOrgName)
	if err != nil {
		return err
	}
	println(res)
	/*
		var ti TableID
		err = json.Unmarshal(tableAsJson,&ti)


		db := libs.DB

		var tableForWeb TableForWeb
		//add table to mysql
		tableForWeb.TID = ti.TId+"@"+userOrgName
		tableForWeb.TName = ti.TName

		if !db.Where("tid = ? ", ti.TId).First(&tableForWeb).RecordNotFound() {
			return errors.New("table id"+ti.TId+" already exits")
		}


		if err := db.Create(&tableForWeb).Error; err != nil {
			return errors.New("fail to add")
		}

	*/
	return nil
}
func (this *TableForWeb) QueryTable(tableID string, userOrgName string) (string, fab.TransactionID, error) {

	res, txID, err := Services.HLservice.SearchTableByIdService(tableID, userOrgName)
	if err != nil {
		return "error", "", err
	}
	return res, txID, nil
	//query through chaincode
	//show table
}

func (this *TableForWeb) QueryBlockByTxID(tableID string, userOrgName string, txID fab.TransactionID) (*common.Block, error) {

	block, err := Services.HLservice.QueryBlockByTx(txID)
	if err != nil {
		return nil, err
	}
	return block, nil
}
