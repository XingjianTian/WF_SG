package main

import (
	ds "WF_SG/Chaincode/DataStructure"
	"encoding/json"
	//"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	//"strconv"
	sig "WF_SG/Chaincode/Utils"
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

	if function == "addContract" { //contract
		return t.addContract(stub, args)
	} else if function == "searchTxIDByContractKey" {
		return t.searchTxIDByContractKey(stub, args)
	} else if function == "queryContractByKey" {
		return t.queryContractByKey(stub, args)
	} else if function == "queryAllContracts" {
		return t.queryAllContracts(stub)
	} else if function == "addIed" { //IED
		return t.addIed(stub, args)
	} else if function == "searchTxIDByIedId" {
		return t.searchTxIDByIedId(stub, args)
	} else if function == "queryIedByAccount" {
		return t.queryIedByAccount(stub, args)
	} else if function == "queryIedById" {
		return t.queryIedById(stub, args)
	} else if function == "queryAllIeds" {
		return t.queryAllIeds(stub)
	}

	return shim.Error("Invalid SmartContract function name")
}
func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new MySmartContract: %s", err)
	}
}

///Contract chaincode

func doubleVerify(contractJson []byte) (bool, error) {

	//verigy company
	var bid ds.BidModel
	err := json.Unmarshal(contractJson, &bid)
	if err != nil {
		return false, err
	}

	var bidBeforeCompanySig = bid
	bidBeforeCompanySig.ContractCompanyOwnerSig = ""
	bidBeforeCompanySigAsJsonBytes, err := json.Marshal(bidBeforeCompanySig)
	if err != nil {
		return false, err
	}
	bVerified, _ := sig.Verify(bidBeforeCompanySigAsJsonBytes, bid.ContractCompanyOwnerSig, bid.ContractCompanyOwnerAccount)
	if bVerified {
		fmt.Println("bid verified ok")
	} else {
		fmt.Println("bid failed to verify")
		return false, nil
	}

	//verigy user
	var contract ds.ContractModel
	err = json.Unmarshal(contractJson, &contract)
	if err != nil {
		return false, err
	}

	var contractBeforeUserSig = contract
	contractBeforeUserSig.ContractUserSig = ""
	contractBeforeUserSigAsJsonBytes, err := json.Marshal(contractBeforeUserSig)
	if err != nil {
		return false, err
	}
	bVerified, _ = sig.Verify(contractBeforeUserSigAsJsonBytes, contract.ContractUserSig,
		contract.ContractUserAccount)
	if bVerified {
		fmt.Println("contract verified ok")
	} else {
		fmt.Println("contract failed to verify")
		return false, nil
	}

	return bVerified, nil

}

//read json, make contract, add contract
func (t *SmartContract) addContract(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// args[0] are already json, args[1] is userid

	contractAsJsonBytes := []byte(args[0])
	//contractKey = "contract-"+contract.ContractId +"-"+contract.ContractVersion)
	contractKey := args[1]

	bverified, err := doubleVerify(contractAsJsonBytes)
	if bverified {
		fmt.Println("double verification passed")
	} else {
		return shim.Error("double verification failed")
	}

	//check if table and prior tables exist
	allContractKeys := queryAllContractKeys(stub)
	for _, id := range allContractKeys {
		if id == contractKey {
			return shim.Error("table " + id + " already exists")
		}
	}

	//create composite key for PutStates(), as main key
	contractIdKey, err := stub.CreateCompositeKey("Contract", []string{"_Key", contractKey})
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

//require all Contract keys
func queryAllContractKeys(stub shim.ChaincodeStubInterface) []string {
	//composite key query
	cKeyIter, err := stub.GetStateByPartialCompositeKey("Contract", []string{"_Key"})
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

func (t *SmartContract) searchTxIDByContractKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Args wrong, expecting 1 argument")
	}
	//use composite key to search
	contractKey, _ := stub.CreateCompositeKey("Contract", []string{"_Key", args[0]})

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
		return shim.Error("Args wrong, expecting 1 argument")
	}

	//use composite key to search
	contractKey, _ := stub.CreateCompositeKey("Contract", []string{"_Key", args[0]})
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
	cKeyIter, err := stub.GetStateByPartialCompositeKey("Contract", []string{"_Key"})
	if err != nil {
		fmt.Println("error")
		return shim.Error(err.Error())
	}
	defer cKeyIter.Close()

	contractsMap := make(map[string]ds.ContractModel)

	//iteration
	for i := 0; cKeyIter.HasNext(); i++ {
		responseRange, err := cKeyIter.Next()
		if err != nil {
			fmt.Println("error")
			return shim.Error(err.Error())
		}
		var contract ds.ContractModel
		err = json.Unmarshal(responseRange.Value, &contract)
		if err != nil {
			fmt.Println("error")
			return shim.Error(err.Error())
		}
		contractsMap[contract.ContractId] = contract
	}

	payload, _ := json.Marshal(contractsMap)

	fmt.Println("payload is", string(payload))

	return shim.Success(payload)
}

