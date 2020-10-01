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
	ContractName        string `gorm:"type:varchar(50);not null;" json:"contract_name"`
	ContractCompanyName string `gorm:"type:varchar(50);not null;" json:"contract_company_name"`
	ContractCompanySig  string `gorm:"type:varchar(50);not null;" json:"contract_company_sig"`
	ContractDetails     string `gorm:"type:varchar(50);not null;" json:"contract_details"`
	EnergyType          string `gorm:"type:varchar(50);not null;" json:"energy_type"`
	EnergyPrice         string `gorm:"type:varchar(50);not null;" json:"energy_price"`
	ContractLastTime    string `gorm:"type:varchar(50);not null;" json:"contract_last_time"`
}

func (this *ContractModel) TableName() string {
	return "contract"
}

func (this *ContractModel) ContractInfo(contractId string) (ContractModel, fab.TransactionID, error) {

	res, txID, err := Services.HLservice.QueryContractByIdService(contractId)
	if err != nil {
		return ContractModel{}, "", err
	}
	var contract ContractModel
	_ = json.Unmarshal([]byte(res), &contract)

	return contract, txID, nil
}
func (this *ContractModel) ContractList(page int) ([]ContractModel, int, int) {
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

	return data, totalCount, totalPages
}
func (this *ContractModel) ContractAdd(postValues map[string][]string, filePath string, acc string) error {

	var contract ContractModel

	contract.ContractId = postValues["contractId"][0]
	contract.ContractName = postValues["contractName"][0]
	contract.ContractVersion = postValues["contractVersion"][0]
	contract.ContractCompanyName = postValues["contractCompany"][0]
	contract.ContractDetails = postValues["contractDetails"][0]
	contract.EnergyType = postValues["energyType"][0]
	contract.EnergyPrice = postValues["energyPrice"][0]
	contract.ContractLastTime = postValues["contractLast"][0]

	contractJson, err := json.Marshal(&contract)
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
