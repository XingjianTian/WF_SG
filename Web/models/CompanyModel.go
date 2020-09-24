package models

import (
	"WF_SG/Web/common"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	config "github.com/spf13/viper"
	"log"
	"math"
	"time"
)

type CompanyModel struct {
	gorm.Model
	//ID       uint    `gorm:"primary_key;auto_increment"`
	CompanyName         string `gorm:"type:varchar(50);not null;"`
	CompanyOwnerAccount string `gorm:"type:varchar(50);not null;"`
	CompanyBuildYear    string `gorm:"type:char(32);not null;"`
	EmailAddress        string `gorm:"type:char(50);not null;"`
	PhoneNumber         string `gorm:"type:char(50);not null;"`
	Descript            string `gorm:"type:varchar(255);DEFAULT '';"`
	Headico             string `gorm:"type:varchar(200);DEFAULT '';"`
	CompanyInvest       string `gorm:"type:varchar(50);DEFAULT '';"`
	CompanySize         string `gorm:"type:varchar(50);DEFAULT '';"`
	CompanyLocation     string `gorm:"type:varchar(50);DEFAULT '';"`
	CompanyWebsite      string `gorm:"type:varchar(50);DEFAULT '';"`
}

// 设置User的表名为`profiles`
func (this *CompanyModel) TableName() string {
	return "company"
}
func (this *CompanyModel) CompanyList(page int) ([]CompanyModel, int, int) {
	var data []CompanyModel
	var totalCount int
	limit := config.GetInt("pagination.PageSize")
	offset := (page - 1) * limit
	db := common.DB
	db.Find(&data).Count(&totalCount)
	err := db.Offset(offset).Limit(limit).Order("id desc").Find(&data).Error
	if err != nil {
		log.Fatalln(err)
	}
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))
	return data, totalCount, totalPages
}

func (this *CompanyModel) CompanyInfo(companyName string) (CompanyModel, error) {
	var company CompanyModel
	db := common.DB
	if db.Where("company_name = ?", companyName).First(&company).RecordNotFound() {
		return CompanyModel{}, errors.New("Company not found")
	}
	return company, nil
}

func (this *CompanyModel) CompanyInfoByOwnerAcc(acc string, page int) ([]CompanyModel, int, int) {

	var data []CompanyModel
	var totalCount int
	limit := config.GetInt("pagination.PageSize")
	//offset := (page - 1) * limit
	db := common.DB
	db.Where("company_owner_account = ?", acc).Find(&data).Count(&totalCount)
	/*err := db.Offset(offset).Limit(limit).Order("id desc").Where("company_owner_account = ?", "acc").Find(&data).Error
	if err != nil {
		log.Fatalln(err)
	}
	*/

	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))
	return data, totalCount, totalPages
}

func (this *CompanyModel) CompanyUpdate(postValues map[string][]string, companyName string, filePath string) error {
	db := common.DB
	company, err := this.CompanyInfo(companyName)
	if err != nil {
		return err
	}
	company.Descript = postValues["descript"][0]
	company.EmailAddress = postValues["emailAddress"][0]
	company.PhoneNumber = postValues["phoneNumber"][0]
	company.CompanyInvest = postValues["companyInvest"][0]
	company.CompanySize = postValues["companySize"][0]
	company.CompanyLocation = postValues["companyLocation"][0]
	company.CompanyWebsite = postValues["companyWebsite"][0]

	if filePath != "" {
		company.Headico = filePath
	}
	if err := db.Save(&company).Error; err != nil {
		return errors.New("Failed to update usr")
	}

	return nil
}
func (this *CompanyModel) CompanyAdd(postValues map[string][]string, filePath string, acc string) error {
	var company CompanyModel

	db := common.DB

	if !db.Where("company_name = ? ", postValues["companyName"][0]).First(&company).RecordNotFound() {
		return errors.New("company already exits")
	}

	company.CompanyName = postValues["companyName"][0]
	company.CompanyOwnerAccount = acc
	company.CompanyBuildYear = string(time.Now().Year())
	company.Descript = postValues["descript"][0]
	company.EmailAddress = postValues["emailAddress"][0]
	company.PhoneNumber = postValues["phoneNumber"][0]
	company.CompanyInvest = postValues["companyInvest"][0]
	company.CompanySize = postValues["companySize"][0]
	company.CompanyLocation = postValues["companyLocation"][0]
	company.CompanyWebsite = postValues["CompanyWebsite"][0]

	if filePath != "" {
		company.Headico = filePath
	}
	if err := db.Create(&company).Error; err != nil {
		return errors.New("Failed to register company")
	}

	return nil
}

func (this *CompanyModel) CompanyDel(companyNamae string) error {
	var company CompanyModel
	db := common.DB
	if err := db.Where("company_name = ?", companyNamae).Delete(&company).Error; err != nil {
		return errors.New("Failed to delete company")
	}
	return nil
}
