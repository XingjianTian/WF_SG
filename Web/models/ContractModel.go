package models

import (
	"WF_SG/Services"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/jinzhu/gorm"
	config "github.com/spf13/viper"
	"math"
)

type ContractModel struct {
	gorm.Model
	ContractId      string `gorm:"primary_key;type:varchar(50);not null;" json:"contract_id"`
	ContractVersion string `gorm:"type:varchar(50);not null;" json:"contract_version"`
	//fabric key:id+"-"+version
	ContractName                string `gorm:"type:varchar(50);not null;" json:"contract_name"`
	ContractCompanyName         string `gorm:"type:varchar(50);not null;" json:"contract_company_name"`
	ContractCompanyOwnerAccount string `gorm:"type:varchar(50);not null;" json:"contract_company_owner_account"`
	ContractCompanyOwnerSig     string `gorm:"type:varchar(255);not null;" json:"contract_company_owner_sig"`
	ContractDetails             string `gorm:"type:varchar(50);not null;" json:"contract_details"`
	EnergyType                  string `gorm:"type:varchar(50);not null;" json:"energy_type"`
	EnergyPrice                 string `gorm:"type:varchar(50);not null;" json:"energy_price"`
	ContractLastTime            string `gorm:"type:varchar(50);not null;" json:"contract_last_time"`

	ContractSignTime    string `gorm:"type:varchar(50);not null;" json:"contract_last_time"`
	ContractUserAccount string `gorm:"type:varchar(50);not null;" json:"contract_user_account"`
	ContractUserSig     string `gorm:"type:varchar(50);not null;" json:"contract_user_sig"`
}

func (this *ContractModel) TableName() string {
	return "contract"
}
func (this *ContractModel) ContractKey() string {
	return this.ContractId + "-" + this.ContractVersion + "-" + this.ContractUserAccount
}

func (this *ContractModel) ContractInfo(contractId string) (ContractModel, fab.TransactionID, error) {

	res, txID, err := Services.HLservice.QueryContractByKeyService(contractId)
	if err != nil {
		return ContractModel{}, "", err
	}
	var contract ContractModel
	_ = json.Unmarshal([]byte(res), &contract)

	return contract, txID, nil
}
func (this *ContractModel) ContractList(page int) ([]ContractModel, int, int, error) {
	var data []ContractModel
	var totalCount int
	limit := config.GetInt("pagination.PageSize")
	respJson, err := Services.HLservice.QueryAllContractsService()
	//transform
	err = json.Unmarshal([]byte(respJson), &data)
	if err != nil {
		fmt.Print(err.Error())
	}
	totalCount = len(data)
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))

	return data, totalCount, totalPages, err
}
func (this *ContractModel) ContractAdd(userAcc string, filePath string, bidId string) error {

	var bid BidModel
	bid, err := bid.BidInfo(bidId)
	if err != nil {
		return err
	}

	contractJson, err := json.Marshal(&bid)
	if err != nil {
		return err
	}
	//var err error
	_, err = Services.HLservice.AddContractService(contractJson)
	if err != nil {
		return err
	}
	return nil
}
