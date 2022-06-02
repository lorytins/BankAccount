package responses

import (
	"github.com/lorytins/BankAcountApi/models"
	"github.com/lorytins/BankAcountApi/utils"
)

type AccountResponse struct {
	Id            int             `json:"AccountId"`
	AccountNumber string          `json:"AccountNumber"`
	Balance       float64         `json:"Balance"`
	Customer      models.Customer `json:"Holder"`
}

func CreateAccountResponse(account models.Account, customer models.Customer) AccountResponse {
	accountResponse := AccountResponse{Id: account.Id, Customer: customer, AccountNumber: account.AccountNumber,
		Balance: utils.ConvertToReal(account.Balance)}
	return accountResponse
}
