package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SampleChaincode struct {
}

func (s *SampleChaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// CityHandler The CityHandler
func CityHandler(res http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal("{'cities': 'San Franciso, Amsterdam, Berlin'}")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func main() {
	http.HandleFunc("/cities.json", CityHandler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
