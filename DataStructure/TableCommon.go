package DataStructure

type Common struct {
	OrgEngineeringName string `json:"Organization_engineeringName"`
	DepEngineeringName string `json:"Department_engineeringName"`
	SubEngineeringName string `json:"Sub_engineeringName"`

	ConstructionOrgName string `json:"Construction_orgName"`
	ConstructionLeader  string `'json:"Construction_leader"`
	TestCapacity        string `'json:"Test_capacity"`

	ContractorOrgName string `json:"Contractor_orgName"`
	ContractorLeader  string `'json:"Contractor_leader"`
	TestPart          string `'json:"Test_Part"`

	ConstructionAccordance string `json:"Construction_accordance"`
	InspectionAccordance   string `json:"Inspection_accordance"`
}
