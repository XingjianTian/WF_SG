package DataStructure

import (
	"strings"
	"time"
)

func AiteBefore(s string) (r string) {
	aiteIndex := strings.Index(s, "@")
	return s[:aiteIndex]
}

func AiteAfter(s string) (r string) {
	aiteIndex := strings.Index(s, "@")
	return s[aiteIndex+1:]
}

type TableForWebinCC struct {
	RealID             string //xxx
	TID                string //xxx@yyy.com
	TName              string
	OrgEngineeringName string
	DepEngineeringName string
	SubEngineeringName string

	TestPart string

	CreatedAt      time.Time //.Format("2006-01-02 15:04")
	CreatedTimeWeb string
	State          string
	Operator       string
	OrgName        string
}

type Table struct {
	TId   string `json:"table_id"`
	TName string `json:"table_name"`

	TCommon Common      `json:"table_common"`
	Cs      []Container `json:"table_containers"`

	//RelatedProjectID []string `json:"table_relatedProject_id"`
	//SignedByObj[] string `json:"table_signedByObj"`
	//TriggerEvent string `json:"table_triggerEvent"`
	//cs[]Container `json:table_containers`

	Sig        []string  `json:"table_signature"`
	PriorTIds  []string  `json:"table_priorTableIds"`
	LastSigner string    `json:"table_lastSigner"`
	TimeStamp  time.Time `json:"table_timeStamp"`

	IotData []Property `json:"Iot_Data"`
}

func (t *Table) CheckAllPropertyValue() {
	//var bAddable = true
	for i, _ := range t.Cs {
		for j, _ := range t.Cs[i].Ps {
			//bAddable=bAddable&&p.CheckRange()
			t.Cs[i].Ps[j].CheckRange()
		}
	}

	for i, _ := range t.IotData {
		t.IotData[i].CheckRange()
	}

	//return bAddable
}
