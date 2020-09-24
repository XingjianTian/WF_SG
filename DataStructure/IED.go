package DataStructure

type IED struct {
	DeviceName       string         `json:"device_name"`
	DeviceProducer   string         `json:"device_producer"`
	DeviceWokingInfo IEDWorkingInfo `json:"device_workingInfo"`
	DeviceUser       UserInfo       `json:"device_user"`
	DeviceDownInfos  []IEDDownInfo  `json:"device_downInfos"`
}

type IEDDownInfo struct {
	DownTimeLast          string  `json:"down_timeLas"`
	DownReason            string  `json:"down_reason"`
	DownResponsiblePerson string  `json:"down_responsiblePerson"`
	RepairPerson          string  `json:"repair_person"`
	DownTimeFrom          string  `json:"down_timeFrom"`
	DownTimeTo            string  `json:"down_timeTo"`
	RepairCost            float64 `json:"repair_cost"`
}

type IEDWorkingInfo struct {
	WoringTimeFrom   string       `json:"working_timeFrom"`
	WoringTimeLast   string       `json:"working_timeLast"`
	BelongToIEM      string       `json:"belong_toIEM"`
	AveCostPerDay    float64      `json:"Ave_costPerDay"`
	AveCostPerW      float64      `json:"Ave_costPerW"`
	EnergyInfoByDays []EnergyInfo `json:"energy_infoByDays"`
}
type EnergyInfo struct {
	EnergyTimeFrom string       `json:"energy_timeFrom"`
	EnergyTimeTo   string       `json:"energy_timeTo"`
	EnergyConsumed float64      `json:"energy_consumed"`
	EnergyProduced float64      `json:"energy_produced"`
	EnergyCost     float64      `json:"energy_cost"`
	CarbonProduced float64      `json:"carbon_cost"`
	EnergyContract ContractInfo `json:"energy_contract"`
}
