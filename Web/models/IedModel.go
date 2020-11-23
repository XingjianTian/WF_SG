package models

import (
	"WF_SG/Web/common"
	"errors"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/jinzhu/gorm"
	config "github.com/spf13/viper"
	"math"
)

type IedModel struct {
	gorm.Model

	DeviceId           uint                `gorm:"primary_key;type:uint;not null;" json:"device_id"`
	DeviceName         string              `gorm:"type:varchar(50);not null;" json:"device_name"`
	DeviceProducer     string              `gorm:"type:varchar(50);not null;" json:"device_producer"`
	DeviceWorkingDays  int                 `gorm:"type:int;not null;" json:"device_working_days"`
	DeviceBelongIem    string              `gorm:"type:varchar(50);not null;" json:"device_belong_iem"`
	DeviceUserAccount  string              `gorm:"type:varchar(50);not null;" json:"device_user_account"`
	DeviceDownInfos    []DeviceDownInfo    `gorm:"foreignkey:DownId;association_foreignkey:DeviceId;" json:"device_down_infos"`
	DeviceWorkingInfos []DeviceWorkingInfo `gorm:"foreignkey:WorkingId;association_foreignkey:DeviceId;" json:"device_working_infos"`
}

func (this *IedModel) TableName() string {
	return "ied"
}

type DeviceDownInfo struct {
	gorm.Model
	DownId       uint                `gorm:"primary_key;auto_increment"`
	DownTimeFrom timestamp.Timestamp `gorm:"type:timestamp;" json:"down_time_from"`
	DownTimeTo   timestamp.Timestamp `gorm:"type:timestamp;" json:"down_time_to"`
	DownReason   string              `gorm:"type:varchar(50);not null;" json:"down_reason"`
	RepairPerson string              `gorm:"type:varchar(50);not null;" json:"repair_person"`
	RepairCost   string              `gorm:"type:varchar(50);not null;" json:"repair_cost"`
}
type DeviceWorkingInfo struct {
	gorm.Model
	WorkingId      uint                `gorm:"primary_key;auto_increment"`
	WorkingDay     timestamp.Timestamp `gorm:"type:timestamp;" json:"working_day"`
	EnergyConsumed string              `gorm:"type:varchar(50);not null;" json:"energy_consumed"`
	EnergyProduced string              `gorm:"type:varchar(50);not null;" json:"energy_produced"`
	EnergyCost     string              `gorm:"type:varchar(50);not null;" json:"energy_cost"`
	ContractName   string              `gorm:"type:varchar(50);not null;" json:"contract_name"`
}

func (this *IedModel) IedList(page int) ([]IedModel, int, int, error) {
	var IEDs []IedModel
	var totalCount int
	limit := config.GetInt("pagination.PageSize")
	offset := (page - 1) * limit
	db := common.DB
	db.Find(&IEDs).Count(&totalCount)
	err := db.Offset(offset).Limit(limit).Order("id desc").Find(&IEDs).Error
	if err != nil {
		return IEDs, 0, 0, err
	}
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))
	return IEDs, totalCount, totalPages, nil
}

func (this *IedModel) IedInfo(id string) (IedModel, fab.TransactionID, error) {
	var ied IedModel
	db := common.DB
	if db.Where("device_id = ?", id).First(&ied).RecordNotFound() {
		return IedModel{}, fab.TransactionID(""), errors.New("ied not found")
	}
	return ied, fab.TransactionID(""), nil
}

func (this *IedModel) IedUpdate(postValues map[string][]string) error {
	db := common.DB

	id := postValues["deviceId"][0]

	ied, _, err := this.IedInfo(id)
	if err != nil {
		return err
	}

	if err := db.Save(&ied).Error; err != nil {
		return errors.New("Failed to update usr")
	}

	return nil
}
func (this *IedModel) IedAdd(postValues map[string][]string, acc string) error {
	var company IedModel

	db := common.DB

	if !db.Where("company_name = ? ", postValues["companyName"][0]).First(&company).RecordNotFound() {
		return errors.New("company already exits")
	}

	if err := db.Create(&company).Error; err != nil {
		return errors.New("Failed to register company")
	}

	return nil
}

func (this *IedModel) CompanyDel(companyNamae string) error {
	var company IedModel
	db := common.DB
	if err := db.Where("company_name = ?", companyNamae).Delete(&company).Error; err != nil {
		return errors.New("Failed to delete company")
	}
	return nil
}
