package models

import (
	"WF_SG/Web/common"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	config "github.com/spf13/viper"
	"log"
	"math"
	"sync"
)

var mu sync.Mutex

type UserModel struct {
	gorm.Model
	//ID       uint    `gorm:"primary_key;auto_increment"`
	Account      string `gorm:"type:varchar(20);not null;index:username"`
	Password     string `gorm:"type:char(32);not null;"`
	EmailAddress string `gorm:"type:char(32);not null;"`
	PhoneNumber  string `gorm:"type:char(32);not null;"`
	Descript     string `gorm:"type:varchar(255);DEFAULT '';"`
	Headico      string `gorm:"type:varchar(200);DEFAULT '';"`
	Type         string `gorm:"type:varchar(20);not null;"`
}

// 设置User的表名为`profiles`
func (this *UserModel) TableName() string {
	return "user"
}

func (this *UserModel) UserList(page int) ([]UserModel, int, int) {
	var data []UserModel
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

func (this *UserModel) UserLogin(acc string, passwd string) (UserModel, error) {
	if acc == "" || passwd == "" {
		return UserModel{}, errors.New("acc or passwd null")
	}
	db := common.DB
	var user UserModel
	has := md5.Sum([]byte(passwd))
	md5_passwd := fmt.Sprintf("%x", has)
	if db.Where(&UserModel{Account: acc, Password: md5_passwd}).First(&user).RecordNotFound() {
		return UserModel{}, errors.New("wrong acc or passwd")
	}
	return user, nil
}

func (this *UserModel) UserInfo(acc string) (UserModel, error) {
	var userModel UserModel
	db := common.DB
	if db.Where("account = ?", acc).First(&userModel).RecordNotFound() {
		return UserModel{}, errors.New("acc not found")
	}
	return userModel, nil
}

func (this *UserModel) UserPasswdUpdate(acc string, passwd, repasswd string) error {
	if passwd == "" || repasswd == "" {
		return errors.New("no nil passwd")
	}
	if passwd != repasswd {
		return errors.New("not equal")
	}
	db := common.DB
	user, err := this.UserInfo(acc)

	if err != nil {
		return err
	}

	has := md5.Sum([]byte(passwd))
	md5_password := fmt.Sprintf("%x", has) //将[]byte转成16进制

	if err := db.Model(&user).Update("password", md5_password).Error; err != nil {
		return errors.New("fail to update passwd")
	}

	return nil
}
func (this *UserModel) UserUpdate(postValues map[string][]string, acc string, filePath string) error {
	db := common.DB
	user, err := this.UserInfo(acc)
	if err != nil {
		return err
	}
	user.Descript = postValues["descript"][0]
	user.EmailAddress = postValues["emailAddress"][0]
	user.PhoneNumber = postValues["phoneNumber"][0]
	if filePath != "Error while uploading: <b>http: no such file</b>" {
		user.Headico = filePath
	}
	if err := db.Save(&user).Error; err != nil {
		return errors.New("fail to update usr")
	}

	return nil
}

func (this *UserModel) UserAdd(postValues map[string][]string, filePath string) error {
	var user UserModel

	if postValues["password"][0] == "" || postValues["Repassword"][0] == "" {
		return errors.New("no nil passwd")
	}
	if postValues["password"][0] != postValues["Repassword"][0] {
		return errors.New("not equal")
	}
	delete(postValues, "Repassword")
	has := md5.Sum([]byte(postValues["password"][0]))
	postValues["password"][0] = fmt.Sprintf("%x", has) //将[]byte转成16进制

	db := common.DB

	if !db.Where("account = ? ", postValues["account"][0]).First(&user).RecordNotFound() {
		return errors.New("usr already exits")
	}

	user.Account = postValues["account"][0]
	user.Password = postValues["password"][0]
	user.Descript = postValues["descript"][0]
	user.Type = "user"
	if filePath != "" {
		user.Headico = filePath
	}
	if err := db.Create(&user).Error; err != nil {
		return errors.New("failed to add user")
	}

	return nil
}

func (this *UserModel) UserDel(acc string) error {
	var user UserModel
	db := common.DB
	if err := db.Where("account = ?", acc).Delete(&user).Error; err != nil {
		return errors.New("failed to delete user")
	}
	return nil
}