///IED chaincode

func (t *SmartContract) addIed(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// args[0] are already json, args[1] is ied id

	iedAsJsonBytes := []byte(args[0])
	//iedkey = IED.DeviceId(or MAC ADDRESS)
	iedKey := args[1]
	allIedKeys := queryAllIedKeys(stub)
	for _, id := range allIedKeys {
		if id == iedKey {
			return shim.Error("Ied id: " + id + " already exists")
		}
	}
	//create composite key for PutStates(), as main key
	iedCompositeKey, err := stub.CreateCompositeKey("Ied", []string{"_Id", iedKey})
	if err != nil {
		return shim.Error(err.Error())
	}

	//remarshal after changing pcheckresult
	err = stub.PutState(iedCompositeKey, iedAsJsonBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.SetEvent("eventAddIed", []byte(iedKey))
	if err != nil {
		return shim.Error(err.Error())
	}

	payload := []byte("Ied id: " + iedKey + " successfully added")
	return shim.Success(payload)
}

//require all Ied keys
func queryAllIedKeys(stub shim.ChaincodeStubInterface) []string {
	//composite key query
	cKeyIter, err := stub.GetStateByPartialCompositeKey("Ied", []string{"_Id"})
	if err != nil {
		return nil
	}
	defer cKeyIter.Close()
	keys := make([]string, 0)

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
		keys = append(keys, contractKey)
	}

	return keys
}
func (t *SmartContract) queryIedById(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("Args wrong, expecting 1 like (Ied_Id)")
	}

	//use composite key to search
	compositeKey, _ := stub.CreateCompositeKey("Ied", []string{"_Id", args[0]})
	iedAsBytes, err := stub.GetState(compositeKey)
	if err != nil {
		return shim.Error("Failed to get ied info:" + err.Error())
	} else if iedAsBytes == nil {
		return shim.Error("Ied id: " + args[0] + " does not exist")
	}
	fmt.Printf("Search Response: %s\n", string(iedAsBytes))
	return shim.Success(iedAsBytes)
}

func (t *SmartContract) queryIedByAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("Args wrong, expecting 1 like (Ied_Id)")
	}
	//composite key query
	cKeyIter, err := stub.GetStateByPartialCompositeKey("Ied", []string{"_Id"})
	if err != nil {
		fmt.Println("error")
		return shim.Error(err.Error())
	}
	for i := 0; cKeyIter.HasNext(); i++ {
		responseRange, err := cKeyIter.Next()
		if err != nil {
			fmt.Println("error")
			return shim.Error(err.Error())
		}
		var ied ds.IedModel
		err = json.Unmarshal(responseRange.Value, &ied)
		if err != nil {
			fmt.Println("error")
			return shim.Error(err.Error())
		}
		if ied.DeviceUserAccount == args[0] {
			fmt.Printf("Search Response: %s\n", string(responseRange.Value))
			return shim.Success(responseRange.Value)
		}
	}

	return shim.Error("Ied User Account: " + args[0] + " does not exist")
}

func (t *SmartContract) searchTxIDByIedId(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Args wrong, expecting 1 like (Ied_Id)")
	}
	//use composite key to search
	compositeKey, _ := stub.CreateCompositeKey("Ied", []string{"_Id", args[0]})

	resultsIterator, err := stub.GetHistoryForKey(compositeKey)
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

	//payload:= append([]byte("table id " + compositeKey + " successfully queried, the json is: "),tableAsBytes...)
	return shim.Success([]byte(TxID))

}
func (t *SmartContract) queryAllIeds(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("queryAllIeds")
	//composite key query
	cKeyIter, err := stub.GetStateByPartialCompositeKey("Ied", []string{"_Id"})
	if err != nil {
		fmt.Println("error")
		return shim.Error(err.Error())
	}
	defer cKeyIter.Close()

	iedsMap := make(map[string]ds.IedModel)

	//iteration
	for i := 0; cKeyIter.HasNext(); i++ {
		responseRange, err := cKeyIter.Next()
		if err != nil {
			fmt.Println("error")
			return shim.Error(err.Error())
		}
		var ied ds.IedModel
		err = json.Unmarshal(responseRange.Value, &ied)
		if err != nil {
			fmt.Println("error")
			return shim.Error(err.Error())
		}
		iedsMap[ied.DeviceId] = ied
	}

	payload, _ := json.Marshal(iedsMap)

	fmt.Println("payload is", string(payload))

	return shim.Success(payload)
}
