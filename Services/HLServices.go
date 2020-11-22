package Services

import (
	ds "WF_SG/DataStructure"
	sig "WF_SG/Utils"
	//"WF_SG/Web/models"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"strings"
	"time"
)

type ServiceSetup struct {
	ChaincodeID string
	Clients     map[string]*channel.Client
	Ledgers     map[string]*ledger.Client
}

var HLservice ServiceSetup

func eventRegister(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("Chaincode registering failed: %s", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("Chaincode event received: %v\n", ccEvent)
	case <-time.After(time.Second * 50):
		return fmt.Errorf("Can't receive corresponding chaincode event according to event id (%s)", eventID)
	}
	return nil
}

//table
func (t *ServiceSetup) AddTableService(tableAsJsonBytes []byte, userid string, userOrgName string) (string, error) {
	eventID := "eventAddTable"
	cli := t.Clients["WH-zhijianju"]
	//cli := t.Clients[userOrgName]
	reg, notifer := eventRegister(cli, t.ChaincodeID, eventID)
	defer cli.UnregisterChaincodeEvent(reg)

	var table ds.Table
	err := json.Unmarshal(tableAsJsonBytes, &table)
	if err != nil {
		return "", err
	}
	aiteIndex := strings.Index(table.TId, "@")
	var tmpID string
	if aiteIndex != -1 {
		tmpID = table.TId[:aiteIndex]
	} else {
		tmpID = table.TId
	}
	table.LastSigner = userid
	//table.LastSigner = userOrgName
	table.TId = tmpID + "@" + userOrgName
	table.TimeStamp = time.Now()

	tableAsJsonBytes, err = json.Marshal(table)
	//test use private key to sign
	signature, _ := sig.Sign(tableAsJsonBytes, userid)
	table.Sig = append(table.Sig, signature)
	tableAsJsonBytes, err = json.Marshal(table)

	if err != nil {
		return "", err
	}
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "addTable", Args: [][]byte{tableAsJsonBytes, []byte(userid)}}
	response, err := cli.Execute(req)
	if err != nil {
		return "", err
	}
	err = eventResult(notifer, eventID)
	if err != nil {
		return "", err
	}
	return string(response.TransactionID), nil
}

func (t *ServiceSetup) SearchTxIDByTableId(tableID string, userOrgName string) (fab.TransactionID, error) {
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "searchTxIDByTableId", Args: [][]byte{[]byte(tableID)}}

	//cli:=t.Clients[userOrgName]
	cli := t.Clients["WH-zhijianju"]
	response, err := cli.Query(req)
	if err != nil {
		return "", err
	}
	return fab.TransactionID(response.Payload), nil

}

func (t *ServiceSetup) QueryAllTablesWithoutExclude(userOrgName string) (string, error) {
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryAllTablesWithoutExclude"}

	//cli:=t.Clients[userOrgName]
	cli := t.Clients["WH-zhijianju"]
	response, err := cli.Query(req)
	if err != nil {
		return "", err
	}
	return string(response.Payload), nil
}

func (t *ServiceSetup) SearchTableByIdService(tableID string, userOrgName string) (string, fab.TransactionID, error) {
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "searchTableById", Args: [][]byte{[]byte(tableID)}}

	//cli:=t.Clients[userOrgName]
	cli := t.Clients["WH-zhijianju"]
	response, err := cli.Query(req)
	if err != nil {
		return "", "", err
	}

	txID, err := t.SearchTxIDByTableId(tableID, userOrgName)

	return string(response.Payload), txID, nil

}

func (t *ServiceSetup) QueryBlockByTx(txID fab.TransactionID) (*common.Block, error) {
	//cli:=t.Clients[userOrgName]
	cli := t.Ledgers["WH-zhijianju"]

	response, err := cli.QueryBlockByTxID(txID)
	if err != nil {
		if strings.Contains(err.Error(), "Entry not found in index") {
			return nil, nil
		}
		return nil, err
	}

	return response, nil

}

func (t *ServiceSetup) QueryAllTables(userOrgName string) (string, error) {
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryAllTables"}

	//cli:=t.Clients[userOrgName]
	cli := t.Clients["WH-zhijianju"]
	response, err := cli.Query(req)
	if err != nil {
		return "", err
	}
	return string(response.Payload), nil

}

//contract
func (t *ServiceSetup) AddContractService(contractJson []byte) (string, error) {
	eventID := "eventAddContract"
	cli := t.Clients["WH-zhijianju"]
	//cli := t.Clients[userOrgName]
	reg, notifer := eventRegister(cli, t.ChaincodeID, eventID)
	defer cli.UnregisterChaincodeEvent(reg)

	var contract ds.ContractInfo
	err := json.Unmarshal(contractJson, &contract)
	if err != nil {
		return "", err
	}
	//use private key to sign
	signature, _ := sig.Sign(contractJson, contract.ContractCompanyName)
	contract.ContractCompanySig = signature
	contractJson, err = json.Marshal(contract)

	if err != nil {
		return "", err
	}
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "addContract", Args: [][]byte{contractJson,
		[]byte("contract-" + contract.ContractId + "-" + contract.ContractVersion)}}

	response, err := cli.Execute(req)
	if err != nil {
		return "", err
	}
	err = eventResult(notifer, eventID)
	if err != nil {
		return "", err
	}
	return string(response.TransactionID), nil
}

func (t *ServiceSetup) QueryAllContractsService() (string, error) {
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryAllContracts"}

	//cli:=t.Clients[userOrgName]
	cli := t.Clients["WH-zhijianju"]
	response, err := cli.Query(req)
	if err != nil {
		return "", err
	}
	return string(response.Payload), nil

}

func (t *ServiceSetup) QueryContractByIdService(contractId string) (string, fab.TransactionID, error) {
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryContractById", Args: [][]byte{[]byte(contractId)}}

	//cli:=t.Clients[userOrgName]
	cli := t.Clients["WH-zhijianju"]
	response, err := cli.Query(req)
	if err != nil {
		return "", "", err
	}

	txID, err := t.SearchTxIDByContractId(contractId)

	return string(response.Payload), txID, nil

}
func (t *ServiceSetup) SearchTxIDByContractId(contractId string) (fab.TransactionID, error) {
	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "searchTxIDByContractId", Args: [][]byte{[]byte(contractId)}}

	//cli:=t.Clients[userOrgName]
	cli := t.Clients["WH-zhijianju"]
	response, err := cli.Query(req)
	if err != nil {
		return "", err
	}
	return fab.TransactionID(response.Payload), nil

}
