package Services

import (
	ds "WF_SG/DataStructure"
	sig "WF_SG/Utils"
	"encoding/json"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"strings"
	"time"
)

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
