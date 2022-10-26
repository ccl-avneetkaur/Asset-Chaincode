package chaincode

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
type Organization struct {
	OrgName     string   `json:"mspID"`
	UserName    string   `json:"commonName"`
	CompanyName string   `json:"compName"`
	TimeStamp   int64    `json:"timeStamp"`
	MemberList  []string `json:"memberList"`
}

type Response struct {
	TxID    string      `json:"txId"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
