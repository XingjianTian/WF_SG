package DataStructure

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type IedModel struct {
	DeviceId           string              `gorm:"primary_key;type:varchar(50);not null;" json:"device_id"`
	DeviceName         string              `gorm:"type:varchar(50);not null;" json:"device_name"`
	DeviceProducer     string              `gorm:"type:varchar(50);not null;" json:"device_producer"`
	DeviceWorkingDays  int                 `gorm:"type:int;not null;" json:"device_working_days"`
	DeviceBelongIem    string              `gorm:"type:varchar(50);not null;" json:"device_belong_iem"`
	DeviceUserAccount  string              `gorm:"type:varchar(50);not null;" json:"device_user_account"`
	DeviceDownInfos    []DeviceDownInfo    `gorm:"foreignkey:DownId;association_foreignkey:DeviceId;" json:"device_down_infos"`
	DeviceWorkingInfos []DeviceWorkingInfo `gorm:"foreignkey:WorkingId;association_foreignkey:DeviceId;" json:"device_working_infos"`
}

type DeviceDownInfo struct {
	DownId       uint                `gorm:"primary_key;auto_increment"`
	DownTimeFrom timestamp.Timestamp `gorm:"type:timestamp;" json:"down_time_from"`
	DownTimeTo   timestamp.Timestamp `gorm:"type:timestamp;" json:"down_time_to"`
	DownReason   string              `gorm:"type:varchar(50);not null;" json:"down_reason"`
	RepairPerson string              `gorm:"type:varchar(50);not null;" json:"repair_person"`
	RepairCost   string              `gorm:"type:varchar(50);not null;" json:"repair_cost"`
}
type DeviceWorkingInfo struct {
	WorkingId      uint                `gorm:"primary_key;auto_increment"`
	WorkingDay     timestamp.Timestamp `gorm:"type:timestamp;" json:"working_day"`
	EnergyConsumed string              `gorm:"type:varchar(50);not null;" json:"energy_consumed"`
	EnergyProduced string              `gorm:"type:varchar(50);not null;" json:"energy_produced"`
	EnergyCost     string              `gorm:"type:varchar(50);not null;" json:"energy_cost"`
	ContractName   string              `gorm:"type:varchar(50);not null;" json:"contract_name"`
}
