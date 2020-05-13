package main

import (


"encoding/json"
"fmt"




"github.com/hyperledger/fabric/core/chaincode/shim"
peer "github.com/hyperledger/fabric/protos/peer"

)

type LoanChainCode struct {
}

type loan struct{
ObjectType string `json:"docType"`
CustomerId string `json:"customerId"` //0
AddressId string `json:"addressId"`//1
BankId string `json:"bankId"`//2
ApplicationId string `json:"applicationId"`//3 
LoanId string `json:loanId`//4
LoanAmount string `json:loanAmt`//5
EMI string `json:emi`//6
Tenure string `json:tenure`//7
Name string `json:name`//8
Salary string `json:salary`//9
DOB string `json:dob`//10
Gender string `json:gender`//11
Collateral string `json:collateral` //12
Status string `json:status` //13
}
func main () {
err := shim.Start(new(LoanChainCode))
if err != nil{
fmt.Printf("Error starting Pets Trace chaincode: %s",err)
}
}
func (t *LoanChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response{
return shim.Success(nil)
}

func (t *LoanChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
function,args := stub.GetFunctionAndParameters()
fmt.Println("Invoke is running " + function)

if function == "CreateLoan"{
return t.CreateLoan (stub,args)
} else if function == "queryLoan" {
return t.queryLoan(stub, args)
}

fmt.Println ("Invoke did not find any func name : " +function)
return shim.Error("Received unknown function invocation")
}

func (t * LoanChainCode) CreateLoan(stub shim.ChaincodeStubInterface, args []string) peer.Response {
var err error

if len(args) != 14{
return shim.Error("Invoke chaincode with defined number of arguments")
}

if len(args[0]) <=0{
return shim.Error("1st argument must be non-empty string")
}
if len(args[1]) <= 0{
return shim.Error("2nd argument must be non-empty string")
}
if len(args[2])<=0{
return shim.Error("3rd argument must be non-empty string")
}
if len(args[3])<=0{
return shim.Error("4th argument must be non-empty string")
}
if len(args[4])<=0{
return shim.Error("tth argument must be non-empty string")
}
if len(args[5])<=0{
return shim.Error("6th argument must be non-empty string")
}
if len(args[6])<=0{
return shim.Error("7th argument must be non-empty string")
}
if len(args[7])<=0{
return shim.Error("8th argument must be non-empty string")
}
if len(args[8])<=0{
return shim.Error("9th argument must be non-empty string")
}
if len(args[9])<=0{
return shim.Error("10th argument must be non-empty string")
}
if len(args[10])<=0{
	return shim.Error("11th argument must be non-empty string")
}
if len(args[11])<=0{
	return shim.Error("12th argument must be non-empty string")
}
if len(args[12])<=0{
	return shim.Error("13th argument must be non-empty string")
}
customerId := args[0]
addressId := args[1]
bankId :=args[2]
applicationId := args[3]
loanId :=args[4] //4
loanAmt :=args[5] //5
emi:=args[6] //6
tenure := args[7] //7
name := args[8] //8
salary := args[9] //9
dob := args[10] //10
gender := args[11] //11
collateral := args[12] //12
status := args[13] //13




CustomerasBytes,err := stub.GetState(customerId)
if err != nil{
return shim.Error ("Failed to get Applicant : " + err.Error())
}
if CustomerasBytes != nil{
return shim.Error("Applicant already exist: "+ customerId)
}

objectType := "loan"
loan := &loan{objectType,customerId,addressId,bankId,applicationId,loanId,loanAmt,emi,tenure,name,salary,dob,gender,collateral,status}
ApplicantasJsonBytes, err := json.Marshal(loan)
if err != nil {
return shim.Error(err.Error())
}
err = stub.PutState(customerId,ApplicantasJsonBytes)
if err != nil{
return shim.Error("Cannot create state for Loan Applicant as: "+ err.Error())
}
fmt.Println("end of Create Applicant")
return shim.Success(nil)
}

func (t *LoanChainCode) queryLoan(stub shim.ChaincodeStubInterface, args []string) peer.Response {

if len(args) != 1 {
return shim.Error("Incorrect number of arguments. Expecting 1")
}

carAsBytes, _ := stub.GetState(args[0])
return shim.Success(carAsBytes)
}

