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
		return t.get_ShipmentId_ByTier(stub, args)
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
	
	if args[0]=="Tier-3"{
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
	
	
	if args[0]=="Tier-2"{
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
	
	if args[0]=="Tier-1"{
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
	
	
	if args[0]=="OEM"{
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
	if args[0]=="Tier-3"{
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
		jsonResp = "{\"ShipmetList\":\"["
		for i,val := range Tier3ShipmentIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tier3 Shipments")
			jsonResp = jsonResp + val
			if i < len(Tier3ShipmentIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	fmt.Println("len(ShipmentIndex) : ")
	fmt.Println(len(Tier3ShipmentIndex))
	jsonResp = jsonResp +"]\""+ "}"
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched All Tier3 Shipments successfully.")
	return []byte(jsonResp), nil	
	}
	
	if args[0]=="Tier-2"{
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
		jsonResp = "{\"ShipmetList\":\"["
		for i,val := range Tier2ShipmentIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tier2 Shipments")
			jsonResp = jsonResp + val
			if i < len(Tier2ShipmentIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	fmt.Println("len(ShipmentIndex) : ")
	fmt.Println(len(Tier2ShipmentIndex))
	jsonResp = jsonResp +"]\""+ "}"
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched All Tier2 Shipments successfully.")
	return []byte(jsonResp), nil	
	}
	
	if args[0]=="Tier-1"{
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
		jsonResp = "{\"ShipmetList\":\"["
		for i,val := range Tier1ShipmentIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Tier1 Shipments")
			jsonResp = jsonResp + val
			if i < len(Tier1ShipmentIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	fmt.Println("len(ShipmentIndex) : ")
	fmt.Println(len(Tier1ShipmentIndex))
	jsonResp = jsonResp +"]\""+ "}"
	fmt.Println([]byte(jsonResp))
	fmt.Println("Fetched All Tier1 Shipments successfully.")
	return []byte(jsonResp), nil	
	}
	
	if args[0]=="OEM"{
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
		jsonResp = "{\"ShipmetList\":\"["
		for i,val := range OemShipmentIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all OEM Shipments")
			jsonResp = jsonResp + val
			if i < len(OemShipmentIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	fmt.Println("len(ShipmentIndex) : ")
	fmt.Println(len(OemShipmentIndex))
	jsonResp =jsonResp +"]\""+ "}"
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
		`"file_hash": "`+res.File_hash+`" , `+
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
 	//(....'ship_frm_country','ship_frm_city','ship_to_country','ship_to_city','truck_details'....)
 	//(....'logistics_agency_details','air/ship_way_bill_details','flight/vessel_details'....)
 	//(....'departing_port','arriving_port','scheduled_departure_date_ts','actual_arrival_date_ts')
 	//(....'vendor_name','tier_type','ipfs_hash')
 	//totan no of new arguments=25
 	
 	var err error
 	var valIndex Form
 	if len(args) != 25 {
 		return nil, errors.New("Incorrect number of arguments. Expecting 24")
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
 	if len(args[10]) <= 0 {
 		return nil, errors.New("11th argument must be a non-empty string")
 	}
 	if len(args[11]) <= 0 {
 		return nil, errors.New("12th argument must be a non-empty string")
 	}
 	if len(args[12]) <= 0 {
 		return nil, errors.New("13th argument must be a non-empty string")
 	}
 	if len(args[13]) <= 0 {
 		return nil, errors.New("14th argument must be a non-empty string")
 	}
 	if len(args[14]) <= 0 {
 		return nil, errors.New("15th argument must be a non-empty string")
 	}
 	if len(args[15]) <= 0 {
 		return nil, errors.New("16th argument must be a non-empty string")
 	}
 	if len(args[16]) <= 0 {
 		return nil, errors.New("17th argument must be a non-empty string")
 	}
 	if len(args[17]) <= 0 {
 		return nil, errors.New("18th argument must be a non-empty string")
 	}
 	if len(args[18]) <= 0 {
 		return nil, errors.New("19th argument must be a non-empty string")
 	}
 	if len(args[19]) <= 0 {
 		return nil, errors.New("20th argument must be a non-empty string")
 	}
 	if len(args[20]) <= 0 {
 		return nil, errors.New("21th argument must be a non-empty string")
 	}
 	if len(args[21]) <= 0 {
 		return nil, errors.New("22th argument must be a non-empty string")
 	}
 	if len(args[22]) <= 0 {
 		return nil, errors.New("23th argument must be a non-empty string")
 	}
 	if len(args[23]) <= 0 {
 		return nil, errors.New("24th argument must be a non-empty string")
 	}
 	if len(args[24]) <= 0 {
 		return nil, errors.New("25th argument must be a non-empty string")
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
 	ship_frm_country:=args[10]
 	ship_frm_city:=args[11]
 	ship_to_country:=args[12]
 	ship_to_city:=args[13]
 	truck_details:=args[14]
 	logistics_agency_details:=args[15]
 	air_ship_way_bill_details:=args[16]
 	flight_vessel_details:=args[17]
 	departing_port:=args[18]
 	arriving_port:=args[19]
 	scheduled_departure_date_ts:=args[20]
 	actual_arrival_date_ts:=args[21]
 	vendor_name:=args[22]
 	tier_type:=args[23]
 	file_hash:=args[24]
 	
 	
 	
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
 		`"status": "` + status + `" , `+ 
 		`"ship_frm_country": "`+ship_frm_country+ `" , `+ 
 		`"ship_frm_city": "`+ship_frm_city+ `" , `+
 		`"ship_to_country": "`+ship_to_country+ `" , `+
 		`"ship_to_city": "`+ship_to_city+ `" , `+
 		`"truck_details": "`+truck_details+`" , `+
 		`"logistics_agency_details": "`+logistics_agency_details+`" , `+
 		`"air_ship_way_bill_details": "`+air_ship_way_bill_details+`" , `+
 		`"flight_vessel_details": "`+flight_vessel_details+`" , `+
 		`"departing_port": "`+departing_port+`" , `+
 		`"arriving_port": "`+arriving_port+`" , `+
 		`"scheduled_departure_date_ts": "`+scheduled_departure_date_ts+`" , `+
 		`"actual_arrival_date_ts": "`+actual_arrival_date_ts+`" , `+
 		`"vendor_name": "`+vendor_name+`" , `+
 		`"file_hash": "`+file_hash+`" , `+
 		`"tier_type": "`+tier_type+`" `+
 
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
 	
 	
 	
 	
 	//get the Shipment index
 	if tier_type=="Tier-3"{
 		
 		Tier3ShipmentIndexAsBytes, err := stub.GetState(Tier3ShipmentIndexStr)
 		if err != nil {
 			return nil, errors.New("Failed to get Shipment index")
 		}
 		var Tier3ShipmentIndex []string
 		fmt.Print("Tier3ShipmentIndexAsBytes: ")
 		fmt.Println(Tier3ShipmentIndexAsBytes)
 
 		json.Unmarshal(Tier3ShipmentIndexAsBytes, &Tier3ShipmentIndex)							//un stringify it aka JSON.parse()
 		fmt.Print("Tier3ShipmentIndex after unmarshal..before append: ")
 		fmt.Println(Tier3ShipmentIndex)
 		//append
 		Tier3ShipmentIndex = append(Tier3ShipmentIndex, shipmentId)									//add Shipment transID to index list
 		fmt.Println("!Tier3 Shipment index after appending shipmentId: ", Tier3ShipmentIndex)
 		jsonAsBytes, _ := json.Marshal(Tier3ShipmentIndex)
 		fmt.Print("jsonAsBytes: ")
 		fmt.Println(jsonAsBytes)
 		err = stub.PutState(Tier3ShipmentIndexStr, jsonAsBytes)						//store name of Shipment
 		if err != nil {
 			return nil, err
 		}
 	}
 	
 	if tier_type=="Tier-2"{
 		
 		Tier2ShipmentIndexAsBytes, err := stub.GetState(Tier2ShipmentIndexStr)
 		if err != nil {
 			return nil, errors.New("Failed to get Tier2 Shipment index")
 		}
 		var Tier2ShipmentIndex []string
 		fmt.Print("Tier2ShipmentIndexAsBytes: ")
 		fmt.Println(Tier2ShipmentIndexAsBytes)
 
 		json.Unmarshal(Tier2ShipmentIndexAsBytes, &Tier2ShipmentIndex)							//un stringify it aka JSON.parse()
 		fmt.Print("Tier2ShipmentIndex after unmarshal..before append: ")
 		fmt.Println(Tier2ShipmentIndex)
 		//append
 		Tier2ShipmentIndex = append(Tier2ShipmentIndex, shipmentId)									//add Shipment transID to index list
 		fmt.Println("!Tier2 Shipment index after appending shipmentId: ", Tier2ShipmentIndex)
 		jsonAsBytes, _ := json.Marshal(Tier2ShipmentIndex)
 		fmt.Print("jsonAsBytes: ")
 		fmt.Println(jsonAsBytes)
 		err = stub.PutState(Tier2ShipmentIndexStr, jsonAsBytes)						//store name of Shipment
 		if err != nil {
 			return nil, err
 		}
 	}
 	
 	if tier_type=="Tier-1"{
 		
 		Tier1ShipmentIndexAsBytes, err := stub.GetState(Tier1ShipmentIndexStr)
 		if err != nil {
 			return nil, errors.New("Failed to get Tier1 Shipment index")
 		}
 		var Tier1ShipmentIndex []string
 		fmt.Print("Tier1ShipmentIndexAsBytes: ")
 		fmt.Println(Tier1ShipmentIndexAsBytes)
 
 		json.Unmarshal(Tier1ShipmentIndexAsBytes, &Tier1ShipmentIndex)							//un stringify it aka JSON.parse()
 		fmt.Print("Tier1ShipmentIndex after unmarshal..before append: ")
 		fmt.Println(Tier1ShipmentIndex)
 		//append
 		Tier1ShipmentIndex = append(Tier1ShipmentIndex, shipmentId)									//add Shipment transID to index list
 		fmt.Println("!Tier1 Shipment index after appending shipmentId: ", Tier1ShipmentIndex)
 		jsonAsBytes, _ := json.Marshal(Tier1ShipmentIndex)
 		fmt.Print("jsonAsBytes: ")
 		fmt.Println(jsonAsBytes)
 		err = stub.PutState(Tier1ShipmentIndexStr, jsonAsBytes)						//store name of Shipment
 		if err != nil {
 			return nil, err
 		}
 	}
 	
 	if tier_type=="OEM"{
 		
 		OemShipmentIndexAsBytes, err := stub.GetState(OemShipmentIndexStr)
 		if err != nil {
 			return nil, errors.New("Failed to get OEM Shipment index")
 		}
 		var OemShipmentIndex []string
 		fmt.Print("OemShipmentIndexAsBytes: ")
 		fmt.Println(OemShipmentIndexAsBytes)
 
 		json.Unmarshal(OemShipmentIndexAsBytes, &OemShipmentIndex)							//un stringify it aka JSON.parse()
 		fmt.Print("OemShipmentIndex after unmarshal..before append: ")
 		fmt.Println(OemShipmentIndex)
 		//append
 		OemShipmentIndex = append(OemShipmentIndex, shipmentId)									//add Shipment transID to index list
 		fmt.Println("!Oem Shipment index after appending shipmentId: ", OemShipmentIndex)
 		jsonAsBytes, _ := json.Marshal(OemShipmentIndex)
 		fmt.Print("jsonAsBytes: ")
 		fmt.Println(jsonAsBytes)
 		err = stub.PutState(OemShipmentIndexStr, jsonAsBytes)						//store name of Shipment
 		if err != nil {
 			return nil, err
 		}
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
