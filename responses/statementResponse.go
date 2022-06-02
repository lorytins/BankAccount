package responses

import (
	"github.com/lorytins/BankAcountApi/models"
	"github.com/lorytins/BankAcountApi/utils"
)

type StatementResponse struct {
	AccountId      int                  `json:"AccountId"`
	CurrentBalance float64              `json:"CurrentBalance"`
	Statement      []OperationsResponse `json:"Statement"`
}

type OperationsResponse struct {
	Id                   int     `json:"OperationId"`
	OriginAccountId      int     `json:"OriginAccountId"`
	DestinationAccountId int     `json:"DestinationAccountId"`
	OperationType        string  `json:"OperationType"`
	Amount               float64 `json:"Amount"`
	ServiceCharge        float64 `json:"ServiceCharge"`
	OperationDate        string  `json:"OperationDate"`
}

func CreateOperationsResponse(operations []models.Operation, accountId int) []OperationsResponse {
	var operationsResponses []OperationsResponse

	for _, operation := range operations {
		var operationsResponse OperationsResponse

		if operation.OperationType == "Transfer" && operation.DestinationAccountId == accountId {
			operation.ServiceCharge = 0
		}

		operationsResponse.Id = operation.Id
		operationsResponse.OriginAccountId = operation.OriginAccountId
		operationsResponse.DestinationAccountId = operation.DestinationAccountId
		operationsResponse.OperationType = operation.OperationType
		operationsResponse.Amount = utils.ConvertToReal(operation.Amount)
		operationsResponse.ServiceCharge = utils.ConvertToReal(operation.ServiceCharge)
		operationsResponse.OperationDate = operation.OperationDate

		operationsResponses = append(operationsResponses, operationsResponse)
	}

	return operationsResponses
}
