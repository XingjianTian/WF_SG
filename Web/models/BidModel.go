package models

import (
	"WF_SG/Chaincode/Utils"
	"WF_SG/Web/common"
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	config "github.com/spf13/viper"
	"log"
	"math"
)

type BidModel struct {
	gorm.Model
	ContractId      string `gorm:"primary_key;type:varchar(50);not null;" json:"contract_id"`
	ContractVersion string `gorm:"type:varchar(50);not null;" json:"contract_version"`
	//fabric key:id+"-"+version
	ContractName                string `gorm:"type:varchar(50);not null;" json:"contract_name"`
	ContractCompanyName         string `gorm:"type:varchar(50);not null;" json:"contract_company_name"`
	ContractCompanyOwnerAccount string `gorm:"type:varchar(50);not null;" json:"contract_company_owner_account"`
	ContractCompanyOwnerSig     string `gorm:"type:varchar(255);not null;" json:"contract_company_owner_sig"`
	ContractDetails             string `gorm:"type:varchar(255);not null;" json:"contract_details"`
	EnergyType                  string `gorm:"type:varchar(50);not null;" json:"energy_type"`
	EnergyPrice                 string `gorm:"type:varchar(50);not null;" json:"energy_price"`
	ContractLastTime            string `gorm:"type:varchar(50);not null;" json:"contract_last_time"`
}

func (this *BidModel) TableName() string {
	return "bid"
}
func (this *BidModel) BidList(page int) ([]BidModel, int, int) {
	var data []BidModel
	var totalCount int
	limit := config.GetInt("pagination.PageSize")
	offset := (page - 1) * limit
	db := common.DB
	db.Find(&data).Count(&totalCount)
	err := db.Offset(offset).Limit(limit).Order("contract_id").Find(&data).Error
	if err != nil {
		log.Fatalln(err)
	}
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))
	return data, totalCount, totalPages
}

func (this *BidModel) BidInfo(bidId string) (BidModel, error) {
	var bid BidModel
	db := common.DB
	if db.Where("contract_id = ?", bidId).First(&bid).RecordNotFound() {
		return BidModel{}, errors.New("bid not found")
	}
	return bid, nil
}

func (this *BidModel) BidUpdate(postValues map[string][]string, bidId string) error {
	db := common.DB
	bid, err := this.BidInfo(bidId)
	if err != nil {
		return err
	}
	bid.ContractVersion = postValues["contractVersion"][0]
	//fabric key:id+"-"+version
	bid.ContractName = postValues["contractName"][0]
	bid.ContractCompanyName = postValues["contractCompany"][0]
	bid.ContractDetails = postValues["contractDetails"][0]
	bid.EnergyType = postValues["energyType"][0]
	bid.EnergyPrice = postValues["energyPrice"][0]
	bid.ContractLastTime = postValues["companyLast"][0]

	bidJson, err := json.Marshal(bid)
	if err != nil {
		return nil
	}

	bid.ContractCompanyOwnerSig, err = Utils.Sign(bidJson, bid.ContractCompanyOwnerAccount)
	if err != nil {
		return nil
	}

	if err != nil {
		return err
	}
	if err := db.Save(&bid).Error; err != nil {
		return errors.New("Failed to update bid")
	}

	return nil
}
func (this *BidModel) BidAdd(postValues map[string][]string, userAcc string) error {
	var bid BidModel

	bid.ContractId = postValues["contractId"][0]
	bid.ContractVersion = postValues["contractVersion"][0]
	//fabric key:id+"-"+version
	bid.ContractName = postValues["contractName"][0]
	bid.ContractCompanyName = postValues["contractCompany"][0]
	bid.ContractCompanyOwnerAccount = userAcc
	bid.ContractDetails = postValues["contractDetails"][0]
	bid.EnergyType = postValues["energyType"][0]
	bid.EnergyPrice = postValues["energyPrice"][0]
	bid.ContractLastTime = postValues["companyLast"][0]

	bidJson, err := json.Marshal(bid)
	if err != nil {
		return nil
	}

	bid.ContractCompanyOwnerSig, err = Utils.Sign(bidJson, userAcc)
	if err != nil {
		return nil
	}

	db := common.DB

	if !db.Where("contract_id = ? ", bid.ContractId).First(&bid).RecordNotFound() {
		return errors.New("bid already exits")
	}

	if err := db.Create(&bid).Error; err != nil {
		return errors.New("Failed to add bid")
	}

	return nil
}

func (this *BidModel) BidDel(bidId string) error {
	var bid BidModel
	db := common.DB
	if err := db.Where("contract_id = ?", bidId).Delete(&bid).Error; err != nil {
		return errors.New("Failed to delete bid")
	}
	return nil
}
