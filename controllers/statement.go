package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lorytins/BankAcountApi/repositories"
	"github.com/lorytins/BankAcountApi/requests"
	"github.com/lorytins/BankAcountApi/responses"
	"github.com/lorytins/BankAcountApi/utils"
)

func StatementHandler(w http.ResponseWriter, r *http.Request) {
	var request requests.StatementRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("Unable to decode the request body %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Request fields are invalid")
	}

	account, err := repositories.FindAccountByAccountNumber(request.AccountNumber)

	if err == nil {
		operations := repositories.GetStatementByAccountId(account.Id)
		operationsResponse := responses.CreateOperationsResponse(operations, account.Id)
		statement := responses.StatementResponse{AccountId: account.Id, CurrentBalance: utils.ConvertToReal(account.Balance),
			Statement: operationsResponse}

		json.NewEncoder(w).Encode(statement)

	} else {
		w.WriteHeader(http.StatusBadRequest)
		message := "Account data is invalid or null"
		json.NewEncoder(w).Encode(message)
	}
}
