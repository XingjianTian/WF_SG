package main

import (
	ds "WF_SG/DataStructure"
	sig "WF_SG/Utils"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
)

var userid string = "Admin@HUST.builder.com"

func mockInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func addTable(t *testing.T, stub *shim.MockStub, args []string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("addTable"), []byte(args[0]), []byte(userid)})

	if res.Status != shim.OK {
		fmt.Println("AddTable failed:", args[0], string(res.Message))
		t.FailNow()
	}
}
func searchTableById(t *testing.T, stub *shim.MockStub, tableId string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("searchTableById"), []byte(tableId)})
	if res.Status != shim.OK {
		fmt.Println("SearchTableById :", tableId, ", failed :", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("SearchTableById :", tableId, " failed to get value")
		t.FailNow()
	}
}
func queryAllTables(t *testing.T, stub *shim.MockStub) {
	res := stub.MockInvoke("1", [][]byte{[]byte("queryAllTables")})
	if res.Status != shim.OK {
		fmt.Println("query all tables: failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("query all tables failed to get value")
		t.FailNow()
	}
}
func queryAllTablesWithoutExclude(t *testing.T, stub *shim.MockStub) {
	res := stub.MockInvoke("1", [][]byte{[]byte("queryAllTablesWithoutExclude")})
	if res.Status != shim.OK {
		fmt.Println("query all tables without exclude: failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("query all tables failed to get value")
		t.FailNow()
	}
}

// test addTable and search
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
	*/

	//test use private key to sign
	signature1, _ := sig.Sign(tableAsJsonBytes1, userid)
	table1.Sig = append(table1.Sig, signature1)
	tableAsJsonBytes1, _ = json.Marshal(table1)

	//test use private key to sign
	/*
		signature2,_ :=sig.Sign(tableAsJsonBytes2,userid)
		table2.Sig=append(table1.Sig, signature2)
		tableAsJsonBytes2,err = json.Marshal(table2)
	*/

	addTable(t, stub, []string{string(tableAsJsonBytes1)})
	searchTableById(t, stub, "1@supervisor.com")
	//addTable(t, stub, []string{string(tableAsJsonBytes)})
	queryAllTables(t, stub)
	queryAllTablesWithoutExclude(t, stub)
	return
}
