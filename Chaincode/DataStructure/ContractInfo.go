package DataStructure

type ContractModel struct {
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

	ContractSignTime    string `gorm:"type:varchar(50);not null;" json:"contract_sign_time"`
	ContractUserAccount string `gorm:"type:varchar(50);not null;" json:"contract_user_account"`
	ContractUserSig     string `gorm:"type:varchar(50);not null;" json:"contract_user_sig"`
}

func (this *ContractModel) ContractKey() string {
	return this.ContractId + "-" + this.ContractVersion + "-" + this.ContractUserAccount
}

type BidModel struct {
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
