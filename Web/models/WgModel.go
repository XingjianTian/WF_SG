package models

import (
	"WF_SG/Web/common"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/skip2/go-qrcode"
	config "github.com/spf13/viper"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"math"
)

type AllocatedIpModel struct {
	gorm.Model
	WgAccount string `gorm:"type:varchar(50);not null;" json:"account"`
	IpAddress string `gorm:"type:varchar(50);not null;" json:"ip_address"`
}
type AllowedIpModel struct {
	gorm.Model
	WgAccount string `gorm:"type:varchar(50);not null;" json:"account"`
	IpAddress string `gorm:"type:varchar(50);not null;" json:"ip_address"`
}
type ListenIpModel struct {
	gorm.Model
	WgAccount string `gorm:"type:varchar(50);not null;" json:"account"`
	IpAddress string `gorm:"type:varchar(50);not null;" json:"ip_address"`
}
type WgModel struct {
	gorm.Model
	//client
	//ID       uint    `gorm:"primary_key;auto_increment"`
	CreatedAt    timestamp.Timestamp `gorm:"type:timestamp;" json:"created_at"`
	UpdatedAt    timestamp.Timestamp `gorm:"type:timestamp;" json:"updated_at"`
	DeletedAt    timestamp.Timestamp `gorm:"type:timestamp;" json:"deleted_at"`
	Account      string              `gorm:"type:varchar(50);not null;" json:"account"`
	Enabled      bool                `gorm:"type:tinyint(1);not null;" json:"enabled"`
	PrivateKey   string              `gorm:"type:varchar(200);not null;" json:"private_key"`
	PublicKey    string              `gorm:"type:varchar(200);not null;" json:"public_key"`
	PresharedKey string              `gorm:"type:varchar(200);not null;" json:"preshared_key"`
	AllocatedIps []AllocatedIpModel  `gorm:"foreignkey:WgAccount;association_foreignkey:Account;" json:"allocated_ips"` //- split
	AllowedIps   []AllowedIpModel    `gorm:"foreignkey:WgAccount;association_foreignkey:Account;" json:"allowed_ips"`   //- split
	Qrcode       string              `gorm:"type:varchar(1500);not null;" json:"qrcode"`

	//server
	ListenAddresses []ListenIpModel `gorm:"foreignkey:ID;association_foreignkey:Account;" json:"listen_addresses"`

	ListenPort string `gorm:"type:varchar(50);" json:"listen_port"`
	PostUp     string `gorm:"type:varchar(50);" json:"post_up"`
	PostDown   string `gorm:"type:varchar(50);" json:"post_down"`
}

func (this *WgModel) TableName() string {
	return "wg"
}

func (this *WgModel) WgKeyGenerate() error {
	// gen Wireguard key pair
	key, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return errors.New("Failed to generate wireguard key pair: " + err.Error())
	}
	presharedKey, err := wgtypes.GenerateKey()
	if err != nil {
		return errors.New("Failed to generate wireguard preshared key: " + err.Error())
	}
	this.PrivateKey = key.String()
	this.PublicKey = key.PublicKey().String()
	this.PresharedKey = presharedKey.String()

	return nil
}

func (this *WgModel) WgInfo(acc string) (WgModel, error) {
	var wgModel WgModel
	db := common.DB
	if db.Where("account = ?", acc).First(&wgModel).RecordNotFound() {
		return WgModel{}, errors.New("acc not found")
	}
	db.Preload("AllocatedIps").
		Preload("AllowedIps").
		Preload("ListenAddresses").Where("account = ?", acc).First(&wgModel)
	return wgModel, nil
}

