package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var medicalRecord3 = "{\"ID\":\"104\",\"Name\":\"varun\",\"Weight\":\"56\",\"Age\":\"32\"}"
var medicalRecord2 = "{\"id\":\"103\",\"name\":\"akhil\",\"weight\":\"54\",\"age\":\"17\"}"
var newmedicalrecord = "{\"ID\":\"104\",\"Name\":\"varun\",\"Weight\":\"56\",\"Age\":\"32\"}"
var addrecord = "{\"MedicalRecord5\",\"110\",\"rajesh\",\"55\",\"65\"}"
var allrecord = "[{\"Key\":\"MedicalRecord0\", \"Record\":{\"Age\":\"22\",\"id\":\"101\",\"name\":\"shazu\",\"weight\":\"65\"}},{\"Key\":\"MedicalRecord1\", \"Record\":{\"Age\":\"24\",\"id\":\"102\",\"name\":\"rakhi\",\"weight\":\"70\"}},{\"Key\":\"MedicalRecord2\", \"Record\":{\"Age\":\"17\",\"id\":\"103\",\"name\":\"akhil\",\"weight\":\"54\"}},{\"Key\":\"MedicalRecord3\", \"Record\":{\"Age\":\"32\",\"id\":\"104\",\"name\":\"varun\",\"weight\":\"56\"}},{\"Key\":\"MedicalRecord4\", \"Record\":{\"Age\":\"23\",\"id\":\"105\",\"name\":\"anula\",\"weight\":\"90\"}}]"

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
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
		fmt.Println("Query value", Key, "was not", value, "as expected but gives", string(res.Payload))
		t.FailNow()
	}
}

func checkAllQuery(t *testing.T, stub *shim.MockStub, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("queryAllMedicalRecords")})
	if res.Status != shim.OK {
		fmt.Println("Query failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query failed to get value")
		t.FailNow()
	}
	if string(res.Payload) != value {
		fmt.Println("Query value was not", value, "as expected but gives", string(res.Payload))
		t.FailNow()
	}

}

func testinvokeAddMedicalRecord(t *testing.T, stub *shim.MockStub, key string, id string, name string, age string, weight string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("addMedicalRecord"), []byte(key), []byte(id), []byte(name), []byte(age), []byte(weight)})
	if res.Status != shim.OK {
		fmt.Println("Invoke", id, "failed", string(res.Message))
		t.FailNow()
	}

}

func testinvokeUpdateMedicalRecord(t *testing.T, stub *shim.MockStub, key string, value string, change string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("updateMedicalRecord"), []byte(key), []byte(value), []byte(change)})
	if res.Status != shim.OK {
		fmt.Println("Invoke", key, "failed", string(res.Message))
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
		// invoke init ledger
		checkInvoke(t, stub, [][]byte{[]byte("initLedger")})
		// check the existance of medical record form invoke ledger
		checkQuery(t, stub, "MedicalRecord2", "{\"id\":\"103\",\"name\":\"akhil\",\"weight\":\"54\",\"age\":\"17\"}")

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

		//add medicalrecord
		testinvokeAddMedicalRecord(t, stub, "MedicalRecord0", "110", "rajesh", "55", "65")
		//query medical record to verify that medicalrecord is added
		checkQuery(t, stub, "MedicalRecord0", "{\"id\":\"110\",\"name\":\"rajesh\",\"weight\":\"55\",\"age\":\"65\"}")

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
		//add medicalrecord
		testinvokeAddMedicalRecord(t, stub, "MedicalRecord0", "110", "rajesh", "55", "65")
		//query medical record to verify that medicalrecord is added
		checkQuery(t, stub, "MedicalRecord0", "{\"id\":\"110\",\"name\":\"rajesh\",\"weight\":\"55\",\"age\":\"65\"}")
		//update medical record (changes weight from 55=>98)
		testinvokeUpdateMedicalRecord(t, stub, "MedicalRecord0", "98", "weight")
		//check weather record has been updated
		checkQuery(t, stub, "MedicalRecord0", "{\"id\":\"110\",\"name\":\"rajesh\",\"weight\":\"98\",\"age\":\"65\"}")
		//update medical record change age to 100
		testinvokeUpdateMedicalRecord(t, stub, "MedicalRecord0", "100", "age")
		//check weather record has been updated
		checkQuery(t, stub, "MedicalRecord0", "{\"id\":\"110\",\"name\":\"rajesh\",\"weight\":\"98\",\"age\":\"100\"}")

	}

}

