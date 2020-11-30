package main

import (
	ds "WF_SG/Chaincode/DataStructure"
	sig "WF_SG/Chaincode/Utils"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
)

var userAccount = "Admin@HUST.builder.com"

func mockInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func addIed(t *testing.T, stub *shim.MockStub, args []string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("addIed"), []byte(args[0]), []byte(args[1])})

	if res.Status != shim.OK {
		fmt.Println("AddIed failed:", args[0], string(res.Message))
		t.FailNow()
	}
}
func queryIedById(t *testing.T, stub *shim.MockStub, id string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("queryIedById"), []byte(id)})
	if res.Status != shim.OK {
		fmt.Println("queryIedById :", id, ", failed :", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("queryIedById :", id, " failed to get value")
		t.FailNow()
	}
}
func queryAllIeds(t *testing.T, stub *shim.MockStub) {
	res := stub.MockInvoke("1", [][]byte{[]byte("queryAllIeds")})
	if res.Status != shim.OK {
		fmt.Println("query all Ieds: failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("query all Ieds failed to get value")
		t.FailNow()
	}
}

func addContract(t *testing.T, stub *shim.MockStub, args []string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("addContract"), []byte(args[0]), []byte(args[1])})

	if res.Status != shim.OK {
		fmt.Println("AddIed failed:", args[0], string(res.Message))
		t.FailNow()
	}
}
func queryContractById(t *testing.T, stub *shim.MockStub, id string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("queryContractByKey"), []byte(id)})
	if res.Status != shim.OK {
		fmt.Println("queryContractByKey :", id, ", failed :", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("queryContractByKey :", id, " failed to get value")
		t.FailNow()
	}
}
func queryAllContracts(t *testing.T, stub *shim.MockStub) {
	res := stub.MockInvoke("1", [][]byte{[]byte("queryAllContracts")})
	if res.Status != shim.OK {
		fmt.Println("query all Contracts: failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("query all Contracts failed to get value")
		t.FailNow()
	}
}

func Test(t *testing.T) {
	smartContract := new(SmartContract)
	stub := shim.NewMockStub("SmartContract", smartContract)
	mockInit(t, stub, nil)

	var ied1 = ds.IedModel{
		DeviceId:           "1",
		DeviceName:         "Intelligent Device Generation-2",
		DeviceProducer:     "Smart Bridge Company",
		DeviceWorkingDays:  0,
		DeviceBelongIem:    "Lakeview Community",
		DeviceUserAccount:  userAccount,
		DeviceDownInfos:    nil,
		DeviceWorkingInfos: nil,
	}
	iedAsJsonBytes1, _ := json.Marshal(ied1)
	addIed(t, stub, []string{string(iedAsJsonBytes1), ied1.DeviceId})
	queryIedById(t, stub, ied1.DeviceId)
	//addTable(t, stub, []string{string(tableAsJsonBytes)})
	queryAllIeds(t, stub)

	///contractt

	var contract1 = ds.ContractModel{
		ContractId:                  "1231",
		ContractVersion:             "1.0",
		ContractName:                "Smart Water",
		ContractCompanyName:         "SanFrancisco Power",
		ContractCompanyOwnerAccount: userAccount,
		ContractCompanyOwnerSig:     "",
		ContractDetails:             "saeoijufhnoquawehjnfoiq",
		EnergyType:                  "water",
		EnergyPrice:                 "1.39",
		ContractLastTime:            "2020.01.12",
		ContractSignTime:            "",
		ContractUserAccount:         userAccount,
		ContractUserSig:             "",
	}

	contractAsJsonBytes1, _ := json.Marshal(contract1)
	var bid ds.BidModel
	_ = json.Unmarshal(contractAsJsonBytes1, &bid)
	bidAsJsonBytes, _ := json.Marshal(bid)

	//test use private key to sign
	signature1, _ := sig.Sign(bidAsJsonBytes, bid.ContractCompanyOwnerAccount)
	contract1.ContractCompanyOwnerSig = signature1
	contractAsJsonBytes1, _ = json.Marshal(contract1)

	signature2, _ := sig.Sign(contractAsJsonBytes1, contract1.ContractUserAccount)
	contract1.ContractUserSig = signature2
	contractAsJsonBytes1, _ = json.Marshal(contract1)

	addContract(t, stub, []string{string(contractAsJsonBytes1), contract1.ContractId})
	queryContractById(t, stub, contract1.ContractId)
	queryAllContracts(t, stub)

	return
}

/*
func TestAddTable(t *testing.T) {
	smartContract := new(SmartContract)
	stub := shim.NewMockStub("SmartContract", smartContract)
	mockInit(t, stub, nil)

	var table1 = ds.Table{
		TId:   "1@supervisor.com",
		TName: "testTable1",
		IotData: []ds.Property{
			{
				PId:      "1",
				PName:    "吊绳拉力",
				PDefault: "KN",
				PRange:   "0,100",
				PValue:   "90.62KN",
			},
		},

		Cs: []ds.Container{
			{
				CName: "Container1",
				//UpperId:"none",
				Ps: []ds.Property{
					{
						PId:      "1",
						PName:    "propertyone",
						PDefault: "mm",
						PRange:   "-10,10",
						PValue:   "3mm 11mm 90mm",
					},
					{
						PId:      "2",
						PName:    "propertytwo",
						PDefault: "mm",
						PRange:   "-1,1",
						PValue:   "0",
					},
				},
			},
		},
	}
	table1.LastSigner = userid
	tableAsJsonBytes1, _ := json.Marshal(table1)
	/*
		var table2 = ds.Table{
			TId:"2@builder.com",
			TName:"testTable2",
			PriorTIds:[]string{"1"},
			Cs: []ds.Container{
				{
					CName:"Container1",
					//UpperId:"none",
					Ps: []ds.Property{
						{
							PId:"1",
							PName:"propertyone",
							PDefault:"none",
							PValue:"defaultvalue",
						},
						{
							PId:"2",
							PName:"propertytwo",
							PDefault:"none",
							PValue:"defaultvalue",
						},
					},
				},

				{
					CName:"Container2",
					//UpperId:"none",
					Ps: []ds.Property{
						{
							PId:"3",
							PName:"propertyone",
							PDefault:"none",
							PValue:"defaultvalue",
						},
					},
				},
			},
		}
		table2.LastSigner = userid
		tableAsJsonBytes1,err := json.Marshal(table1)

		if err!=nil{
			t.FailNow()
		}
		tableAsJsonBytes2,_ := json.Marshal(table2)

	//test use private key to sign
	signature1, _ := sig.Sign(tableAsJsonBytes1, userid)
	table1.Sig = append(table1.Sig, signature1)
	tableAsJsonBytes1, _ = json.Marshal(table1)

	//test use private key to sign
		signature2,_ :=sig.Sign(tableAsJsonBytes2,userid)
		table2.Sig=append(table1.Sig, signature2)
		tableAsJsonBytes2,err = json.Marshal(table2)

	addTable(t, stub, []string{string(tableAsJsonBytes1)})
	searchTableById(t, stub, "1@supervisor.com")
	//addTable(t, stub, []string{string(tableAsJsonBytes)})
	queryAllTables(t, stub)
	queryAllTablesWithoutExclude(t, stub)
	return
}
*/