func (this *WgModel) WgList(page int) ([]WgModel, int, int) {
	var data []WgModel
	var totalCount int
	limit := config.GetInt("pagination.PageSize")
	db := common.DB

	if !db.HasTable(WgModel{}) {
		return nil, 0, 0
	}

	db.Preload("AllocatedIps").
		Preload("AllowedIps").
		Preload("ListenAddresses").
		Find(&data).Count(&totalCount)
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))
	return data, totalCount, totalPages
}
func (this *WgModel) WgAdd(wgJson []byte) error {
	// read server information
	// validate the input Allocation IPs
	// validate the input AllowedIPs
	// gen ID
	// gen Wireguard key pair
	// write client to the database
	var wg WgModel
	err := json.Unmarshal(wgJson, &wg)
	if err != nil {
		return err
	}
	db := common.DB
	if !db.HasTable(wg) {
		db.CreateTable(wg)
	}
	if !db.HasTable(AllowedIpModel{}) {
		db.CreateTable(AllowedIpModel{})
	}
	if !db.HasTable(AllocatedIpModel{}) {
		db.CreateTable(AllocatedIpModel{})
	}
	if !db.HasTable(ListenIpModel{}) {
		db.CreateTable(ListenIpModel{})
	}

	if !db.Where("account = ? ", wg.Account).First(&wg).RecordNotFound() {
		return errors.New("wireguard instance already exists")
	}

	//key
	err = wg.WgKeyGenerate()
	if err != nil {
		return errors.New("Failed to generate wireguard key pair")
	}

	//qrcode

	png, _ := qrcode.Encode(string(wgJson), qrcode.Low, 256)

	wg.Qrcode = "data:image/png;base64," + base64.StdEncoding.EncodeToString(png)
	sz := len(wg.Qrcode)
	fmt.Print(sz)

	if err := db.Create(&wg).Error; err != nil {
		return errors.New("failed to add wireguard instance")
	}

	return nil
}

func (this *WgModel) WgUpdate(wgJson []byte, acc string) error {
	db := common.DB
	wg, err := this.WgInfo(acc)
	if err != nil {
		return err
	}
	err = json.Unmarshal(wgJson, &wg)
	if err != nil {
		return err
	}

	if err := db.Save(&wg).Error; err != nil {
		return errors.New("fail to update wg")
	}

	return nil
}

func (this *WgModel) WgUpdateStatus(acc string, status bool) error {
	db := common.DB
	wg, err := this.WgInfo(acc)
	if err != nil {
		return err
	}
	wg.Enabled = status
	if err := db.Save(&wg).Error; err != nil {
		return errors.New("fail to update wg")
	}

	return nil
}

func (this *WgModel) WgDel(acc string) error {
	db := common.DB

	db.Where("wg_account = ?", acc).Delete(&AllocatedIpModel{})
	db.Where("wg_account = ?", acc).Delete(&AllowedIpModel{})
	db.Where("wg_account = ?", acc).Delete(&ListenIpModel{})
	db.Where("account = ?", acc).Delete(&WgModel{})

	return nil
}

func (this *WgModel) IpList(ipType string, count int) ([]byte, error) {
	db := common.DB
	if ipType == "allocated" {
		var list []AllocatedIpModel
		db.Find(&list).Limit(count)
		return json.Marshal(list)
	} else if ipType == "allowed" {
		var list []AllowedIpModel
		db.Find(&list).Limit(count)
		return json.Marshal(list)
	} else if ipType == "listen" {
		var list []ListenIpModel
		db.Find(&list).Limit(count)
		return json.Marshal(list)
	} else {
		return nil, errors.New("no such ip table")
	}

}

// GlobalSetting model
type GlobalSetting struct {
	EndpointAddress     string   `gorm:"type:varchar(200);not null;" json:"endpoint_address"`
	DNSServers          []string `gorm:"type:varchar(200);not null;" json:"dns_servers"`
	MTU                 string   `gorm:"type:varchar(50);not null;" json:"mtu"`
	PersistentKeepalive string   `gorm:"type:varchar(50);not null;" json:"persistent_keepalive"`
	ConfigFilePath      string   `gorm:"type:varchar(200);not null;" json:"config_file_path"`
}
