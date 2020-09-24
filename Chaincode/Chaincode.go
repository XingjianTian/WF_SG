package main

import (
	"encoding/json"
	//"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	//"strconv"
	ds "WF_SG/DataStructure"
	sig "WF_SG/Utils"
	"strings"
)

type SmartContract struct {
}

func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Chaincode Init")
	return shim.Success(nil)
}
func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Chaincode Invoke")

	function, args := stub.GetFunctionAndParameters()
	if function == "addTable" {
		return t.addTable(stub, args)
	} else if function == "searchTableById" {
		return t.searchTableById(stub, args)
	} else if function == "searchTxIDByTableId" {
		return t.searchTxIDByTableId(stub, args)
	} else if function == "queryAllTables" {
		return t.queryAllTables(stub)
	} else if function == "queryAllTablesWithoutExclude" {
		return t.queryAllTablesWithoutExclude(stub)
	}

	return shim.Error("Invalid SmartContract function name")
}
func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new MySmartContract: %s", err)
	}
}

//read json, make table, add table
func (t *SmartContract) addTable(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// args[0] are already json, args[1] is userid

	tableAsJsonBytes := []byte(args[0])
	userid := args[1]

	var table ds.Table
	var tableBeforeSig ds.Table

	//change pcheckResult and remarshal

	err := json.Unmarshal(tableAsJsonBytes, &table)
	table.CheckAllPropertyValue()
	if err != nil {
		return shim.Error(err.Error())
	}

	err = json.Unmarshal(tableAsJsonBytes, &tableBeforeSig)
	if err != nil {
		return shim.Error(err.Error())
	}

	lastSignature := table.Sig[len(table.Sig)-1]

	if len(table.Sig) == 1 {
		tableBeforeSig.Sig = nil
	} else {
		tableBeforeSig.Sig = table.Sig[:len(table.Sig)-1]
	}

	tableBeforeSigAsJsonBytes, err := json.Marshal(tableBeforeSig)
	if err != nil {
		return shim.Error(err.Error())
	}

	//use public key to verify
	bVerified, _ := sig.Verify(tableBeforeSigAsJsonBytes, lastSignature, userid)
	if bVerified {
		fmt.Println("user: " + userid + " verified ok")
	} else {
		return shim.Error("user " + userid + " fail to verify")
	}

	//check if table and prior tables exist
	allTableIds := queryAllTableIds(stub)
	for _, id := range allTableIds {
		if id == table.TId {
			return shim.Error("table " + id + " already exists")
		}
	}

	//create composite key for PutStates(), as main key
	tableIdKey, err := stub.CreateCompositeKey("Table", []string{"_TId", table.TId})
	if err != nil {
		return shim.Error(err.Error())
	}

	//remarshal after changing pcheckresult
	tableAsJsonBytes, err = json.Marshal(table)
	err = stub.PutState(tableIdKey, tableAsJsonBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.SetEvent("eventAddTable", []byte(table.TId))
	if err != nil {
		return shim.Error(err.Error())
	}

	payload := []byte("table id: " + table.TId + " successfully added")
	return shim.Success(payload)
}

//require all table id
func queryAllTableIds(stub shim.ChaincodeStubInterface) []string {
	//composite key query
	cKeyIter, err := stub.GetStateByPartialCompositeKey("Table", []string{"_TId"})
	if err != nil {
		return nil
	}
	defer cKeyIter.Close()
	tableIds := make([]string, 0)

	//iteration
	for i := 0; cKeyIter.HasNext(); i++ {
		responseRange, err := cKeyIter.Next()
		if err != nil {
			return nil
		}
		_, cKeyParts, err := stub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return nil
		}
		tableId := cKeyParts[1]
		tableIds = append(tableIds, tableId)
	}

	return tableIds
}

func (t *SmartContract) searchTxIDByTableId(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Args wrong, expecting 1 like (table_Id)")
	}
	tableIdAsStr := args[0]
	//use composite key to search
	tableIdKey, _ := stub.CreateCompositeKey("Table", []string{"_TId", tableIdAsStr})

	resultsIterator, err := stub.GetHistoryForKey(tableIdKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var TxID string
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		TxID = response.TxId
	}

	//payload:= append([]byte("table id " + tableIdAsStr + " successfully queried, the json is: "),tableAsBytes...)
	return shim.Success([]byte(TxID))

}

