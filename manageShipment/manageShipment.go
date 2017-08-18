/*/* modified
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license Shipments.  See the NOTICE file
distributed with this work for additional inShipmentation.
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

var ShipmentIndexStr = "_Shipmentindex"	//name for the key/value that will store a list of all known Shipment
var Tier3ShipmentIndexStr="_Tier3Shipmentindex"  //name for the key/value that will store a list of all known Tier3 Shipments
var Tier2ShipmentIndexStr="_Tier2Shipmentindex"  //name for the key/value that will store a list of all known Tier2 Shipments
var Tier1ShipmentIndexStr="_Tier1Shipmentindex"  //name for the key/value that will store a list of all known Tier1 Shipments
var OemShipmentIndexStr="_OemShipmentindex"      //name for the key/value that will store a list of all known OEM Shipments


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
	Ipfs_hash string `json:"ipfs_hash"`
	
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
	err = stub.PutState(Tier3ShipmentIndexStr, jsonAsBytes)
	if err != nil {
		return nil, err
	}
	err = stub.PutState(Tier2ShipmentIndexStr, jsonAsBytes)
	if err != nil {
		return nil, err
	}
	err = stub.PutState(Tier1ShipmentIndexStr, jsonAsBytes)
	if err != nil {
		return nil, err
	}
	err = stub.PutState(OemShipmentIndexStr, jsonAsBytes)
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
	}else if function == "get_AllShipmentByTier" {													//Read a Shipment by Buyer
		return t.get_AllShipmentByTier(stub, args)
	}else if function == "get_ShipmentId_ByTier" {													//Read a Shipment by Buyer
		return t.get_AllShipmentByTier(stub, args)
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



func (t *ManageShipment) get_AllShipmentByTier(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var jsonResp,errResp string
	fmt.Println("Fetching All Shipments by Tier Type")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments,Expecting one argument")
	}
	
	if args[0]=="Tier3"{
		var Tier3ShipmentIndex []string
		Tier3ShipmentAsBytes, err := stub.GetState(Tier3ShipmentIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get Tier3 Shipment index")
		}
		
		fmt.Print("Tier3ShipmentAsBytes : ")
		fmt.Println(Tier3ShipmentAsBytes)
		json.Unmarshal(Tier3ShipmentAsBytes, &Tier3ShipmentIndex)								//un stringify it aka JSON.parse()
		fmt.Print("Tier3ShipmentIndex : ")
		fmt.Println(Tier3ShipmentIndex)
		
		
		
		jsonResp = "{"
		for i,val := range Tier3ShipmentIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tier3 Shipments")
			valueAsBytes, err := stub.GetState(val)
			if err != nil {
				errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
				return nil, errors.New(errResp)
			}
			fmt.Print("valueAsBytes : ")
			fmt.Println(valueAsBytes)
			jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
			if i < len(Tier3ShipmentIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	fmt.Println("len(ShipmentIndex) : ")
	fmt.Println(len(Tier3ShipmentIndex))
	jsonResp = jsonResp + "}"
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched All Tier3 Shipments successfully.")
	return []byte(jsonResp), nil	
		
	}
	
	
	if args[0]=="Tier2"{
		var Tier2ShipmentIndex []string
		Tier2ShipmentAsBytes, err := stub.GetState(Tier2ShipmentIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get Tier2 Shipment index")
		}
		
		fmt.Print("Tier2ShipmentAsBytes : ")
		fmt.Println(Tier2ShipmentAsBytes)
		json.Unmarshal(Tier2ShipmentAsBytes, &Tier2ShipmentIndex)								//un stringify it aka JSON.parse()
		fmt.Print("Tier2ShipmentIndex : ")
		fmt.Println(Tier2ShipmentIndex)
		
		
		
		jsonResp = "{"
		for i,val := range Tier2ShipmentIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tier2 Shipments")
			valueAsBytes, err := stub.GetState(val)
			if err != nil {
				errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
				return nil, errors.New(errResp)
			}
			fmt.Print("valueAsBytes : ")
			fmt.Println(valueAsBytes)
			jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
			if i < len(Tier2ShipmentIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	fmt.Println("len(ShipmentIndex) : ")
	fmt.Println(len(Tier2ShipmentIndex))
	jsonResp = jsonResp + "}"
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched All Tier2 Shipments successfully.")
	return []byte(jsonResp), nil	
		
	}
	
	if args[0]=="Tier1"{
		var Tier1ShipmentIndex []string
		Tier1ShipmentAsBytes, err := stub.GetState(Tier1ShipmentIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get Tier1 Shipment index")
		}
		
		fmt.Print("Tier1ShipmentAsBytes : ")
		fmt.Println(Tier1ShipmentAsBytes)
		json.Unmarshal(Tier1ShipmentAsBytes, &Tier1ShipmentIndex)								//un stringify it aka JSON.parse()
		fmt.Print("Tier1ShipmentIndex : ")
		fmt.Println(Tier1ShipmentIndex)
		
		
		
		jsonResp = "{"
		for i,val := range Tier1ShipmentIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tier1 Shipments")
			valueAsBytes, err := stub.GetState(val)
			if err != nil {
				errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
				return nil, errors.New(errResp)
			}
			fmt.Print("valueAsBytes : ")
			fmt.Println(valueAsBytes)
			jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
			if i < len(Tier1ShipmentIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	fmt.Println("len(ShipmentIndex) : ")
	fmt.Println(len(Tier1ShipmentIndex))
	jsonResp = jsonResp + "}"
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched All Tier1 Shipments successfully.")
	return []byte(jsonResp), nil	
		
	}
	
	
	if args[0]=="Oem"{
		var OemShipmentIndex []string
		OemShipmentAsBytes, err := stub.GetState(OemShipmentIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get Oem Shipment index")
		}
		
		fmt.Print("OemShipmentAsBytes : ")
		fmt.Println(OemShipmentAsBytes)
		json.Unmarshal(OemShipmentAsBytes, &OemShipmentIndex)								//un stringify it aka JSON.parse()
		fmt.Print("OemShipmentIndex : ")
		fmt.Println(OemShipmentIndex)
		
		
		
		jsonResp = "{"
		for i,val := range OemShipmentIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Oem Shipments")
			valueAsBytes, err := stub.GetState(val)
			if err != nil {
				errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
				return nil, errors.New(errResp)
			}
			fmt.Print("valueAsBytes : ")
			fmt.Println(valueAsBytes)
			jsonResp = jsonResp + "\""+ val + "\":" + string(valueAsBytes[:])
			if i < len(OemShipmentIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	fmt.Println("len(ShipmentIndex) : ")
	fmt.Println(len(OemShipmentIndex))
	jsonResp = jsonResp + "}"
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched All Oem Shipments successfully.")
	return []byte(jsonResp), nil	
		
	}
	return nil,errors.New("Cante fetch forms by tier Type fatal error")
}
	


func (t *ManageShipment) get_ShipmentId_ByTier(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
	var jsonResp string
	fmt.Println("Fetching All Shipment IDS by Tier Type")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments,Expecting one argument")
	}
	if args[0]=="Tier3"{
		var Tier3ShipmentIndex []string
		Tier3ShipmentAsBytes, err := stub.GetState(Tier3ShipmentIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get Tier3 Shipment index")
		}
		
		fmt.Print("Tier3ShipmentAsBytes : ")
		fmt.Println(Tier3ShipmentAsBytes)
		json.Unmarshal(Tier3ShipmentAsBytes, &Tier3ShipmentIndex)								//un stringify it aka JSON.parse()
		fmt.Print("Tier3ShipmentIndex : ")
		fmt.Println(Tier3ShipmentIndex)
		jsonResp = "{Tier3ShipmetList:["
		for i,val := range Tier3ShipmentIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tier3 Shipments")
			jsonResp = jsonResp + val
			if i < len(Tier3ShipmentIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	fmt.Println("len(ShipmentIndex) : ")
	fmt.Println(len(Tier3ShipmentIndex))
	jsonResp = jsonResp + "}"
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched All Tier3 Shipments successfully.")
	return []byte(jsonResp), nil	
	}
	
	if args[0]=="Tier2"{
		var Tier2ShipmentIndex []string
		Tier2ShipmentAsBytes, err := stub.GetState(Tier2ShipmentIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get Tier2 Shipment index")
		}
		
		fmt.Print("Tier2ShipmentAsBytes : ")
		fmt.Println(Tier2ShipmentAsBytes)
		json.Unmarshal(Tier2ShipmentAsBytes, &Tier2ShipmentIndex)								//un stringify it aka JSON.parse()
		fmt.Print("Tier2ShipmentIndex : ")
		fmt.Println(Tier2ShipmentIndex)
		jsonResp = "{Tier2ShipmetList:["
		for i,val := range Tier2ShipmentIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tier2 Shipments")
			jsonResp = jsonResp + val
			if i < len(Tier2ShipmentIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	fmt.Println("len(ShipmentIndex) : ")
	fmt.Println(len(Tier2ShipmentIndex))
	jsonResp = jsonResp + "}"
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched All Tier2 Shipments successfully.")
	return []byte(jsonResp), nil	
	}
	
	if args[0]=="Tier1"{
		var Tier1ShipmentIndex []string
		Tier1ShipmentAsBytes, err := stub.GetState(Tier1ShipmentIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get Tier1 Shipment index")
		}
		
		fmt.Print("Tier1ShipmentAsBytes : ")
		fmt.Println(Tier1ShipmentAsBytes)
		json.Unmarshal(Tier1ShipmentAsBytes, &Tier1ShipmentIndex)								//un stringify it aka JSON.parse()
		fmt.Print("Tier1ShipmentIndex : ")
		fmt.Println(Tier1ShipmentIndex)
		jsonResp = "{Tier1ShipmetList:["
		for i,val := range Tier1ShipmentIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tier1 Shipments")
			jsonResp = jsonResp + val
			if i < len(Tier1ShipmentIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	fmt.Println("len(ShipmentIndex) : ")
	fmt.Println(len(Tier1ShipmentIndex))
	jsonResp = jsonResp + "}"
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched All Tier1 Shipments successfully.")
	return []byte(jsonResp), nil	
	}
	
	if args[0]=="Oem"{
		var OemShipmentIndex []string
		OemShipmentAsBytes, err := stub.GetState(OemShipmentIndexStr)
		if err != nil {
			return nil, errors.New("Failed to get OEM Shipment index")
		}
		
		fmt.Print("OemShipmentAsBytes : ")
		fmt.Println(OemShipmentAsBytes)
		json.Unmarshal(OemShipmentAsBytes, &OemShipmentIndex)								//un stringify it aka JSON.parse()
		fmt.Print("OemShipmentIndex : ")
		fmt.Println(OemShipmentIndex)
		jsonResp = "{OemShipmetList:["
		for i,val := range OemShipmentIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all OEM Shipments")
			jsonResp = jsonResp + val
			if i < len(OemShipmentIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	fmt.Println("len(ShipmentIndex) : ")
	fmt.Println(len(OemShipmentIndex))
	jsonResp = jsonResp + "}"
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched All Oem Shipments successfully.")
	return []byte(jsonResp), nil	
	}
	return nil,errors.New("Cante fetch forms by tier Type fatal error")
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
		`"status": "` + res.Status + `", `+
		`"ship_frm_country": "`+res.Ship_frm_country+ `" , `+ 
		`"ship_frm_city": "`+res.Ship_frm_city+ `" , `+
		`"ship_to_country": "`+res.Ship_to_country+ `" , `+
		`"ship_to_city": "`+res.Ship_to_city+ `" , `+
		`"truck_details": "`+res.Truck_details+`" , `+
		`"logistics_agency_details": "`+res.Logistics_agency_details+`" , `+
		`"air_ship_way_bill_details": "`+res.Air_ship_way_bill_details+`" , `+
		`"flight_vessel_details": "`+res.Flight_vessel_details+`" , `+
		`"departing_port": "`+res.Departing_port+`" , `+
		`"arriving_port": "`+res.Arriving_port+`" , `+
		`"scheduled_departure_date_ts": "`+res.Scheduled_departure_date_ts+`" , `+
		`"actual_arrival_date_ts": "`+res.Actual_arrival_date_ts+`" , `+
		`"vendor_name": "`+res.Vendor_name+`" , `+
		`"ipfs_hash": "`+res.Ipfs_hash+`" , `+
		`"tier_type": "`+res.Tier_type+`" `+

	
	
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
	//createShipment('shipmentId','description','sender','receiver''sender_type','receiver_type','FAA_formNumber','quantity','shipmentDate','chaincodeURL'....)
	//(....'ship_frm_country','ship_frm_city','ship_to_country','ship_to_city','
