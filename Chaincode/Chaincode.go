package main

import (
	"encoding/json"
	//"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	//"strconv"
	sig "WF_SG/Utils"
	"WF_SG/Web/models"
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
	if function == "addContract" {
		return t.addContract(stub, args)
	} else if function == "searchTxIDByContractId" {
		return t.SearchTxIDByKey(stub, args)
	} else if function == "queryContractByKey" {
		return t.queryContractByKey(stub, args)
	} else if function == "queryAllContracts" {
		return t.queryAllContracts(stub)
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
func (t *SmartContract) addContract(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// args[0] are already json, args[1] is userid

	contractAsJsonBytes := []byte(args[0])
	//contractKey = "contract-"+contract.ContractId +"-"+contract.ContractVersion)
	contractKey := args[1]

	var contract models.ContractModel
	var contractBeforeSig models.ContractModel

	//change pcheckResult and remarshal

	err := json.Unmarshal(contractAsJsonBytes, &contractBeforeSig)
	err = json.Unmarshal(contractAsJsonBytes, &contract)
	if err != nil {
		return shim.Error(err.Error())
	}
	contractBeforeSig.ContractCompanySig = ""

	contractBeforeSigAsJsonBytes, err := json.Marshal(contractBeforeSig)
	if err != nil {
		return shim.Error(err.Error())
	}

	//use public key to verify
	bVerified, _ := sig.Verify(contractBeforeSigAsJsonBytes, contract.ContractCompanySig, contract.ContractCompanyName)
	if bVerified {
		fmt.Println("contract verified ok")
	} else {
		return shim.Error("contract failed to verify")
	}

	//check if table and prior tables exist
	allContractKeys := queryAllContractKeys(stub)
	for _, id := range allContractKeys {
		if id == contractKey {
			return shim.Error("table " + id + " already exists")
		}
	}

	//create composite key for PutStates(), as main key
	contractIdKey, err := stub.CreateCompositeKey("Table", []string{"_TId", contractKey})
	if err != nil {
		return shim.Error(err.Error())
	}

	//remarshal after changing pcheckresult
	err = stub.PutState(contractIdKey, contractAsJsonBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.SetEvent("eventAddTable", []byte(contractKey))
	if err != nil {
		return shim.Error(err.Error())
	}

	payload := []byte("table id: " + contractKey + " successfully added")
	return shim.Success(payload)
}

//require all table id
func queryAllContractKeys(stub shim.ChaincodeStubInterface) []string {
	//composite key query
	cKeyIter, err := stub.GetStateByPartialCompositeKey("Table", []string{"_Id"})
	if err != nil {
		return nil
	}
	defer cKeyIter.Close()
	contractKeys := make([]string, 0)

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
		contractKey := cKeyParts[1]
		contractKeys = append(contractKeys, contractKey)
	}

	return contractKeys
}

func (t *SmartContract) SearchTxIDByKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Args wrong, expecting 1 like (table_Id)")
	}
	//use composite key to search
	contractKey, _ := stub.CreateCompositeKey("Table", []string{"_TId", args[0]})

	resultsIterator, err := stub.GetHistoryForKey(contractKey)
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

	//payload:= append([]byte("table id " + contractKey + " successfully queried, the json is: "),tableAsBytes...)
	return shim.Success([]byte(TxID))

}

func (t *SmartContract) queryContractByKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("Args wrong, expecting 1 like (contract_Id)")
	}

	//use composite key to search
	contractKey, _ := stub.CreateCompositeKey("Table", []string{"_TId", args[0]})
	contractAsBytes, err := stub.GetState(contractKey)
	if err != nil {
		return shim.Error("Failed to get contract info:" + err.Error())
	} else if contractAsBytes == nil {
		return shim.Error("Contract does not exist")
	}
	fmt.Printf("Search Response: %s\n", string(contractAsBytes))

	//payload:= append([]byte("table id " + tableIdAsStr + " successfully queried, the json is: "),tableAsBytes...)
	return shim.Success(contractAsBytes)
}
func (t *SmartContract) queryAllContracts(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("queryAllContracts")
	//composite key query
	cKeyIter, err := stub.GetStateByPartialCompositeKey("Table", []string{"_Id"})
	if err != nil {
		fmt.Println("error")
		return shim.Error(err.Error())
	}
	defer cKeyIter.Close()

	contractsMap := make(map[string]models.ContractModel)

	//iteration
	for i := 0; cKeyIter.HasNext(); i++ {
		responseRange, err := cKeyIter.Next()
		if err != nil {
			fmt.Println("error")
			return shim.Error(err.Error())
		}
		var contract models.ContractModel
		err = json.Unmarshal(responseRange.Value, &contract)
		if err != nil {
			fmt.Println("error")
			return shim.Error(err.Error())
		}

		_, ok := contractsMap[contract.ContractId] /*如果确定是真实的,则存在,否则不存在 */
		/*fmt.Println(capital) */
		/*fmt.Println(ok) */
		if ok {
			fmt.Println("map created ok")
		} else {
			fmt.Println("failed to create map")
		}

		if err != nil {
			return shim.Error(err.Error())
		}
	}

	payload, _ := json.Marshal(contractsMap)

	fmt.Println("payload is", string(payload))

	return shim.Success(payload)
}
