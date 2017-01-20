package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChainCode struct {
}

//Init method called when creating the chaincode at the time of deployment
func (t *SimpleChainCode) Init(stub shim.ChaincodeStubInterface) ([]byte, error) {
	stub.PutState("Initial Value", []byte("Test"))
	return nil, nil
}

//Invoke method
func (t *SimpleChainCode) Invoke(stub shim.ChaincodeStubInterface) ([]byte, error) {
	function, _ := stub.GetFunctionAndParameters()

	switch function {
	case "read":
		val, _ := stub.GetState("Initial Value")
		return []byte(val), nil
	default:
		return nil, nil
	}
}

func main() {
	err := shim.Start(new(SimpleChainCode))

	if err != nil {
		fmt.Printf("Error starting simple chain code: %s", err)
	}
}