func (t *SmartContract) searchTableById(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("Args wrong, expecting 1 like (table_Id)")
	}
	tableIdAsStr := args[0]

	//use composite key to search
	tableIdKey, _ := stub.CreateCompositeKey("Table", []string{"_TId", tableIdAsStr})
	tableAsBytes, err := stub.GetState(tableIdKey)
	if err != nil {
		return shim.Error("Failed to get table info:" + err.Error())
	} else if tableAsBytes == nil {
		return shim.Error("Table does not exist")
	}
	fmt.Printf("Search Response: %s\n", string(tableAsBytes))

	//payload:= append([]byte("table id " + tableIdAsStr + " successfully queried, the json is: "),tableAsBytes...)
	return shim.Success(tableAsBytes)
}
func (t *SmartContract) queryAllTables(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("queryAllTables")
	//composite key query
	cKeyIter, err := stub.GetStateByPartialCompositeKey("Table", []string{"_TId"})
	if err != nil {
		fmt.Println("error000")
		return shim.Error(err.Error())
	}
	defer cKeyIter.Close()

	tablesMap := make(map[string]ds.TableForWebinCC)

	//iteration
	for i := 0; cKeyIter.HasNext(); i++ {
		responseRange, err := cKeyIter.Next()
		if err != nil {
			fmt.Println("error1")
			return shim.Error(err.Error())
		}
		var table ds.Table
		err = json.Unmarshal(responseRange.Value, &table)
		if err != nil {
			fmt.Println("error2")
			return shim.Error(err.Error())
		}

		var tempState string
		var orgnization string
		if strings.Contains(table.LastSigner, "builder") {
			tempState = "等待Supervisor签名"
			orgnization = "builder"
		} else if strings.Contains(table.LastSigner, "supervisor") {
			tempState = "等待Constructor签名"
			orgnization = "supervisor"
		} else if strings.Contains(table.LastSigner, "constructor") {
			tempState = "已完成"
			orgnization = "constructor"
		}

		fmt.Println("Thestate is", tempState)

		tableForWeb := ds.TableForWebinCC{
			TID:                table.TId,
			TName:              table.TName,
			OrgEngineeringName: table.TCommon.OrgEngineeringName,
			DepEngineeringName: table.TCommon.DepEngineeringName,
			SubEngineeringName: table.TCommon.SubEngineeringName,
			TestPart:           table.TCommon.TestPart,
			State:              tempState,
			CreatedAt:          table.TimeStamp,
			Operator:           table.LastSigner,
			OrgName:            orgnization,
		}
		//tablesForWebInit = append(tablesForWebInit,tableForWeb)

		/*查看元素在集合中是否存在 */

		realID := ds.AiteBefore(table.TId)
		fmt.Println("realid is", realID)

		_, ok := tablesMap[realID] /*如果确定是真实的,则存在,否则不存在 */
		/*fmt.Println(capital) */
		/*fmt.Println(ok) */
		if ok {
			if tableForWeb.CreatedAt.After(tablesMap[realID].CreatedAt) {
				tablesMap[realID] = tableForWeb
			}
		} else {
			tablesMap[realID] = tableForWeb
		}

		if err != nil {
			return shim.Error(err.Error())
		}
	}

	payload, _ := json.Marshal(FormatTablesForWeb(tablesMap))

	fmt.Println("payload is", string(payload))

	return shim.Success(payload)
}

func (t *SmartContract) queryAllTablesWithoutExclude(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("queryAllTablesWithoutExclude")
	//composite key query
	cKeyIter, err := stub.GetStateByPartialCompositeKey("Table", []string{"_TId"})
	if err != nil {
		fmt.Println("error000")
		return shim.Error(err.Error())
	}
	defer cKeyIter.Close()

	var tablesList []ds.TableForWebinCC

	//iteration
	for i := 0; cKeyIter.HasNext(); i++ {
		responseRange, err := cKeyIter.Next()
		if err != nil {
			fmt.Println("error1")
			return shim.Error(err.Error())
		}
		var table ds.Table
		err = json.Unmarshal(responseRange.Value, &table)
		if err != nil {
			fmt.Println("error2")
			return shim.Error(err.Error())
		}

		var tempState string
		var orgnization string
		if strings.Contains(table.LastSigner, "builder") {
			tempState = "等待Supervisor签名"
			orgnization = "builder"
		} else if strings.Contains(table.LastSigner, "supervisor") {
			tempState = "等待Constructor签名"
			orgnization = "supervisor"
		} else if strings.Contains(table.LastSigner, "constructor") {
			tempState = "已完成"
			orgnization = "constructor"
		}

		fmt.Println("Thestate is", tempState)

		tableForWeb := ds.TableForWebinCC{
			TID:                table.TId,
			TName:              table.TName,
			OrgEngineeringName: table.TCommon.OrgEngineeringName,
			DepEngineeringName: table.TCommon.DepEngineeringName,
			SubEngineeringName: table.TCommon.SubEngineeringName,
			TestPart:           table.TCommon.TestPart,
			State:              tempState,
			CreatedAt:          table.TimeStamp,
			Operator:           table.LastSigner,
			OrgName:            orgnization,
		}
		//tablesForWebInit = append(tablesForWebInit,tableForWeb)

		/*查看元素在集合中是否存在 */

		comma := strings.Index(table.TId, "@")
		realID := table.TId[0:comma]
		fmt.Println("realid is", realID)
		tableForWeb.RealID = realID

		tablesList = append(tablesList, tableForWeb)

	}

	payload, _ := json.Marshal(tablesList)

	fmt.Println("payload is", string(payload))

	return shim.Success(payload)
}

func FormatTablesForWeb(tablesMap map[string]ds.TableForWebinCC) []ds.TableForWebinCC {

	var result []ds.TableForWebinCC

	for k, v := range tablesMap {
		temp := v
		temp.RealID = k
		result = append(result, temp)
	}
	return result

}

//customize searching, given id,output container or property
