package main

import (


"encoding/json"
"fmt"




"github.com/hyperledger/fabric/core/chaincode/shim"
peer "github.com/hyperledger/fabric/protos/peer"

)

type OnboardingChainCode struct {
}

type Onboarding struct{
ObjectType string `json:"docType"`
EmployeeId string `json:"employeeid"` //0
Name string `json:"name"`//1
Email string `json:"email"`//2
MobileNo string `json:"mobileNo"`//3 
PreviousOrg string `json:previousorg`//4
PreviousRole string `json:previousrole`//5
AddressId string `json:addressId`//6
Offered_role string `json:offered_role`//7
Joining_date string `json:joining_date`//8
Joining_location string `json:joining_location`//9
DOB string `json:dob`
Gender string `json:gender`
}
func main () {
err := shim.Start(new(OnboardingChainCode))
if err != nil{
fmt.Printf("Error starting Pets Trace chaincode: %s",err)
}
}
func (t *OnboardingChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response{
return shim.Success(nil)
}

func (t *OnboardingChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
function,args := stub.GetFunctionAndParameters()
fmt.Println("Invoke is running " + function)

if function == "CreateEmployee"{
return t.CreateEmployee (stub,args)
} else if function == "queryEmployee" {
return t.queryEmployee(stub, args)
}

fmt.Println ("Invoke did not find any func name : " +function)
return shim.Error("Received unknown function invocation")
}

func (t *OnboardingChainCode) CreateEmployee(stub shim.ChaincodeStubInterface, args []string) peer.Response {
var err error

if len(args) != 10{
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
// if len(args[12])<=0{
// 	return shim.Error("13th argument must be non-empty string")
// }
employeeId := args[0]
name := args[1]
email :=args[2]
mobileno := args[3]
previousorg :=args[4] //4
previousrole :=args[5] //5
addressid:=args[6] //6
offeredrole := args[7] //7
joining_date := args[8] //8
joining_location := args[9] //9
dob := args[10] //10
gender := args[11] //11





EmployeeasBytes,err := stub.GetState(employeeId)
if err != nil{
return shim.Error ("Failed to get Applicant : " + err.Error())
}
if EmployeeasBytes != nil{
return shim.Error("Applicant already exist: "+ employeeId)
}

objectType := "Onboarding"
Onboarding := &Onboarding{objectType,employeeId,name,email,mobileno,previousorg,previousrole,addressid,offeredrole,joining_date,joining_location,dob,gender}
ApplicantasJsonBytes, err := json.Marshal(Onboarding)
if err != nil {
return shim.Error(err.Error())
}
err = stub.PutState(employeeId,ApplicantasJsonBytes)
if err != nil{
return shim.Error("Cannot create state for Employee Applicant as: "+ err.Error())
}
fmt.Println("end of Create Applicant")
return shim.Success(nil)
}

func (t *OnboardingChainCode) queryEmployee(stub shim.ChaincodeStubInterface, args []string) peer.Response {

if len(args) != 1 {
return shim.Error("Incorrect number of arguments. Expecting 1")
}

carAsBytes, _ := stub.GetState(args[0])
return shim.Success(carAsBytes)
}
