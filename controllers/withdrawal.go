package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lorytins/BankAcountApi/models"
	"github.com/lorytins/BankAcountApi/repositories"
	"github.com/lorytins/BankAcountApi/requests"
	"github.com/lorytins/BankAcountApi/utils"
)

func WithdrawalHandler(w http.ResponseWriter, r *http.Request) {
	request := new(requests.WithdrawalRequest)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("Unable to decode the request body %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Request fields are invalid")
	}

	account, err := repositories.FindAccountByAccountNumber(request.AccountNumber)

	if err == nil {
		operation := models.NewWithdrawaOperation(account.Id, utils.ConvertToCents(request.Amount))

		messageErr := account.Withdrawal(operation)
		if messageErr == "" {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode("Withdrawal made successfully")
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(messageErr)
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		message := "Account data is invalid or null"
		json.NewEncoder(w).Encode(message)
	}
}
