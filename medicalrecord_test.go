package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

var medicalRecord3 = "{\"ID\":\"104\",\"Name\":\"varun\",\"Weight\":\"56\",\"Age\":\"32\"}"

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

func checkState(t *testing.T, stub *shim.MockStub, id string, name string, weight string, age string) {
	bytes := stub.State[id]
	if bytes == nil {
		fmt.Println("State", id, "failed to get value")
		t.FailNow()
	}
	if string(bytes) != name {
		fmt.Println("State id", id, "was not", name, "as expected")
		t.FailNow()
	}
}

func checkQuery(t *testing.T, stub *shim.MockStub, id string, name string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("queryMedicalRecord"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Println("Query", id, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query", id, "failed to get value")
		t.FailNow()
	}
	if string(res.Payload) != name {
		fmt.Println("Query value", id, "was not", name, "as expected")
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
		checkQuery(t, stub, [][]byte{[]byte("MedicalRecord3"), []byte("103"), []byte("varun"), []byte("56"), []byte("32")})
		checkState(t, stub, "104", "varun", "56", "32")
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
		checkInvoke(t, stub, [][]byte{[]byte("addMedicalRecord"), []byte("211"), []byte("Ragu"), []byte("30"), []byte("55")})
		checkInvoke(t, stub, [][]byte{[]byte("addMedicalRecord"), []byte("221"), []byte("lambu"), []byte("40"), []byte("78")})
		//check weather the record have been updated
		checkInvoke(t, stub, [][]byte{[]byte("queryMedicalRecord"), []byte("MedicalRecord6"), []byte("211"), []byte("Ragu"), []byte("30"), []byte("55")})
		checkInvoke(t, stub, [][]byte{[]byte("queryMedicalRecord"), []byte("MedicalRecord7"), []byte("221"), []byte("lambu"), []byte("40"), []byte("78")})

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
