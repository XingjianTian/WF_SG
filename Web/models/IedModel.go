package models

import (
	"WF_SG/Services"
	"WF_SG/Web/common"
	"encoding/json"
	"errors"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/jinzhu/gorm"
	config "github.com/spf13/viper"
	"math"
)

type IedModel struct {
	gorm.Model

	DeviceId           string              `gorm:"primary_key;type:varchar(50);not null;" json:"device_id"`
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
	var ieds []IedModel

	respJson, err := Services.HLservice.QueryAllIedsService()
	//transform
	err = json.Unmarshal([]byte(respJson), &ieds)

	var totalCount int
	limit := config.GetInt("pagination.PageSize")
	totalCount = len(ieds)
	if err != nil {
		return ieds, 0, 0, err
	}
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))
	return ieds, totalCount, totalPages, nil
}

func (this *IedModel) IedInfo(id string) (IedModel, fab.TransactionID, error) {

	respJson, txID, err := Services.HLservice.QueryIedByIdService(id)
	if err != nil {
		return IedModel{}, "", err
	}

	var ied IedModel
	err = json.Unmarshal([]byte(respJson), &ied)

	return ied, txID, nil
}

func (this *IedModel) IedUpdate(postValues map[string][]string) error {
	db := common.DB

	var ied IedModel
	ied.DeviceId = postValues["deviceId"][0]
	ied.DeviceName = postValues["deviceName"][0]
	ied.DeviceProducer = postValues["deviceProducer"][0]
	ied.DeviceBelongIem = postValues["deviceBelongIem"][0]

	if err := db.Save(&ied).Error; err != nil {
		return errors.New("Failed to update IED")
	}

	return nil
}
func (this *IedModel) IedAdd(postValues map[string][]string, userAcc string) error {

	var ied IedModel
	ied.DeviceId = postValues["deviceId"][0]
	ied.DeviceName = postValues["deviceName"][0]
	ied.DeviceProducer = postValues["deviceProducer"][0]
	ied.DeviceBelongIem = postValues["deviceBelongIem"][0]

	ied.DeviceUserAccount = userAcc

	iedJson, _ := json.Marshal(ied)

	res, err := Services.HLservice.AddIedService(iedJson)
	if err != nil {
		return errors.New("Failed to add IED")
	}
	println(res)
	return nil
}
