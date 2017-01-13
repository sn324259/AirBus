/*/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license Shipments.  See the NOTICE file
distributed with this work for additional inShipmentation
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

// ManageShipment example simple Chaincode implementation
type ManageShipment struct {
}

var ShipmentIndexStr = "_Shipmentindex"				//name for the key/value that will store a list of all known Shipments

type Form struct{
								// Attributes of a Form 
	FAA_FormNumber string `json:"FAA_formNumber"`	
	Quantity string `json:"quantity"`
	FAA_FormURL string `json:"FAA_formUrl"`
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
}
// ============================================================================================================================
// Main - start the chaincode for Shipment management
// ============================================================================================================================
func main() {			
	err := shim.Start(new(ManageShipment))
	if err != nil {
		fmt.Printf("Error starting Shipment management chaincode: %s", err)
	}
}
// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================
func (t *ManageShipment) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var msg string
	var err error
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	// Initialize the chaincode
	msg = args[0]
	fmt.Println("ManageShipment chaincode is deployed successfully.");
	
	// Write the state to the ledger
	err = stub.PutState("abc", []byte(msg))				//making a test var "abc", I find it handy to read/write to it right away to test the network
	if err != nil {
		return nil, err
	}
	
	var empty []string
	jsonAsBytes, _ := json.Marshal(empty)								//marshal an emtpy array of strings to clear the index
	err = stub.PutState(ShipmentIndexStr, jsonAsBytes)
	if err != nil {
		return nil, err
	}
	
	return nil, nil
}
// ============================================================================================================================
// Run - Our entry Shipmentint for Invocations - [LEGACY] obc-peer 4/25/2016
// ============================================================================================================================
func (t *ManageShipment) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("run is running " + function)
	return t.Invoke(stub, function, args)
}
// ============================================================================================================================
// Invoke - Our entry Shipmentint for Invocations
// ============================================================================================================================
func (t *ManageShipment) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	}else if function == "createShipment" {											//create a new Shipment
		return t.createShipment(stub, args)
	}else if function == "updateShipment" {									//update a Shipment
		return t.updateShipment(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)					//error
	return nil, errors.New("Received unknown function invocation")
}
// ============================================================================================================================
// Query - Our entry Shipmentint for Queries
// ============================================================================================================================
func (t *ManageShipment) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "getShipment_byReceiver" {													//Read a Shipment by ShipmentID
		return t.getShipment_byReceiver(stub, args)
	} else if function == "getShipment_bySender" {													//Read a Shipment by Buyer
		return t.getShipment_bySender(stub, args)
	} else if function == "get_AllShipment" {													//Read all Shipments
		return t.get_AllShipment(stub, args)
	}else if function == "getShipment_byId" {													//Read a Shipment by Buyer
		return t.getShipment_byId(stub, args)
	}

	fmt.Println("query did not find func: " + function)						//error
	return nil, errors.New("Received unknown function query")
}
// ============================================================================================================================
//  getShipment_byId - get Shipment details by Shipment ID from chaincode state
// ============================================================================================================================
func (t *ManageShipment) getShipment_byId(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//getForm_byID('shipmentId')
	var shipmentId, jsonResp string
	var err error
	fmt.Println("Fetching shipment Form by shipmentId")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting ID of the shipment to query")
	}
	// set shipmentId
	shipmentId = args[0]
	valAsbytes, err := stub.GetState(shipmentId)									//get the shipmentId from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + shipmentId + "\"}"
		return nil, errors.New(jsonResp)
	}
	fmt.Print("valAsbytes : ")
	fmt.Println(valAsbytes)
	fmt.Println("Fetched Form by shipmentId")
	return valAsbytes, nil	
}
// ============================================================================================================================
//  getShipment_byReceiver - get Shipment details by Receiver from chaincode state
// ============================================================================================================================
func (t *ManageShipment) getShipment_byReceiver(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//getShipment_byReceiver('receiver')
	var jsonResp, receiver, errResp string
	var ShipmentIndex []string
	var valIndex Shipment
	fmt.Println("Fetching Shipment by Receiver")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1 argument")
	}
	// set receiver
	receiver = args[0]
	fmt.Println("receiver : " + receiver)
	ShipmentAsBytes, err := stub.GetState(ShipmentIndexStr)
	if err != nil {
		return nil, errors.New("Failed to get Shipment index string")
	}
	fmt.Print("ShipmentAsBytes : ")
	fmt.Println(ShipmentAsBytes)
	json.Unmarshal(ShipmentAsBytes, &ShipmentIndex)									//un stringify it aka JSON.parse()
	fmt.Print("ShipmentIndex : ")
	fmt.Println(ShipmentIndex)
	fmt.Println("len(ShipmentIndex) : ")
	fmt.Println(len(ShipmentIndex))
	jsonResp = "{"
	for i,val := range ShipmentIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for getShipment_byReceiver")
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
		if valIndex.Receiver == receiver{
			fmt.Println("Receiver found")
			jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
			fmt.Println("jsonResp inside if")
			fmt.Println(jsonResp)
			if i < len(ShipmentIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	}
	jsonResp = jsonResp + "}"
	fmt.Println("jsonResp : " + jsonResp)
	fmt.Print("jsonResp in bytes : ")
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched all Shipments by Receiver")
	return []byte(jsonResp), nil											//send it onward
}

// ============================================================================================================================
//  getShipment_bySender - get Shipment details by Sender from chaincode state
// ============================================================================================================================
func (t *ManageShipment) getShipment_bySender(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//getShipment_bySender('sender')
	var jsonResp, sender, errResp string
	var ShipmentIndex []string
	var valIndex Shipment
	fmt.Println("Fetching Shipment by Sender")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1 argument")
	}
	// set sender
	sender = args[0]
	fmt.Println("sender : " + sender)
	ShipmentAsBytes, err := stub.GetState(ShipmentIndexStr)
	if err != nil {
		return nil, errors.New("Failed to get Shipment index string")
	}
	fmt.Print("ShipmentAsBytes : ")
	fmt.Println(ShipmentAsBytes)
	json.Unmarshal(ShipmentAsBytes, &ShipmentIndex)									//un stringify it aka JSON.parse()
	fmt.Print("ShipmentIndex : ")
	fmt.Println(ShipmentIndex)
	fmt.Println("len(ShipmentIndex) : ")
	fmt.Println(len(ShipmentIndex))
	jsonResp = "{"
	for i,val := range ShipmentIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for getShipment_bySender")
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
		if valIndex.Sender == sender{
			fmt.Println("Sender found")
			jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
			fmt.Println("jsonResp inside if")
			fmt.Println(jsonResp)
			if i < len(ShipmentIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	}
	jsonResp = jsonResp + "}"
	fmt.Println("jsonResp : " + jsonResp)
	fmt.Print("jsonResp in bytes : ")
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched all Shipments by Sender")
	return []byte(jsonResp), nil											//send it onward
}

// ============================================================================================================================
//  get_AllShipment- get details of all Shipment from chaincode state
// ============================================================================================================================
func (t *ManageShipment) get_AllShipment(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var jsonResp,errResp string
	var ShipmentIndex []string
	fmt.Println("Fetching All Shipments")
	var err error
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting single space as an argument")
	}
	// fetching all Shipments
	ShipmentAsBytes, err := stub.GetState(ShipmentIndexStr)
	if err != nil {
		return nil, errors.New("Failed to get Shipment index")
	}
	fmt.Print("ShipmentAsBytes : ")
	fmt.Println(ShipmentAsBytes)
	json.Unmarshal(ShipmentAsBytes, &ShipmentIndex)								//un stringify it aka JSON.parse()
	fmt.Print("ShipmentIndex : ")
	fmt.Println(ShipmentIndex)
	// Shipment data
	jsonResp = "{"
	for i,val := range ShipmentIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Shipments")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return nil, errors.New(errResp)
		}
		fmt.Print("valueAsBytes : ")
		fmt.Println(valueAsBytes)
		jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
		if i < len(ShipmentIndex)-1 {
			jsonResp = jsonResp + ","
		}
	}
	fmt.Println("len(ShipmentIndex) : ")
	fmt.Println(len(ShipmentIndex))
	jsonResp = jsonResp + "}"
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched All Shipments successfully.")
	return []byte(jsonResp), nil
}

// ============================================================================================================================
// updateShipment - update Shipment status into chaincode state
// ============================================================================================================================
func (t *ManageShipment) updateShipment(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//updateShipment('faa_ShipmentNumber')
	var jsonResp string
	var err error
	fmt.Println("Updating Shipment status")
	if len(args) != 1{
		return nil, errors.New("Incorrect number of arguments. Expecting 1.")
	}
	// set shipmentId
	shipmentId := args[0]
	ShipmentAsBytes, err := stub.GetState(shipmentId)									//get the Shipment for the specified ShipmentId from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + shipmentId + "\"}"
		return nil, errors.New(jsonResp)
	}
	fmt.Print("ShipmentAsBytes in update Shipment")
	fmt.Println(ShipmentAsBytes);
	res := Shipment{}
	json.Unmarshal(ShipmentAsBytes, &res)
	if res.ShipmentID == shipmentId{
		fmt.Println("Shipment found with shipmentId : " + shipmentId)
		fmt.Println(res);
		res.Status = "Consumed"
	}
	
	//build the Shipment json string manually
	input := 	`{`+
		`"shipmentId": "` + res.ShipmentID + `" , `+
		`"description": "` + res.Description + `" , `+ 
		`"sender": "` + res.Sender + `" , `+
		`"senderType": "` + res.SenderType + `" , `+
		`"receiver": "` + res.Receiver + `" , `+
		`"receiverType": "` + res.ReceiverType + `" , `+
		`"FAA_formNumber": "` + res.FAA_FormNumber + `" , `+
		`"quantity": "` + res.Quantity + `" , `+ 
		`"shipmentDate": "` + res.ShipmentDate + `" , `+ 
		`"status": "` + res.Status + `"`+ 
	    `}`
		
	err = stub.PutState(shipmentId, []byte(input))									//store Shipment with id as key
	if err != nil {
		return nil, err
	}
	fmt.Println("Shipment status updated successfully.")
	return nil, nil
}
// ============================================================================================================================
// create Shipment - create a new Shipment, store into chaincode state
// ============================================================================================================================
func (t *ManageShipment) createShipment(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//createShipment('shipmentId','description','sender','receiver','FAA_formNumber','quantity','shipmentDate')
	var err error
	var valIndex Form
	if len(args) != 10 {
		return nil, errors.New("Incorrect number of arguments. Expecting 10")
	}
	fmt.Println("Creating a new Shipment")
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
	
	shipmentId := args[0]
	description := args[1]
	sender := args[2]
	senderType := args[3]
	receiver := args[4] 
	receiverType := args[5] 
	FAA_formNumber := args[6]
	quantity := args[7]
	shipmentDate := args[8]
	status := "Created"
	chaincodeURL := args[9]
	// Adding Rule for senderType and receiverType
	if(senderType == "Tier-3" && receiverType != "Tier-2"){
		return nil,errors.New("Tier-3 can send shipment to Tier-2 only")
	}else if(senderType == "Tier-2" && receiverType != "Tier-1"){
		return nil,errors.New("Tier-2 can send shipment to Tier-1 only")
	}else if(senderType == "Tier-1" && receiverType != "OEM"){
		return nil,errors.New("Tier-1 can send shipment to OEM only")
	}
	fmt.Print("senderType: ")
	fmt.Println(senderType)
	fmt.Print("receiverType: ")
	fmt.Println(receiverType)

	// calculating available quantity by fetching total approved quantity and quantity from 'manageForm' chaincode
	f := "getForm_byID"
	queryArgs := util.ToChaincodeArgs(f, FAA_formNumber)
	valueAsBytes, err := stub.QueryChaincode(chaincodeURL, queryArgs)
	if err != nil {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", err.Error())
		fmt.Printf(errStr)
		return nil, errors.New(errStr)
	} 	
	fmt.Print("valueAsBytes : ")
	fmt.Println(valueAsBytes)
	json.Unmarshal(valueAsBytes, &valIndex)
	fmt.Print("valIndex: ")
	fmt.Println(valIndex)
	
	qty,err := strconv.Atoi(quantity)
	if err != nil {
		fmt.Sprintf("Error while converting string 'quantity' to int : %s", err.Error())
		return nil, errors.New("Error while converting string 'quantity' to int ")
	}
	fmt.Print("qty: ")
	fmt.Println(qty)
	// Fetch quantity from form
	formQty,err := strconv.Atoi(valIndex.Quantity)
	if err != nil {
		return nil, errors.New("Error while converting string 'form quantity' to int ")
	}
	fmt.Print("formQty: ")
	fmt.Println(formQty)
	// Fetch Total approved quantity from form
	/*approvedQty,err := strconv.Atoi(valIndex.Total_approvedQty)
	if err != nil {
		return nil, errors.New("Error while converting string 'approvedQty' to int ")
	}*/

	//Shipped quantity cannot be greater than Form’s quantity
	if(qty > formQty){
		return nil,errors.New("Shipped quantity cannot be greater than Form’s quantity")
	}	

	// fetching shipments from chaincode
	ShipmentAsBytes, err := stub.GetState(shipmentId) 
	if err != nil {
		return nil, errors.New("Failed to get Shipment ID")
	}
	fmt.Print("ShipmentAsBytes: ")
	fmt.Println(ShipmentAsBytes)
	res := Shipment{}
	json.Unmarshal(ShipmentAsBytes, &res)
	fmt.Print("res: ")
	fmt.Println(res)
	
	// Shipments marked “Consumed” cannot be used for creating new Forms
	if res.Status == "Consumed"{
		fmt.Println("This Shipment is already consumed. New form cannot be created")
		return nil,errors.New("New form cannot be created as this Shipment is already consumed.")
	}
	
	//build the Shipment json string manually
	input := 	`{`+
		`"shipmentId": "` + shipmentId + `" , `+
		`"description": "` + description + `" , `+ 
		`"sender": "` + sender + `" , `+
		`"senderType": "` + senderType + `" , `+
		`"receiver": "` + receiver + `" , `+
		`"receiverType": "` + receiverType + `" , `+
		`"FAA_formNumber": "` + FAA_formNumber + `" , `+
		`"quantity": "` + quantity + `" , `+ 
		`"shipmentDate": "` + shipmentDate + `" , `+ 
		`"status": "` + status + `"`+ 
	    `}`
	fmt.Println("input: " + input)
	fmt.Print("input in bytes array: ")
	fmt.Println([]byte(input))
	err = stub.PutState(shipmentId, []byte(input))									//store Shipment with FAA_FormNumber as key
	if err != nil {
		return nil, err
	}
	//get the Shipment index
	ShipmentIndexAsBytes, err := stub.GetState(ShipmentIndexStr)
	if err != nil {
		return nil, errors.New("Failed to get Shipment index")
	}
	var ShipmentIndex []string
	fmt.Print("ShipmentIndexAsBytes: ")
	fmt.Println(ShipmentIndexAsBytes)
	
	json.Unmarshal(ShipmentIndexAsBytes, &ShipmentIndex)							//un stringify it aka JSON.parse()
	fmt.Print("ShipmentIndex after unmarshal..before append: ")
	fmt.Println(ShipmentIndex)
	//append
	ShipmentIndex = append(ShipmentIndex, shipmentId)									//add Shipment transID to index list
	fmt.Println("! Shipment index after appending shipmentId: ", ShipmentIndex)
	jsonAsBytes, _ := json.Marshal(ShipmentIndex)
	fmt.Print("jsonAsBytes: ")
	fmt.Println(jsonAsBytes)
	err = stub.PutState(ShipmentIndexStr, jsonAsBytes)						//store name of Shipment
	if err != nil {
		return nil, err
	}

	fmt.Println("Shipment created successfully.")
	// calculate quantity left after shipment creation
	var remainingQty int
	remainingQty = formQty - qty

	fmt.Print("remainingQty : ")
	fmt.Println(remainingQty)
	// Forms should be updated to reflect the actual quantity left after shipment
	function := "update_Form"
	invokeArgs := util.ToChaincodeArgs(function, FAA_formNumber,strconv.Itoa(remainingQty),senderType)
	valAsBytes, err := stub.InvokeChaincode(chaincodeURL, invokeArgs)
	if err != nil {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", err.Error())
		fmt.Printf(errStr)
		return nil, errors.New(errStr)
	}
	fmt.Print("valAsBytes : ")
	fmt.Println(valAsBytes)
	fmt.Printf("Form updated successfully after successful Shipment creation.")
	return nil, nil
}
