package DataStructure

type CompanyInfo struct {
	CompanyName       string  `json:"company_name"`
	CompanyOwner      string  `json:"company_owner"`
	CompanyEmail      string  `json:"company_email"`
	CompanyPhone      string  `json:"company_phone"`
	CompanyTimeFrom   string  `json:"company_timeFrom"`
	CompanyInvestment float64 `json:"company_investment"`
	CompanySize       string  `json:"company_size"`
	CompanyLocation   string  `json:"company_location"`
	CompanyWebsite    string  `json:"company_website"`
}
