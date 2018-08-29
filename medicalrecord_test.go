package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

var medicalRecord3 = "{\"ID\":\"104\",\"Name\":\"varun\",\"Weight\":\"56\",\"Age\":\"32\"}"
var medicalRecord2 = "{\"id\":\"103\",\"name\":\"akhil\",\"weight\":\"54\",\"age\":\"17\"}"
var newmedicalrecord = "{\"ID\":\"104\",\"Name\":\"varun\",\"Weight\":\"56\",\"Age\":\"32\"}"
var addrecord = "{\"MedicalRecord5\",\"110\",\"rajesh\",\"55\",\"65\"}"

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}

func checkInvoke2(t *testing.T, stub *shim.MockStub, args [][]byte, response []byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}

	if response != nil {
		if res.Payload == nil {
			fmt.Printf("Invoke returned nil, expected %s", string(response))
			t.FailNow()
		}
		if string(res.Payload) != string(response) {
			fmt.Printf("Invoke returned %s, expected %s", string(res.Payload), string(response))
			t.FailNow()
		}
	}
}

func checkQuery(t *testing.T, stub *shim.MockStub, Key string, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("queryMedicalRecord"), []byte(Key)})
	if res.Status != shim.OK {
		fmt.Println("Query", Key, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query", Key, "failed to get value")
		t.FailNow()
	}
	if string(res.Payload) != value {
		fmt.Println("Query value", Key, "was not", value, "as expected", string(res.Payload))
		t.FailNow()
	}
}

func TestSmartContract_queryMedicalRecord(t *testing.T) {
	type args struct {
		APIstub shim.ChaincodeStubInterface
		args    []string
	}

	{
		scc := new(SmartContract)
		stub := shim.NewMockStub("medicalrecord", scc)
		checkInvoke(t, stub, [][]byte{[]byte("initLedger")})

		checkQuery(t, stub, "MedicalRecord2", "{\"id\":\"103\",\"name\":\"akhil\",\"weight\":\"54\",\"age\":\"17\"}")
		// TODO: Add test cases.
	}

}

func TestSmartContract_addMedicalRecord(t *testing.T) {
	type args struct {
		APIstub shim.ChaincodeStubInterface
		args    []string
	}

	{
		scc := new(SmartContract)
		stub := shim.NewMockStub("medicalrecord", scc)

		//add medicalrecords
		checkInvoke(t, stub, [][]byte{[]byte("addMedicalRecord"), []byte("{\"MedicalRecord5\",\"110\",\"rajesh\",\"55\",\"65\"}")})

		// TODO: Add test cases.
	}

}

func TestSmartContract_updateMedicalRecord(t *testing.T) {
	type args struct {
		APIstub shim.ChaincodeStubInterface
		args    []string
	}

	{
		scc := new(SmartContract)
		stub := shim.NewMockStub("medicalrecord", scc)
		//add a medicalrecord
		checkInvoke(t, stub, [][]byte{[]byte("addMedicalRecord"), []byte("211"), []byte("Ragu"), []byte("30"), []byte("55")})
		//check weather medicalrecord is added
		checkInvoke(t, stub, [][]byte{[]byte("queryMedicalRecord"), []byte("MedicalRecord6"), []byte("211"), []byte("Ragu"), []byte("30"), []byte("55")})
		//update medicalrecord weight
		checkInvoke(t, stub, [][]byte{[]byte("updateMedicalRecord"), []byte("MedicalRecord6"), []byte("89")})
		//check weather weight is updated
		checkInvoke(t, stub, [][]byte{[]byte("queryMedicalRecord"), []byte("MedicalRecord6"), []byte("211"), []byte("Ragu"), []byte("30"), []byte("89")})
		// TODO: Add test cases.
	}

}

func TestSmartContract_queryAllMedicalRecords(t *testing.T) {
	type args struct {
		APIstub shim.ChaincodeStubInterface
	}
	tests := []struct {
		name string
		s    *SmartContract
		args args
		want sc.Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SmartContract{}
			if got := s.queryAllMedicalRecords(tt.args.APIstub); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SmartContract.queryAllMedicalRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmartContract_Invoke(t *testing.T) {
	type args struct {
		APIstub shim.ChaincodeStubInterface
	}
	tests := []struct {
		name string
		s    *SmartContract
		args args
		want sc.Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SmartContract{}
			if got := s.Invoke(tt.args.APIstub); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SmartContract.Invoke() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmartContract_initLedger(t *testing.T) {
	type args struct {
		APIstub shim.ChaincodeStubInterface
	}
	tests := []struct {
		name string
		s    *SmartContract
		args args
		want sc.Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SmartContract{}
			if got := s.initLedger(tt.args.APIstub); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SmartContract.initLedger() = %v, want %v", got, tt.want)
			}
		})
	}
}
