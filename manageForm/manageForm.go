/*/*
-a-Licensed to the Apache Software Foundation (ASF) under one
or more contributor license Forms.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
"errors"
"fmt"
"strconv"
"encoding/json"

"github.com/hyperledger/fabric/core/chaincode/shim"
"github.com/hyperledger/fabric/core/util"
)

// ManageForm example simple Chaincode implementation
type ManageForm struct {
}

var Tier3FormIndexStr = "_Tier3Formindex"				//name for the key/value that will store a list of all known  Tier3 Form
var Tier2FormIndexStr = "_Tier2Formindex"				//name for the key/value that will store a list of all known Tier2 Form
var Tier1FormIndexStr = "_Tier1Formindex"				//name for the key/value that will store a list of all known Tier1 Form
var OEMFormIndexStr = "_OEMFormindex"				//name for the key/value that will store a list of all known OEM Form

type Form struct{
								// Attributes of a Form 
	FAA_FormNumber string `json:"FAA_formNumber"`	
	Quantity string `json:"quantity"`
	FAA_FormURL string `json:"FAA_formUrl"`
	File_hash string `json:"fileHash"`
	User string `json:"user"`					
	ItemType string `json:"itemType"`
	Part_number string `json:"part_number"`
	Total_approvedQty string `json:"total_approvedQty"`
	ApprovalDate string `json:"approvalDate"`	
	Authorization_number string `json:"authorization_number"`
	Tier3_Form_number string `json:"tier3_Form_number"`
	Tier2_Form_number string `json:"tier2_Form_number"`
	Tier1_Form_number string `json:"tier1_Form_number"`
	UserType string `json:"userType"`

}
type Shipment struct{
								// Attributes of a Shipment 

	ShipmentID string `json:"shipmentId"`
	Description string `json:"description"`
	Sender string `json:"sender"`
	SenderType string `json:"senderType"`		
	Receiver string `json:"receiver"`
	ReceiverType string `json:"receiverType"`
	FAA_FormNumber string `json:"FAA_formNumber"`	
	Quantity string `json:"quantity"`
	ShipmentDate string `json:"shipmentDate"`	
	ReceivedDate string `json:"receivedDate"`
	Status string `json:"status"`
	ChaincodeURL string `json:"chaincodeURL"`
	Ship_frm_country string `json:"ship_frm_country"`
	Ship_frm_city string `json:"ship_frm_city"`
	Ship_to_country string `json:"ship_to_country"`
	Ship_to_city string `json:"ship_to_city"`
	Truck_details string `json:"truck_details"`
	Logistics_agency_details string `json:"logistics_agency_details"`
	Air_ship_way_bill_details string `json:"air_ship_way_bill_way_details"`
	Flight_vessel_details string `json:"flight_vessel_details"`
	Departing_port string `json:"departing_port"`
	Arriving_port string `json:"arriving_port"`
	Scheduled_departure_date_ts string `json:"scheduled_departure_date_ts"`
	Actual_arrival_date_ts	string `json:"actual_arriving_date_ts"`
	Vendor_name string `json:"vendor_name"`
	Tier_type string `json:"tier_type"`
	File_hash string `json:"file_hash"`
	
	
	
	
	
	
}
// ============================================================================================================================
// Main - start the chaincode for Form management
// ============================================================================================================================
func main() {			
	err := shim.Start(new(ManageForm))
	if err != nil {
		fmt.Printf("Error starting Form management chaincode: %s", err)
	}
}
// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================
func (t *ManageForm) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var msg string
	var err error
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	// Initialize the chaincode
	msg = args[0]
	fmt.Println("ManageForm chaincode is deployed successfully.");
	
	// Write the state to the ledger
	err = stub.PutState("abc", []byte(msg))				//making a test var "abc", I find it handy to read/write to it right away to test the network
	if err != nil {
		return nil, err
	}
	
	var emptyOEM []string
	OEMjsonAsBytes, _ := json.Marshal(emptyOEM)								//marshal an emtpy array of strings to clear the index
	err = stub.PutState(OEMFormIndexStr, OEMjsonAsBytes)
	if err != nil {
		return nil, err
	}
	
	var Tier3empty []string
	Tier3jsonAsBytes, _ := json.Marshal(Tier3empty)								//marshal an emtpy array of strings to clear the index
	err = stub.PutState(Tier3FormIndexStr, Tier3jsonAsBytes)
	if err != nil {
		return nil, err
	}

	var Tier2empty []string
	Tier2jsonAsBytes, _ := json.Marshal(Tier2empty)								//marshal an emtpy array of strings to clear the index
	err = stub.PutState(Tier2FormIndexStr, Tier2jsonAsBytes)
	if err != nil {
		return nil, err
	}

	var Tier1empty []string
	Tier1jsonAsBytes, _ := json.Marshal(Tier1empty)								//marshal an emtpy array of strings to clear the index
	err = stub.PutState(Tier1FormIndexStr, Tier1jsonAsBytes)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
// ============================================================================================================================
// Run - Our entry Formint for Invocations - [LEGACY] obc-peer 4/25/2016
// ============================================================================================================================
func (t *ManageForm) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("run is running " + function)
	return t.Invoke(stub, function, args)
}
// ============================================================================================================================
// Invoke - Our entry Formint for Invocations
// ============================================================================================================================
func (t *ManageForm) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "createForm_Tier3" {											//create a new Form
		return t.createForm_Tier3(stub, args)
	} else if function == "createForm_Tier2" {											//create a new Form
		return t.createForm_Tier2(stub, args)
	}else if function == "createForm_Tier1" {											//create a new Form
		return t.createForm_Tier1(stub, args)
	}else if function == "createForm_OEM" {											//create a new Form
		return t.createForm_OEM(stub, args)
	}else if function == "update_Form" {									//update a Form
		return t.update_Form(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)	
	jsonResp := "Error : Received unknown function invocation: "+ function 				//error
	return nil, errors.New(jsonResp)
}
// ============================================================================================================================
// Query - Our entry Formint for Queries
// ============================================================================================================================
func (t *ManageForm) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "getForm_byID" {													//Read a Form by FormID
		return t.getForm_byID(stub, args[0])
	} else if function == "getForm_byUser" {													//Read a Form by Buyer
		return t.getForm_byUser(stub, args)
	} else if function == "get_AllForm" {													//Read all Forms
		return t.get_AllForm(stub, args)
	} else if function == "get_FormId_ByTier" {													//Read a Shipment by Buyer
		return t.get_FormId_ByTier(stub, args)
	}else if function == "get_AllFormByTier" {													//Read a Shipment by Buyer
		return t.get_AllFormByTier(stub, args)
	}
	

	fmt.Println("query did not find func: " + function)				//error
	jsonResp := "Error : Received unknown function query: "+ function 
	return nil, errors.New(jsonResp)
}
// ============================================================================================================================
// getForm_byID - get Form details for a specific FormID from chaincode state
// ============================================================================================================================
func (t *ManageForm) getForm_byID(stub shim.ChaincodeStubInterface, args string) ([]byte, error) {
	//getForm_byID('FAA_formNumber')
	var FAA_formNumber, jsonResp string
	var err error
	fmt.Println("Fetching Form by FAA_formNumber")
	if args == "" {
		return nil, errors.New("Incorrect number of arguments. Expecting Form ID to query")
	}
	// set FAA_formNumber
	FAA_formNumber = args
	valAsbytes, err := stub.GetState(FAA_formNumber)									//get the FAA_formNumber from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + FAA_formNumber + "\"}"
		return nil, errors.New(jsonResp)
	}
	fmt.Print("valAsbytes : ")
	fmt.Println(valAsbytes)
	fmt.Println("Fetched Form by FAA_formNumber")
	return valAsbytes, nil													//send it onward
}




func (t *ManageForm) get_FormId_ByTier(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var jsonResp string
	fmt.Println("Fetching All Form IDS by Tier Type")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments,Expecting one argument")
	}
	
	if args[0]=="Tier-3"{
		var Tier3FormIndex []string
		Tier3FormAsBytes, err := stub.GetState(Tier3FormIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get Tier3 Form index")
		}
		fmt.Print("Tier3FormAsBytes : ")
		fmt.Println(Tier3FormAsBytes)
		json.Unmarshal(Tier3FormAsBytes, &Tier3FormIndex)								//un stringify it aka JSON.parse()
		fmt.Print("Tier3FormIndex : ")
		fmt.Println(Tier3FormIndex)
		jsonResp = "{\"FormList\":\"["
		for i,val := range Tier3FormIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tier3 Forms")
			jsonResp = jsonResp + val
			if i < len(Tier3FormIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
		
		fmt.Println("len(Tier3FormIndex) : ")
		fmt.Println(len(Tier3FormIndex))
		jsonResp = jsonResp +"]\""+ "}"
		fmt.Println([]byte(jsonResp))
		fmt.Println("Fetched All Tier3 Formssuccessfully.")
		return []byte(jsonResp), nil	
	}
	
	if args[0]=="Tier-2"{
		var Tier2FormIndex []string
		Tier2FormAsBytes, err := stub.GetState(Tier2FormIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get Tier32Form index")
		}
		fmt.Print("Tier2FormAsBytes : ")
		fmt.Println(Tier2FormAsBytes)
		json.Unmarshal(Tier2FormAsBytes, &Tier2FormIndex)								//un stringify it aka JSON.parse()
		fmt.Print("Tier2FormIndex : ")
		fmt.Println(Tier2FormIndex)
		jsonResp = "{\"FormList\":\"["
		for i,val := range Tier2FormIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tier2 Forms")
			jsonResp = jsonResp + val
			if i < len(Tier2FormIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
		
		fmt.Println("len(Tier2FormIndex) : ")
		fmt.Println(len(Tier2FormIndex))
		jsonResp = jsonResp +"]\""+ "}"
		fmt.Println([]byte(jsonResp))
		fmt.Println("Fetched All Tier2 Formssuccessfully.")
		return []byte(jsonResp), nil	
	}
	if args[0]=="Tier-1"{
		var Tier1FormIndex []string
		Tier1FormAsBytes, err := stub.GetState(Tier1FormIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get Tier1 Form index")
		}
		fmt.Print("Tier1FormAsBytes : ")
		fmt.Println(Tier1FormAsBytes)
		json.Unmarshal(Tier1FormAsBytes, &Tier1FormIndex)								//un stringify it aka JSON.parse()
		fmt.Print("Tier1FormIndex : ")
		fmt.Println(Tier1FormIndex)
		jsonResp = "{\"FormList\":\"["
		for i,val := range Tier1FormIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tier1 Forms")
			jsonResp = jsonResp + val
			if i < len(Tier1FormIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
		
		fmt.Println("len(Tier1FormIndex) : ")
		fmt.Println(len(Tier1FormIndex))
		jsonResp = jsonResp +"]\""+ "}"
		fmt.Println([]byte(jsonResp))
		fmt.Println("Fetched All Tier1 Formssuccessfully.")
		return []byte(jsonResp), nil	
	}
	if args[0]=="OEM"{
		var OemFormIndex []string
		OemFormAsBytes, err := stub.GetState(OEMFormIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get Oem Form index")
		}
		fmt.Print("OemFormAsBytes : ")
		fmt.Println(OemFormAsBytes)
		json.Unmarshal(OemFormAsBytes, &OemFormIndex)								//un stringify it aka JSON.parse()
		fmt.Print("OemFormIndex : ")
		fmt.Println(OemFormIndex)
		jsonResp = "{\"FormList\":\"["
		for i,val := range OemFormIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all OEM Forms")
			jsonResp = jsonResp + val
			if i < len(OemFormIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
		
		fmt.Println("len(OemFormIndex) : ")
		fmt.Println(len(OemFormIndex))
		jsonResp = jsonResp +"]\""+ "}"
		fmt.Println([]byte(jsonResp))
		fmt.Println("Fetched All OEM Formssuccessfully.")
		return []byte(jsonResp), nil	
	}
	return nil,errors.New("Cante fetch forms by tier Type fatal error")
}
	
	
	







func (t *ManageForm) get_AllFormByTier(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var jsonResp,errResp string
	fmt.Println("Fetching All Forms by Tier Type")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments,Expecting one argument")
	}
	
	if args[0]=="Tier-3"{
		var Tier3FormIndex []string
		Tier3FormAsBytes, err := stub.GetState(Tier3FormIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get Tier3 Form index")
		}
		
		fmt.Print("Tier3FormAsBytes : ")
		fmt.Println(Tier3FormAsBytes)
		json.Unmarshal(Tier3FormAsBytes, &Tier3FormIndex)								//un stringify it aka JSON.parse()
		fmt.Print("Tier3FormIndex : ")
		fmt.Println(Tier3FormIndex)
		
		
		
		jsonResp = "{"
		for i,val := range Tier3FormIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tier3 Forms")
			valueAsBytes, err := stub.GetState(val)
			if err != nil {
				errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
				return nil, errors.New(errResp)
			}
			fmt.Print("valueAsBytes : ")
			fmt.Println(valueAsBytes)
			jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
			if i < len(Tier3FormIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	fmt.Println("len(Tier3FormIndex) : ")
	fmt.Println(len(Tier3FormIndex))
	jsonResp = jsonResp + "}"
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched All Tier3 Shipments successfully.")
	return []byte(jsonResp), nil	
		
	}
	
	
	if args[0]=="Tier-2"{
		var Tier2FormIndex []string
		Tier2FormAsBytes, err := stub.GetState(Tier2FormIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get Tier2 Form index")
		}
		
		fmt.Print("Tier2FormAsBytes : ")
		fmt.Println(Tier2FormAsBytes)
		json.Unmarshal(Tier2FormAsBytes, &Tier2FormIndex)								//un stringify it aka JSON.parse()
		fmt.Print("Tier2FormIndex : ")
		fmt.Println(Tier2FormIndex)
		
		
		
		jsonResp = "{"
		for i,val := range Tier2FormIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tier2 Forms")
			valueAsBytes, err := stub.GetState(val)
			if err != nil {
				errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
				return nil, errors.New(errResp)
			}
			fmt.Print("valueAsBytes : ")
			fmt.Println(valueAsBytes)
			jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
			if i < len(Tier2FormIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	fmt.Println("len(Tier2ShipmentIndex) : ")
	fmt.Println(len(Tier2FormIndex))
	jsonResp = jsonResp + "}"
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched All Tier2 Forms successfully.")
	return []byte(jsonResp), nil	
		
	}
	
	if args[0]=="Tier-1"{
		var Tier1FormIndex []string
		Tier1FormAsBytes, err := stub.GetState(Tier1FormIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get Tier1 Form index")
		}
		
		fmt.Print("Tier1FormAsBytes : ")
		fmt.Println(Tier1FormAsBytes)
		json.Unmarshal(Tier1FormAsBytes, &Tier1FormIndex)								//un stringify it aka JSON.parse()
		fmt.Print("Tier1FormIndex : ")
		fmt.Println(Tier1FormIndex)
		
		
		
		jsonResp = "{"
		for i,val := range Tier1FormIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tier1 Forms")
			valueAsBytes, err := stub.GetState(val)
			if err != nil {
				errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
				return nil, errors.New(errResp)
			}
			fmt.Print("valueAsBytes : ")
			fmt.Println(valueAsBytes)
			jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
			if i < len(Tier1FormIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	fmt.Println("len(Tier1FormIndex) : ")
	fmt.Println(len(Tier1FormIndex))
	jsonResp = jsonResp + "}"
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched All Tier1 Forms successfully.")
	return []byte(jsonResp), nil	
		
	}
	
	
	if args[0]=="OEM"{
		var OemFormIndex []string
		OemFormAsBytes, err := stub.GetState(OEMFormIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get Oem Form index")
		}
		
		fmt.Print("OemFormAsBytes : ")
		fmt.Println(OemFormAsBytes)
		json.Unmarshal(OemFormAsBytes, &OemFormIndex)								//un stringify it aka JSON.parse()
		fmt.Print("OemFormIndex : ")
		fmt.Println(OemFormIndex)
		
		
		
		jsonResp = "{"
		for i,val := range OemFormIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Oem Forms")
			valueAsBytes, err := stub.GetState(val)
			if err != nil {
				errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
				return nil, errors.New(errResp)
			}
			fmt.Print("valueAsBytes : ")
			fmt.Println(valueAsBytes)
			jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
			if i < len(OemFormIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	fmt.Println("len(OEMFormIndex) : ")
	fmt.Println(len(OemFormIndex))
	jsonResp = jsonResp + "}"
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched All Oem Forms successfully.")
	return []byte(jsonResp), nil	
		
	}
	return nil,errors.New("Cante fetch forms by tier Type fatal error")
}
	



























// ============================================================================================================================
//  getForm_byUser - get Form details by buyer's name from chaincode state
// ============================================================================================================================
func (t *ManageForm) getForm_byUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//getForm_byUser('usertype','user')
	var jsonResp, user,userType, errResp string
	var FormIndex []string
	var valIndex Form
	fmt.Println("Fetching Form by User")
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2 argument")
	}
	// set user and user Type
	userType = args[0]
	user = args[1]
	fmt.Println("user : " + user)
	fmt.Println("userType : " + userType)
	if(userType == "Tier-3"){
		FormAsBytes, err := stub.GetState(Tier3FormIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get Tier-3 Form index string")
		}
		fmt.Print("FormAsBytes : ")
		fmt.Println(FormAsBytes)
		json.Unmarshal(FormAsBytes, &FormIndex)	
	}else if(userType == "Tier-2"){
		FormAsBytes, err := stub.GetState(Tier2FormIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get Tier-2 Form index string")
		}
		fmt.Print("FormAsBytes : ")
		fmt.Println(FormAsBytes)
		json.Unmarshal(FormAsBytes, &FormIndex)	
	}else if(userType == "Tier-1"){
		FormAsBytes, err := stub.GetState(Tier1FormIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get Tier-1 Form index string")
		}
		fmt.Print("FormAsBytes : ")
		fmt.Println(FormAsBytes)
		json.Unmarshal(FormAsBytes, &FormIndex)	
	}else if(userType == "OEM"){
		FormAsBytes, err := stub.GetState(OEMFormIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get OEM Form index string")
		}
		fmt.Print("FormAsBytes : ")
		fmt.Println(FormAsBytes)
		json.Unmarshal(FormAsBytes, &FormIndex)	
	}
	
								//un stringify it aka JSON.parse()
	fmt.Print("FormIndex : ")
	fmt.Println(FormIndex)
	fmt.Println("len(FormIndex) : ")
	fmt.Println(len(FormIndex))
	jsonResp = "{"
	for i,val := range FormIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for getForm_byUser")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return nil, errors.New(errResp)
		}
		fmt.Print("valueAsBytes : ")
		fmt.Println(valueAsBytes)
		json.Unmarshal(valueAsBytes, &valIndex)
		fmt.Print("valIndex: ")
		fmt.Print(valIndex)
		if valIndex.User == user{
			fmt.Println("User found")
			jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
			fmt.Println("jsonResp inside if")
			fmt.Println(jsonResp)
			if i < len(FormIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	}
	jsonResp = jsonResp + "}"
	fmt.Println("jsonResp : " + jsonResp)
	fmt.Print("jsonResp in bytes : ")
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched all forms by User")
	return []byte(jsonResp), nil											//send it onward
}

// ============================================================================================================================
//  get_AllOEMForm- get details of all OEM Form from chaincode state
// ============================================================================================================================
func (t *ManageForm) get_AllOEMForm(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var jsonOEMResp,errResp string
	var OEMFormIndex []string
	fmt.Println("Fetching All OEM Forms")
	var err error
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting single space as an argument")
	}
	// fetching all OEM forms
	OEMFormAsBytes, err := stub.GetState(OEMFormIndexStr)
	if err != nil {
		return nil, errors.New("Failed to get OEM Form index")
	}
	fmt.Print("OEMFormAsBytes : ")
	fmt.Println(OEMFormAsBytes)
	json.Unmarshal(OEMFormAsBytes, &OEMFormIndex)								//un stringify it aka JSON.parse()
	fmt.Print("OEMFormIndex : ")
	fmt.Println(OEMFormIndex)
	// OEM form data
	jsonOEMResp = "{"
	for i,val := range OEMFormIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all OEM Forms")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return nil, errors.New(errResp)
		}
		fmt.Print("valueAsBytes : ")
		fmt.Println(valueAsBytes)
		jsonOEMResp = jsonOEMResp + "\""+ val + "\":" + string(valueAsBytes[:])
		if i < len(OEMFormIndex)-1 {
			jsonOEMResp = jsonOEMResp + ","
		}
	}
	fmt.Println("len(OEMFormIndex) : ")
	fmt.Println(len(OEMFormIndex))
	jsonOEMResp = jsonOEMResp + "}"
	fmt.Println([]byte(jsonOEMResp))
	fmt.Println("Fetched All OEM Forms successfully.")
	return []byte(jsonOEMResp), nil
}

// ============================================================================================================================
//  get_AllTier1Form- get details of all Tier-1 Form from chaincode state
// ============================================================================================================================
func (t *ManageForm) get_AllTier1Form(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var jsonTier1Resp,errResp string
	var Tier1FormIndex []string
	fmt.Println("Fetching All Tier-1 Forms")
	var err error
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting single space as an argument")
	}
	// fetching all tier-1 forms
	Tier1FormAsBytes, err := stub.GetState(Tier1FormIndexStr)
	if err != nil {
		return nil, errors.New("Failed to get Tier-1 Form index")
	}
	fmt.Print("Tier1FormAsBytes : ")
	fmt.Println(Tier1FormAsBytes)
	json.Unmarshal(Tier1FormAsBytes, &Tier1FormIndex)								//un stringify it aka JSON.parse()
	fmt.Print("Tier1FormIndex : ")
	fmt.Println(Tier1FormIndex)
	// Tier-1 forms data
	jsonTier1Resp = "{"
	for i,val := range Tier1FormIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tier-1 Forms")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return nil, errors.New(errResp)
		}
		fmt.Print("valueAsBytes : ")
		fmt.Println(valueAsBytes)
		jsonTier1Resp = jsonTier1Resp + "\""+ val + "\":" + string(valueAsBytes[:])
		if i < len(Tier1FormIndex)-1 {
			jsonTier1Resp = jsonTier1Resp + ","
		}
	}
	fmt.Println("len(Tier1FormIndex) : ")
	fmt.Println(len(Tier1FormIndex))
	jsonTier1Resp = jsonTier1Resp + "}"
	fmt.Println([]byte(jsonTier1Resp))
	fmt.Println("Fetched All Tier-1 Forms successfully.")
	return []byte(jsonTier1Resp), nil
}

// ============================================================================================================================
//  get_AllTier2Form- get details of all Tier-2 Form from chaincode state
// ============================================================================================================================
func (t *ManageForm) get_AllTier2Form(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var jsonTier2Resp,errResp string
	var Tier2FormIndex []string
	fmt.Println("Fetching All Tier-2 Forms")
	var err error
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting single space as an argument")
	}
	// fetching all tier-2 forms
	Tier2FormAsBytes, err := stub.GetState(Tier2FormIndexStr)
	if err != nil {
		return nil, errors.New("Failed to get Tier-2 Form index")
	}
	fmt.Print("Tier2FormAsBytes : ")
	fmt.Println(Tier2FormAsBytes)
	json.Unmarshal(Tier2FormAsBytes, &Tier2FormIndex)								//un stringify it aka JSON.parse()
	fmt.Print("Tier2FormIndex : ")
	fmt.Println(Tier2FormIndex)
	// Tier-2 forms data
	jsonTier2Resp = "{"
	for i,val := range Tier2FormIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tier-2 Forms")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return nil, errors.New(errResp)
		}
		fmt.Print("valueAsBytes : ")
		fmt.Println(valueAsBytes)
		jsonTier2Resp = jsonTier2Resp + "\""+ val + "\":" + string(valueAsBytes[:])
		if i < len(Tier2FormIndex)-1 {
			jsonTier2Resp = jsonTier2Resp + ","
		}
	}
	fmt.Println("len(Tier2FormIndex) : ")
	fmt.Println(len(Tier2FormIndex))
	jsonTier2Resp = jsonTier2Resp + "}"
	fmt.Println([]byte(jsonTier2Resp))
	fmt.Println("Fetched All Tier-2 Forms successfully.")
	return []byte(jsonTier2Resp), nil

}

// ============================================================================================================================
//  get_AllTier3Form- get details of all Tier-3 Form from chaincode state
// ============================================================================================================================
func (t *ManageForm) get_AllTier3Form(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var jsonTier3Resp,errResp string
	var Tier3FormIndex []string
	fmt.Println("Fetching All Tier-3 Forms")
	var err error
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting single space as an argument")
	}
	// fetching all tier-3 forms
	Tier3FormAsBytes, err := stub.GetState(Tier3FormIndexStr)
	if err != nil {
		return nil, errors.New("Failed to get Tier-3 Form index")
	}
	fmt.Print("Tier3FormAsBytes : ")
	fmt.Println(Tier3FormAsBytes)
	json.Unmarshal(Tier3FormAsBytes, &Tier3FormIndex)								//un stringify it aka JSON.parse()
	fmt.Print("Tier3FormIndex : ")
	fmt.Println(Tier3FormIndex)
	// Tier-3 forms data
	jsonTier3Resp = "{"
	for i,val := range Tier3FormIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tier-3 Forms")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return nil, errors.New(errResp)
		}
		fmt.Print("valueAsBytes : ")
		fmt.Println(valueAsBytes)
		jsonTier3Resp = jsonTier3Resp + "\""+ val + "\":" + string(valueAsBytes[:])
		if i < len(Tier3FormIndex)-1 {
			jsonTier3Resp = jsonTier3Resp + ","
		}
	}
	fmt.Println("len(Tier3FormIndex) : ")
	fmt.Println(len(Tier3FormIndex))

	jsonTier3Resp = jsonTier3Resp + "}"
	fmt.Println([]byte(jsonTier3Resp))
	fmt.Println("Fetched All Tier-3 Forms successfully.")
	return []byte(jsonTier3Resp), nil
}

// ============================================================================================================================
//  get_AllForm- get details of all Form from chaincode state
// ============================================================================================================================
func (t *ManageForm) get_AllForm(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//get_AllForm(" ")
	var jsonResp string
	fmt.Println("Fetching All Forms")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting single space as an argument")
	}
	jsonOEMResp,err := t.get_AllOEMForm(stub,args);
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get all OEM forms\"}"
		return nil, errors.New(jsonResp)
	}
	jsonTier3Resp,err := t.get_AllTier3Form(stub,args);
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get all Tier-3 forms\"}"
		return nil, errors.New(jsonResp)
	}
	jsonTier2Resp,err := t.get_AllTier2Form(stub,args);
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get all Tier-2 forms\"}"
		return nil, errors.New(jsonResp)
	}
	jsonTier1Resp,err := t.get_AllTier1Form(stub,args);	
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get all Tier-1 forms\"}"
		return nil, errors.New(jsonResp)
	}
	
	jsonResp = string(jsonTier3Resp) + "," + string(jsonTier2Resp) + "," + string(jsonTier1Resp) + "," + string(jsonOEMResp);
	fmt.Println("jsonResp : " + jsonResp)
	fmt.Print("jsonResp in bytes : ")
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched All Forms successfully.")
	return []byte(jsonResp), nil
											//send it onward
}

// ============================================================================================================================
// Write - update Form into chaincode state
// ============================================================================================================================
func (t *ManageForm) update_Form(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//update_Form('faa_formNumber','quantity','userType')
	var jsonResp string
	var err error
	fmt.Println("Updating Form quantity")
	if len(args) != 3{
		return nil, errors.New("Incorrect number of arguments. Expecting 3.")
	}
	// set FAA_formNumber
	FAA_formNumber := args[0]
	quantity := args[1]
	userType := args[2]
	fmt.Print("quantity")
	fmt.Println(quantity);
	FormAsBytes, err := stub.GetState(FAA_formNumber)									//get the Form for the specified FormId from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + FAA_formNumber + "\"}"
		return nil, errors.New(jsonResp)
	}
	fmt.Print("FormAsBytes in update Form")
	fmt.Println(FormAsBytes);
	res := Form{}
	json.Unmarshal(FormAsBytes, &res)
	if res.FAA_FormNumber == FAA_formNumber{
		fmt.Println("Form found with FAA_formNumber : " + FAA_formNumber)
		fmt.Println(res);
		qty,err := strconv.Atoi(quantity)
		if err != nil {
			return nil, errors.New("Error while converting string 'quantity' to int ")
		}
		approvedQty,err := strconv.Atoi(res.Total_approvedQty)
		if err != nil {
			return nil, errors.New("Error while converting string 'approvedQty' to int ")
		}
		if(qty > approvedQty){
			return nil,errors.New("Quantity should be less than Total Approved Quantity")
		}
		res.Quantity = quantity
	}
	var forms string
	if userType == "Tier-2" {
		forms = `"tier3_Form_number": "` + res.Tier3_Form_number + `" , `
	}else if userType == "Tier-1"{
	  forms = `"tier3_Form_number": "` + res.Tier3_Form_number + `" , `+ `"tier2_Form_number": "` + res.Tier2_Form_number + `" , `
	}else if userType == "OEM"{
	  forms = `"tier3_Form_number": "` + res.Tier3_Form_number + `" , `+ `"tier2_Form_number": "` + res.Tier2_Form_number + `" , ` + `"tier1_Form_number": "` + res.Tier1_Form_number + `" , `
	}
	//build the Form json string manually
	input := 	`{`+
		`"FAA_formNumber": "` + res.FAA_FormNumber + `" , `+
		`"quantity": "` + res.Quantity + `" , `+ 
		`"FAA_formUrl": "` + res.FAA_FormURL + `" , `+ 
		`"fileHash": "`+res.File_hash +`" , `+
		`"user": "` + res.User + `" , `+
		`"itemType": "` + res.ItemType + `" , `+
		`"part_number": "` + res.Part_number + `" , `+ 
		`"total_approvedQty": "` + res.Total_approvedQty + `" , `+ 
		`"approvalDate": "` + res.ApprovalDate + `" , `+ 	
		`"authorization_number": "` + res.Authorization_number + `" , `+ 
		forms +
		`"userType": "` + res.UserType + `"`+
		`}`
	
	err = stub.PutState(FAA_formNumber, []byte(input))									//store Form with id as key
	if err != nil {
		return nil, err
	}
	fmt.Println("Form updated successfully.")
	return nil, nil
}
// ============================================================================================================================
// create Form - create a new Form for Tier-3, store into chaincode state
// ============================================================================================================================
func (t *ManageForm) createForm_Tier3(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	if len(args) != 10 {
		return nil, errors.New("Incorrect number of arguments. Expecting 9")
	}
	fmt.Println("Creating a new Form for Tier-3")
	if len(args[0]) <= 0 {
		return nil, errors.New("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return nil, errors.New("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return nil, errors.New("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return nil, errors.New("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return nil, errors.New("5th argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return nil, errors.New("6th argument must be a non-empty string")
	}
	if len(args[6]) <= 0 {
		return nil, errors.New("7th argument must be a non-empty string")
	}
	if len(args[7]) <= 0 {
		return nil, errors.New("8th argument must be a non-empty string")
	}
	if len(args[8]) <= 0 {
		return nil, errors.New("9th argument must be a non-empty string")
	}
	if len(args[9]) <= 0 {
		return nil, errors.New("10th argument must be a non-empty string")
	}
	
	FAA_formNumber := args[0] // FAA_formNumber or FAA_formNumberber
	quantity := args[1]
	FAA_formUrl := args[2]
	fileHash:=args[3]
	user := args[4]
	itemType := args[5]
	part_number := args[6]
	total_approvedQty := args[7]
	approvalDate	:= args[8]
	authorization_number := args[9]
	userType := "Tier-3"	
	qty,err := strconv.Atoi(quantity)
	if err != nil {
		return nil, errors.New("Error while converting string 'quantity' to int ")
	}
	approvedQty,err := strconv.Atoi(total_approvedQty)
	if err != nil {
		return nil, errors.New("{\"Error\":\"Error while converting string 'approvedQty' to int \"}")
	}
	if(qty > approvedQty){
		jsonResp := "Error: Quantity should be less than Total Approved Quantity"
		return nil,errors.New(jsonResp)
	}	
		
	//build the Form json string manually
	input := 	`{`+
		`"FAA_formNumber": "` + FAA_formNumber + `" , `+
		`"quantity": "` + quantity + `" , `+ 
		`"FAA_formUrl": "` + FAA_formUrl + `" , `+ 
		`"fileHash": "`+fileHash+ `" , `+ 
		`"user": "` + user + `" , `+
		`"itemType": "` + itemType + `" , `+
		`"part_number": "` + part_number + `" , `+ 
		`"total_approvedQty": "` + total_approvedQty + `" , `+ 
		`"approvalDate": "` + approvalDate + `" , `+ 	
		`"authorization_number": "` + authorization_number + `" , `+ 
		`"userType": "` + userType + `"`+
		`}`
		fmt.Println("input: " + input)
		fmt.Print("input in bytes array: ")
		fmt.Println([]byte(input))
	err = stub.PutState(FAA_formNumber, []byte(input))									//store Form with FAA_formNumber as key
	if err != nil {
		return nil, err
	}
	//get the Form index
	Tier3FormIndexAsBytes, err := stub.GetState(Tier3FormIndexStr)
	if err != nil {
		return nil, errors.New("Failed to get Tier-2 Form index")
	}
	var Tier3FormIndex []string
	fmt.Print("Tier3FormIndexAsBytes: ")
	fmt.Println(Tier3FormIndexAsBytes)
	
	json.Unmarshal(Tier3FormIndexAsBytes, &Tier3FormIndex)							//un stringify it aka JSON.parse()
	fmt.Print("Tier3FormIndex after unmarshal..before append: ")
	fmt.Println(Tier3FormIndex)
	//append
	Tier3FormIndex = append(Tier3FormIndex, FAA_formNumber)									//add Form transID to index list
	fmt.Println("! Tier-3 Form index after appending FAA_formNumber: ", Tier3FormIndex)
	jsonAsBytes, _ := json.Marshal(Tier3FormIndex)
	fmt.Print("jsonAsBytes: ")
	fmt.Println(jsonAsBytes)
	err = stub.PutState(Tier3FormIndexStr, jsonAsBytes)						//store name of Form
	if err != nil {
		return nil, err
	}

	fmt.Println("Tier-3 Form created successfully.")
	return nil, nil
}
// ============================================================================================================================
// create Form - create a new Form for Tier-2, store into chaincode state
// ============================================================================================================================
func (t *ManageForm) createForm_Tier2(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var valIndex Shipment
	if len(args) != 13 {
		return nil, errors.New("Incorrect number of arguments. Expecting 13")
	}
	fmt.Println("Creating a new Form for Tier-2")
	if len(args[0]) <= 0 {
		return nil, errors.New("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return nil, errors.New("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return nil, errors.New("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return nil, errors.New("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return nil, errors.New("5th argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return nil, errors.New("6th argument must be a non-empty string")
	}
	if len(args[6]) <= 0 {
		return nil, errors.New("7th argument must be a non-empty string")
	}
	if len(args[7]) <= 0 {
		return nil, errors.New("8th argument must be a non-empty string")
	}
	if len(args[8]) <= 0 {
		return nil, errors.New("9th argument must be a non-empty string")
	}
	if len(args[9]) <= 0 {
		return nil, errors.New("10th argument must be a non-empty string")
	}
	if len(args[10]) <= 0 {
		return nil, errors.New("11th argument must be a non-empty string")
	}
	if len(args[11]) <= 0 {
		return nil, errors.New("12th argument must be a non-empty string")
	}
	if len(args[12]) <= 0 {
		return nil, errors.New("13th argument must be a non-empty string")
	}
	
	
	FAA_formNumber := args[0]
	quantity := args[1]
	FAA_formUrl := args[2]
	fileHash:=args[3]
	user := args[4]
	itemType := args[4]
	part_number := args[6]
	total_approvedQty := args[7]
	approvalDate	:= args[8]
	authorization_number := args[9]
	tier3_Form_number := args[10]
	shipmentId := args[11]
	userType := "Tier-2"
	chaincodeURL := args[12]
	// Fetching shipment status from 'manageShipment' chaincode
	f := "getShipment_byId"
	queryArgs := util.ToChaincodeArgs(f, shipmentId)
	valueAsBytes, err := stub.QueryChaincode(chaincodeURL, queryArgs)
	if err != nil {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", err.Error())
		fmt.Printf(errStr)
		return nil, errors.New(errStr)
	} 	
	json.Unmarshal(valueAsBytes, &valIndex)
	shipmentStatus := valIndex.Status;

	// New Form for Tier-2 can only be created from received forms with “Created” status
	if(shipmentStatus != "Created"){
		fmt.Println("New Form for Tier-2 can only be created from received forms with “Created” status")
		return nil,errors.New("New Form for Tier-2 can only be created from received forms with “Created” status.")
	}

	qty,err := strconv.Atoi(quantity)
	if err != nil {
		return nil, errors.New("Error while converting string 'quantity' to int ")
	}
	approvedQty,err := strconv.Atoi(total_approvedQty)
	if err != nil {
		return nil, errors.New("Error while converting string 'approvedQty' to int ")
	}
	if(qty > approvedQty){
		return nil,errors.New("Quantity should be less than Total Approved Quantity")
	}

	FormAsBytes, err := stub.GetState(FAA_formNumber)
	if err != nil {
		return nil, errors.New("Failed to get Form FAA_formNumber")
	}
	fmt.Print("FormAsBytes: ")
	fmt.Println(FormAsBytes)
	res := Form{}
	json.Unmarshal(FormAsBytes, &res)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.FAA_FormNumber == FAA_formNumber{
		fmt.Println("This Form arleady exists: " + FAA_formNumber)
		fmt.Println(res);
		return nil, errors.New("This Form arleady exists")				//all stop a Form by this name exists
	}
	
	//build the Form json string manually
	input := 	`{`+
		`"FAA_formNumber": "` + FAA_formNumber + `" , `+
		`"quantity": "` + quantity + `" , `+ 
		`"FAA_formUrl": "` + FAA_formUrl + `" , `+ 
		`"fileHash": "`+fileHash+`" , `+
		`"user": "` + user + `" , `+
		`"itemType": "` + itemType + `" , `+
		`"part_number": "` + part_number + `" , `+ 
		`"total_approvedQty": "` + total_approvedQty + `" , `+ 
		`"approvalDate": "` + approvalDate + `" , `+ 	
		`"authorization_number": "` + authorization_number + `" , `+ 
		`"tier3_Form_number": "` + tier3_Form_number + `" , `+
		`"userType": "` + userType + `"`+  
		`}`
		fmt.Println("input: " + input)
		fmt.Print("input in bytes array: ")
		fmt.Println([]byte(input))
	err = stub.PutState(FAA_formNumber, []byte(input))									//store Form with FAA_formNumber as key
	if err != nil {
		return nil, err
	}
	//get the Form index
	Tier2FormIndexAsBytes, err := stub.GetState(Tier2FormIndexStr)
	if err != nil {
		return nil, errors.New("Failed to get Tier-2 Form index")
	}
	var Tier2FormIndex []string
	fmt.Print("Tier2FormIndexAsBytes: ")
	fmt.Println(Tier2FormIndexAsBytes)
	
	json.Unmarshal(Tier2FormIndexAsBytes, &Tier2FormIndex)							//un stringify it aka JSON.parse()
	fmt.Print("Tier2FormIndex after unmarshal..before append: ")
	fmt.Println(Tier2FormIndex)
	//append
	Tier2FormIndex = append(Tier2FormIndex, FAA_formNumber)									//add Form transID to index list
	fmt.Println("! Tier-2 Form index after appending FAA_formNumber: ", Tier2FormIndex)
	jsonAsBytes, _ := json.Marshal(Tier2FormIndex)
	fmt.Print("jsonAsBytes: ")
	fmt.Println(jsonAsBytes)
	err = stub.PutState(Tier2FormIndexStr, jsonAsBytes)						//store name of Form
	if err != nil {
		return nil, err
	}

	fmt.Println("Tier-2 Form created successfully.")

	// Update shipment status of the shipmentId from 'manageForm' chaincode
	function := "updateShipment"
	invokeArgs := util.ToChaincodeArgs(function, shipmentId)
	result, err := stub.InvokeChaincode(chaincodeURL, invokeArgs)
	if err != nil {
		errStr := fmt.Sprintf("Failed to update shipment status from 'manageForm' chaincode. Got error: %s", err.Error())
		fmt.Printf(errStr)
		return nil, errors.New(errStr)
	} 
	fmt.Sprintf("Shipment status updated successfully. Transaction hash : %s",result)
	return nil, nil
}
// ============================================================================================================================
// create Form - create a new Form for Tier-1, store into chaincode state
// ============================================================================================================================
func (t *ManageForm) createForm_Tier1(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var valIndex Shipment
	var formIndex Form
	if len(args) != 13 {
		return nil, errors.New("Incorrect number of arguments. Expecting 13")
	}
	fmt.Println("Creating a new Form for Tier-1")
	if len(args[0]) <= 0 {
		return nil, errors.New("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return nil, errors.New("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return nil, errors.New("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return nil, errors.New("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return nil, errors.New("5th argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return nil, errors.New("6th argument must be a non-empty string")
	}
	if len(args[6]) <= 0 {
		return nil, errors.New("7th argument must be a non-empty string")
	}
	if len(args[7]) <= 0 {
		return nil, errors.New("8th argument must be a non-empty string")
	}
	if len(args[8]) <= 0 {
		return nil, errors.New("9th argument must be a non-empty string")
	}
	if len(args[9]) <= 0 {
		return nil, errors.New("10th argument must be a non-empty string")
	}
	if len(args[10]) <= 0 {
		return nil, errors.New("11th argument must be a non-empty string")
	}
	if len(args[11]) <= 0 {
		return nil, errors.New("12th argument must be a non-empty string")
	}
	if len(args[12]) <= 0 {
		return nil, errors.New("13th argument must be a non-empty string")
	}
	
	
	FAA_formNumber := args[0] // FAA_formNumber or FAA_formNumberber
	quantity := args[1]
	FAA_formUrl := args[2]
	fileHash:=args[3]
	user := args[4]
	itemType := args[5]
	part_number := args[6]
	total_approvedQty := args[7]
	approvalDate	:= args[8]
	authorization_number := args[9]
	tier2_Form_number := args[10]
	shipmentId := args[11]
	userType := "Tier-1"
	chaincodeURL := args[12]
	// Fetching shipment status from 'manageShipment' chaincode
	f := "getShipment_byId"
	queryArgs := util.ToChaincodeArgs(f, shipmentId)
	valueAsBytes, err := stub.QueryChaincode(chaincodeURL, queryArgs)
	if err != nil {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", err.Error())
		fmt.Printf(errStr)
		return nil, errors.New(errStr)
	} 	
	json.Unmarshal(valueAsBytes, &valIndex)
	shipmentStatus := valIndex.Status;
	
	// New Form for Tier-1 can only be created from received forms with “Created” status
	if(shipmentStatus != "Created"){
		fmt.Println("New Form for Tier-1 can only be created from received forms with “Created” status")
		return nil,errors.New("New Form for Tier-1 can only be created from received forms with “Created” status.")
	}

	/*var formArgs []string
	formArgs[0]=tier2_Form_number */
	//  Tier-1 Form should be updated with Tier-3 Form number 

	formAsbytes,err := t.getForm_byID(stub,tier2_Form_number);
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get all OEM forms\"}"
		return nil, errors.New(jsonResp)
	}
	json.Unmarshal(formAsbytes, &formIndex)
	tier3_Form_number := formIndex.Tier3_Form_number

	qty,err := strconv.Atoi(quantity)
	if err != nil {
		return nil, errors.New("Error while converting string 'quantity' to int ")
	}
	approvedQty,err := strconv.Atoi(total_approvedQty)
	if err != nil {
		return nil, errors.New("Error while converting string 'approvedQty' to int ")
	}
	if(qty > approvedQty){
		return nil,errors.New("Quantity should be less than Total Approved Quantity")
	}
		
	FormAsBytes, err := stub.GetState(FAA_formNumber) // need to ask use FAA_formNumber or formid
	if err != nil {
		return nil, errors.New("Failed to get Form FAA_formNumber")
	}
	fmt.Print("FormAsBytes: ")
	fmt.Println(FormAsBytes)
	res := Form{}
	json.Unmarshal(FormAsBytes, &res)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.FAA_FormNumber == FAA_formNumber{
		fmt.Println("This Form arleady exists: " + FAA_formNumber)
		fmt.Println(res);
		return nil, errors.New("This Form arleady exists")				//all stop a Form by this name exists
	}
	
	//build the Form json string manually
	input := 	`{`+
		`"FAA_formNumber": "` + FAA_formNumber + `" , `+
		`"quantity": "` + quantity + `" , `+ 
		`"FAA_formUrl": "` + FAA_formUrl + `" , `+ 
		`"fileHash": "`+fileHash+`" , `+ 
		`"user": "` + user + `" , `+
		`"itemType": "` + itemType + `" , `+
		`"part_number": "` + part_number + `" , `+ 
		`"total_approvedQty": "` + total_approvedQty + `" , `+ 
		`"approvalDate": "` + approvalDate + `" , `+ 	
		`"authorization_number": "` + authorization_number + `" , `+ 
		`"tier2_Form_number": "` + tier2_Form_number + `" , `+
		`"tier3_Form_number": "` + tier3_Form_number + `" , `+
		`"userType": "` + userType + `"`+    
		`}`
		fmt.Println("input: " + input)
		fmt.Print("input in bytes array: ")
		fmt.Println([]byte(input))
	err = stub.PutState(FAA_formNumber, []byte(input))									//store Form with FAA_formNumber as key
	if err != nil {
		return nil, err
	}
	//get the Form index
	Tier1FormIndexAsBytes, err := stub.GetState(Tier1FormIndexStr)
	if err != nil {
		return nil, errors.New("Failed to get Tier-1 Form index")
	}
	var Tier1FormIndex []string
	fmt.Print("Tier1FormIndexAsBytes: ")
	fmt.Println(Tier1FormIndexAsBytes)
	
	json.Unmarshal(Tier1FormIndexAsBytes, &Tier1FormIndex)							//un stringify it aka JSON.parse()
	fmt.Print("Tier1FormIndex after unmarshal..before append: ")
	fmt.Println(Tier1FormIndex)
	//append
	Tier1FormIndex = append(Tier1FormIndex, FAA_formNumber)									//add Form transID to index list
	fmt.Println("! Tier-1 Form index after appending FAA_formNumber: ", Tier1FormIndex)
	jsonAsBytes, _ := json.Marshal(Tier1FormIndex)
	fmt.Print("jsonAsBytes: ")
	fmt.Println(jsonAsBytes)
	err = stub.PutState(Tier1FormIndexStr, jsonAsBytes)						//store name of Form
	if err != nil {
		return nil, err
	}

	fmt.Println("Tier-1 Form created successfully.")

	// Update shipment status of the shipmentId from 'manageForm' chaincode
	function := "updateShipment"
	invokeArgs := util.ToChaincodeArgs(function, shipmentId)
	result, err := stub.InvokeChaincode(chaincodeURL, invokeArgs)
	if err != nil {
		errStr := fmt.Sprintf("Failed to update shipment status from 'manageForm' chaincode. Got error: %s", err.Error())
		fmt.Printf(errStr)
		return nil, errors.New(errStr)
	} 
	fmt.Sprintf("Shipment status updated successfully. Transaction hash : %s",result)
	return nil, nil
}
// ============================================================================================================================
// create Form - create a new Form for OEM, store into chaincode state
// ============================================================================================================================
func (t *ManageForm) createForm_OEM(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var valIndex Shipment
	var	formIndex Form
	if len(args) != 13 {
		return nil, errors.New("Incorrect number of arguments. Expecting 13")
	}
	fmt.Println("Creating a new Form for OEM")
	if len(args[0]) <= 0 {
		return nil, errors.New("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return nil, errors.New("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return nil, errors.New("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return nil, errors.New("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return nil, errors.New("5th argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return nil, errors.New("6th argument must be a non-empty string")
	}
	if len(args[6]) <= 0 {
		return nil, errors.New("7th argument must be a non-empty string")
	}
	if len(args[7]) <= 0 {
		return nil, errors.New("8th argument must be a non-empty string")
	}
	if len(args[8]) <= 0 {
		return nil, errors.New("9th argument must be a non-empty string")
	}
	if len(args[9]) <= 0 {
		return nil, errors.New("10th argument must be a non-empty string")
	}
	if len(args[10]) <= 0 {
		return nil, errors.New("11th argument must be a non-empty string")
	}
	if len(args[11]) <= 0 {
		return nil, errors.New("12th argument must be a non-empty string")
	}
	if len(args[12]) <= 0 {
		return nil, errors.New("13th argument must be a non-empty string")
	}
	
	FAA_formNumber := args[0]
	quantity := args[1]
	FAA_formUrl := args[2]
	fileHash:=args[3]
	user := args[4]
	itemType := args[5]
	part_number := args[6]
	total_approvedQty := args[7]
	approvalDate	:= args[8]
	authorization_number := args[9]
	tier1_Form_number := args[10]
	shipmentId := args[11]
	userType := "OEM"
	chaincodeURL := args[12]
	// Fetching shipment status from 'manageShipment' chaincode
	f := "getShipment_byId"
	queryArgs := util.ToChaincodeArgs(f, shipmentId)
	valueAsBytes, err := stub.QueryChaincode(chaincodeURL, queryArgs)
	if err != nil {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", err.Error())
		fmt.Printf(errStr)
		return nil, errors.New(errStr)
	} 	
	json.Unmarshal(valueAsBytes, &valIndex)
	shipmentStatus := valIndex.Status;
	
	// New Form for OEM can only be created from received forms with “Created” status
	if(shipmentStatus != "Created"){
		fmt.Println("New Form for OEM can only be created from received forms with “Created” status")
		return nil,errors.New("New Form for OEM can only be created from received forms with “Created” status.")
	}
	/*var formArgs []string
	formArgs[0]=tier1_Form_number */
	//  OEM Form should be updated with tier-2 Form number 

	formAsbytes,err := t.getForm_byID(stub,tier1_Form_number);
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get form by tier1_Form_number\"}"
		return nil, errors.New(jsonResp)
	}
	json.Unmarshal(formAsbytes, &formIndex)
	tier2_Form_number := formIndex.Tier2_Form_number
	tier3_Form_number := formIndex.Tier3_Form_number


	qty,err := strconv.Atoi(quantity)
	if err != nil {
		return nil, errors.New("Error while converting string 'quantity' to int ")
	}
	approvedQty,err := strconv.Atoi(total_approvedQty)
	if err != nil {
		return nil, errors.New("Error while converting string 'approvedQty' to int ")
	}
	if(qty > approvedQty){
		return nil,errors.New("Quantity should be less than Total Approved Quantity")
	}
		
	FormAsBytes, err := stub.GetState(FAA_formNumber) // need to ask use FAA_formNumber or formid
	if err != nil {
		return nil, errors.New("Failed to get Form FAA_formNumber")
	}
	fmt.Print("FormAsBytes: ")
	fmt.Println(FormAsBytes)
	res := Form{}
	json.Unmarshal(FormAsBytes, &res)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.FAA_FormNumber == FAA_formNumber{
		fmt.Println("This Form arleady exists: " + FAA_formNumber)
		fmt.Println(res);
		return nil, errors.New("This Form arleady exists")				//all stop a Form by this name exists
	}
	
	//build the Form json string manually
	input := 	`{`+
		`"FAA_formNumber": "` + FAA_formNumber + `" , `+
		`"quantity": "` + quantity + `" , `+ 
		`"FAA_formUrl": "` + FAA_formUrl + `" , `+ 
		`"fileHash": "`+fileHash+ `" , `+
		`"user": "` + user + `" , `+
		`"itemType": "` + itemType + `" , `+
		`"part_number": "` + part_number + `" , `+ 
		`"total_approvedQty": "` + total_approvedQty + `" , `+ 
		`"approvalDate": "` + approvalDate + `" , `+ 	
		`"authorization_number": "` + authorization_number + `" , `+ 
		`"tier1_Form_number": "` + tier1_Form_number + `" , `+
		`"tier2_Form_number": "` + tier2_Form_number + `" , `+
		`"tier3_Form_number": "` + tier3_Form_number + `" , `+
		`"userType": "` + userType + `"`+
		`}`
		fmt.Println("input: " + input)
		fmt.Print("input in bytes array: ")
		fmt.Println([]byte(input))
	err = stub.PutState(FAA_formNumber, []byte(input))									//store Form with FAA_formNumber as key
	if err != nil {
		return nil, err
	}
	//get the Form index
	OEMFormIndexAsBytes, err := stub.GetState(OEMFormIndexStr)
	if err != nil {
		return nil, errors.New("Failed to get OEM Form index")
	}
	var OEMFormIndex []string
	fmt.Print("OEMFormIndexAsBytes: ")
	fmt.Println(OEMFormIndexAsBytes)
	
	json.Unmarshal(OEMFormIndexAsBytes, &OEMFormIndex)							//un stringify it aka JSON.parse()
	fmt.Print("OEMFormIndex after unmarshal..before append: ")
	fmt.Println(OEMFormIndex)
	//append
	OEMFormIndex = append(OEMFormIndex, FAA_formNumber)									//add Form transID to index list
	fmt.Println("! OEM Form index after appending FAA_formNumber: ", OEMFormIndex)
	jsonAsBytes, _ := json.Marshal(OEMFormIndex)
	fmt.Print("jsonAsBytes: ")
	fmt.Println(jsonAsBytes)
	err = stub.PutState(OEMFormIndexStr, jsonAsBytes)						//store name of Form
	if err != nil {
		return nil, err
	}

	fmt.Println("OEM Form created successfully.")
	// Update shipment status of the shipmentId from 'manageForm' chaincode
	function := "updateShipment"
	invokeArgs := util.ToChaincodeArgs(function, shipmentId)
	result, err := stub.InvokeChaincode(chaincodeURL, invokeArgs)
	if err != nil {
		errStr := fmt.Sprintf("Failed to update shipment status from 'manageForm' chaincode. Got error: %s", err.Error())
		fmt.Printf(errStr)
		return nil, errors.New(errStr)
	} 
	fmt.Sprintf("Shipment status updated successfully. Transaction hash : %s",result)
	return nil, nil
}
