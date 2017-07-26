package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func TestHandleIndexReturnsWithStatusOK(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	CityHandler(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Non-expected status code%v:\n\tbody: %v", "200", response.Code)
	}
}

func TestCreateLoanApplication(t *testing.T) {
	stub := shim.NewMockStub("mockstub", new(SampleChaincode))
	if stub == nil {
		t.Fatalf("MockStub creation failed")
	}
}
