package main

import (
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type CustomMockStub struct {
	stub           *MockStub
	CertAttributes map[string][]byte
}

var newmr = `{"id":"110","name":"rajest","weight":"70","age":"21"}`

//test weather addmedicalrecord is invoked sussesfully
func TestaddMedicalRecord(t *testing.T) {
	attributes := make(map[string][]byte)
	attributes["id"] = []byte("218")
	attributes["name"] = []byte("raksho")
	attributes["weight"] = []byte("76")
	attributes["age"] = []byte("45")

	stub := shim.NewCustomMockStub("mockStub", new(SampleChaincode), attributes)
	if stub == nil {
		t.Fatelf("MockStub creation failed")
	}
	_, err := stub.MockInvoke("addMedicalRecord", []string{})
	if err != nil {
		t.Fatalf("Expected addMedicalRecord to be invoked")
	}
}

//
