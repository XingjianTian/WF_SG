package DataStructure

type ContractInfo struct {
	ContractName    string  `json:"contract_name"`
	ContractCompany string  `json:"contract_company"`
	ContractDetails string  `json:"contract_details"`
	EnergyType      string  `json:"energy_type"`
	EnergyPricePerW float64 `json:"energy_pricePerW"`
}
