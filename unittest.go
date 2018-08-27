package shim
import (
	"encoding/json"
	"fmt"
	"testing"
	"github.com/hyperledger/fabric/core/chaincode/shim" 
)

func TestaddMedicalRecord (t *testing.T) {
	attributes := make(map[string][]byte)
	stub :=shim.NewCustomMockStub("mockStub", new(SampleChaincode), attributes)
	if stub == nil{
		t.Fatelf("MockStub creation failed")
	}
}
