package chaincode

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func getCommonName(ctx contractapi.TransactionContextInterface) (string, error) {
	x509, err := ctx.GetClientIdentity().GetX509Certificate()
	if err != nil {
		return "", err
	}
	commonName := x509.Subject.CommonName
	fmt.Println("Common name is: ", commonName)
	return commonName, nil
}