func TestSmartContract_queryAllMedicalRecords(t *testing.T) {
	type args struct {
		APIstub shim.ChaincodeStubInterface
	}
	{
		scc := new(SmartContract)
		stub := shim.NewMockStub("medicalrecord", scc)
		//invoke init ledger to add 5 medical records
		checkInvoke(t, stub, [][]byte{[]byte("initLedger")})
		//query all medical medical records and compre with init ledger record
		checkAllQuery(t, stub, "[{\"Key\":\"MedicalRecord0\", \"Record\":{\"id\":\"101\",\"name\":\"shazu\",\"weight\":\"65\",\"age\":\"22\"}},{\"Key\":\"MedicalRecord1\", \"Record\":{\"id\":\"102\",\"name\":\"rakhi\",\"weight\":\"70\",\"age\":\"24\"}},{\"Key\":\"MedicalRecord2\", \"Record\":{\"id\":\"103\",\"name\":\"akhil\",\"weight\":\"54\",\"age\":\"17\"}},{\"Key\":\"MedicalRecord3\", \"Record\":{\"id\":\"104\",\"name\":\"varun\",\"weight\":\"56\",\"age\":\"32\"}},{\"Key\":\"MedicalRecord4\", \"Record\":{\"id\":\"105\",\"name\":\"anula\",\"weight\":\"90\",\"age\":\"23\"}}]")
		//add 6th record to the blockchain
		testinvokeAddMedicalRecord(t, stub, "MedicalRecord5", "110", "rajesh", "55", "65")
		//do query record and check weather it gives updated allmedicalrecord with 6th record
		checkAllQuery(t, stub, "[{\"Key\":\"MedicalRecord0\", \"Record\":{\"id\":\"101\",\"name\":\"shazu\",\"weight\":\"65\",\"age\":\"22\"}},{\"Key\":\"MedicalRecord1\", \"Record\":{\"id\":\"102\",\"name\":\"rakhi\",\"weight\":\"70\",\"age\":\"24\"}},{\"Key\":\"MedicalRecord2\", \"Record\":{\"id\":\"103\",\"name\":\"akhil\",\"weight\":\"54\",\"age\":\"17\"}},{\"Key\":\"MedicalRecord3\", \"Record\":{\"id\":\"104\",\"name\":\"varun\",\"weight\":\"56\",\"age\":\"32\"}},{\"Key\":\"MedicalRecord4\", \"Record\":{\"id\":\"105\",\"name\":\"anula\",\"weight\":\"90\",\"age\":\"23\"}},{\"Key\":\"MedicalRecord5\", \"Record\":{\"id\":\"110\",\"name\":\"rajesh\",\"weight\":\"55\",\"age\":\"65\"}}]")

	}

}

func TestSmartContract_Invoke(t *testing.T) {
	type args struct {
		APIstub shim.ChaincodeStubInterface
	}
	{
		scc := new(SmartContract)
		stub := shim.NewMockStub("medicalrecord", scc)
		// invoke init ledger
		checkInvoke(t, stub, [][]byte{[]byte("initLedger")})
		// check the existance of medical record form invoke ledger
		checkQuery(t, stub, "MedicalRecord0", "{\"id\":\"101\",\"name\":\"shazu\",\"weight\":\"65\",\"age\":\"22\"}")
		checkQuery(t, stub, "MedicalRecord1", "{\"id\":\"102\",\"name\":\"rakhi\",\"weight\":\"70\",\"age\":\"24\"}")
		checkQuery(t, stub, "MedicalRecord2", "{\"id\":\"103\",\"name\":\"akhil\",\"weight\":\"54\",\"age\":\"17\"}")
		checkQuery(t, stub, "MedicalRecord3", "{\"id\":\"104\",\"name\":\"varun\",\"weight\":\"56\",\"age\":\"32\"}")
		checkQuery(t, stub, "MedicalRecord4", "{\"id\":\"105\",\"name\":\"anula\",\"weight\":\"90\",\"age\":\"23\"}")
		//add medicalrecord
		testinvokeAddMedicalRecord(t, stub, "MedicalRecord5", "110", "rajesh", "55", "65")
		//query medical record to verify that medicalrecord is added
		checkQuery(t, stub, "MedicalRecord5", "{\"id\":\"110\",\"name\":\"rajesh\",\"weight\":\"55\",\"age\":\"65\"}")
		//update medical record (changes weight from 55=>98)
		testinvokeUpdateMedicalRecord(t, stub, "MedicalRecord5", "98", "weight")
		//check weather record has been updated
		checkQuery(t, stub, "MedicalRecord5", "{\"id\":\"110\",\"name\":\"rajesh\",\"weight\":\"98\",\"age\":\"65\"}")
		//update medical record change age to 100
		testinvokeUpdateMedicalRecord(t, stub, "MedicalRecord5", "100", "age")
		//check weather record has been updated
		checkQuery(t, stub, "MedicalRecord5", "{\"id\":\"110\",\"name\":\"rajesh\",\"weight\":\"98\",\"age\":\"100\"}")
	}
}

func TestSmartContract_initLedger(t *testing.T) {
	type args struct {
		APIstub shim.ChaincodeStubInterface
	}
	{
		scc := new(SmartContract)
		stub := shim.NewMockStub("medicalrecord", scc)
		// invoke init ledger
		checkInvoke(t, stub, [][]byte{[]byte("initLedger")})
		// check the existance of medical record form invoke ledger
		checkQuery(t, stub, "MedicalRecord0", "{\"id\":\"101\",\"name\":\"shazu\",\"weight\":\"65\",\"age\":\"22\"}")
		checkQuery(t, stub, "MedicalRecord1", "{\"id\":\"102\",\"name\":\"rakhi\",\"weight\":\"70\",\"age\":\"24\"}")
		checkQuery(t, stub, "MedicalRecord2", "{\"id\":\"103\",\"name\":\"akhil\",\"weight\":\"54\",\"age\":\"17\"}")
		checkQuery(t, stub, "MedicalRecord3", "{\"id\":\"104\",\"name\":\"varun\",\"weight\":\"56\",\"age\":\"32\"}")
		checkQuery(t, stub, "MedicalRecord4", "{\"id\":\"105\",\"name\":\"anula\",\"weight\":\"90\",\"age\":\"23\"}")
	}
}
