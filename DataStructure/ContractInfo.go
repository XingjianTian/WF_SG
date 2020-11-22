package DataStructure

type ContractInfo struct {
	ContractId      string
	ContractVersion string
	//fabric key:id+"-"+version
	ContractName        string
	ContractCompanyName string
	ContractCompanySig  string
	ContractDetails     string
	EnergyType          string
	EnergyPrice         string
	ContractLastTime    string
}
