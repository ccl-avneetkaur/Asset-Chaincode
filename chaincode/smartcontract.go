package chaincode

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// CreateCompany issues a new asset to the world state with given details.
func (s *SmartContract) CreateCompany(ctx contractapi.TransactionContextInterface, compName string) (Response, error) {
	response := Response{
		TxID:    ctx.GetStub().GetTxID(),
		Success: false,
		Message: "",
		Data:    nil,
	}
	commonName, err := getCommonName(ctx)
	if err != nil {
		return response, err
	}
	mspId, _ := ctx.GetClientIdentity().GetMSPID()
	timestamp, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return response, err
	}
	// fmt.Println("timestamp: ", timestamp)

	organization := Organization{
		OrgName:     mspId,
		UserName:    commonName,
		CompanyName: compName,
		TimeStamp:   timestamp.GetSeconds(),
		MemberList:  []string{},
	}
	// fmt.Println("Company details are: ", organization)
	compJSON, err := json.Marshal(organization)
	if err != nil {
		return response, err
	}

	fmt.Println(organization)

	return response, ctx.GetStub().PutState(compName, compJSON)
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadCompany(ctx contractapi.TransactionContextInterface, compName string) (Response, error) {
	response := Response{
		TxID:    ctx.GetStub().GetTxID(),
		Success: false,
		Message: "",
		Data:    nil,
	}
	compJSON, err := ctx.GetStub().GetState(compName)
	if err != nil {
		return response, fmt.Errorf("failed to read from world state: %v", err)
	}
	if compJSON == nil {
		return response, fmt.Errorf("the asset %s does not exist", compName)
	}

	var organization Organization
	err = json.Unmarshal(compJSON, &organization)
	if err != nil {
		return response, err
	}

	return response, err
}

func (s *SmartContract) AddMember(ctx contractapi.TransactionContextInterface, compName string, memberName string) (Response, error) {
	response := Response{
		TxID:    ctx.GetStub().GetTxID(),
		Success: false,
		Message: "",
		Data:    nil,
	}
	commonName, _ := getCommonName(ctx)
	companyDetailsAsBytes, _ := ctx.GetStub().GetState(compName)
	log.Println("Company details as bytes: ", companyDetailsAsBytes)

	var organization Organization
	err := json.Unmarshal(companyDetailsAsBytes, &organization)
	if err != nil {
		return response, err
	}
	companyDetails := string(companyDetailsAsBytes)
	log.Println("Company details are: ", companyDetails)

	if organization.UserName != commonName {
		return response, fmt.Errorf("only user can add members")

	} else {
		organization.MemberList = append(organization.MemberList, memberName)
	}

	compJSON, err := json.Marshal(organization)
	if err != nil {
		return response, err
	}

	return response, ctx.GetStub().PutState(compName, compJSON)
}

func (s *SmartContract) DisplayMembers(ctx contractapi.TransactionContextInterface, compName string) ([]string, error) {
	companyDetailsAsBytes, _ := ctx.GetStub().GetState(compName)
	log.Println("Company details as bytes: ", companyDetailsAsBytes)
	var organization Organization
	err := json.Unmarshal(companyDetailsAsBytes, &organization)
	if err != nil {
		return nil, err
	}
	return organization.MemberList, nil
}

func (s *SmartContract) LeaveCompany(ctx contractapi.TransactionContextInterface, compName string, memberName string) (Response, error) {
	response := Response{
		TxID:    ctx.GetStub().GetTxID(),
		Success: false,
		Message: "",
		Data:    nil,
	}
	companyDetailsAsBytes, _ := ctx.GetStub().GetState(compName)
	var organization Organization
	err := json.Unmarshal(companyDetailsAsBytes, &organization)
	if err != nil {
		return response, err
	}
	for index := 0; index < len(organization.MemberList); index++ {
		if organization.MemberList[index] == memberName {
			organization.MemberList = append(organization.MemberList[:index], organization.MemberList[index+1:]...)
			fmt.Println(organization.MemberList)
		}
	}
	log.Println("Member deleted: ", memberName)
	log.Println(organization.MemberList)

	compJSON, err := json.Marshal(organization)
	if err != nil {
		return response, err
	}

	return response, ctx.GetStub().PutState(compName, compJSON)
}
