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

func TransferHandler(w http.ResponseWriter, r *http.Request) {
	request := new(requests.TransferRequest)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("Unable to decode the request body %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Request fields are invalid")
	}

	originAccount, originErr := repositories.FindAccountByAccountNumber(request.OriginAccountNumber)
	destinationAccount, destinationErr := repositories.FindAccountByAccountNumber(request.DestinationAccountNumber)

	if originErr == nil && destinationErr == nil {
		operation := models.NewTransferOperation(originAccount.Id, destinationAccount.Id, utils.ConvertToCents(request.Amount))

		messageErr := originAccount.Transfer(operation, destinationAccount)
		if messageErr == "" {
			w.WriteHeader(http.StatusCreated)
			message := "Transfer made successfully"
			json.NewEncoder(w).Encode(message)
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
