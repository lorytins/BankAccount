package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lorytins/BankAcountApi/models"
	"github.com/lorytins/BankAcountApi/repositories"
	"github.com/lorytins/BankAcountApi/responses"
)

func NewCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		log.Printf("Unable to decode the request body %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Request fields are invalid")
	} else {
		account, message := repositories.CreateAccount(&customer)
		accountResponse := responses.CreateAccountResponse(account, customer)

		if len(message) == 0 {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(accountResponse)

		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(message)
		}
	}
}

func GetAccountByAccountNumberHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountNumber := vars["accountNumber"]
	var accountResponse responses.AccountResponse

	account, err := repositories.FindAccountByAccountNumber(accountNumber)
	if err == nil {
		customer, err := repositories.FindCustomerById(account.CustomerId)
		if err == nil {
			accountResponse = responses.CreateAccountResponse(account, customer)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(accountResponse)
		} else {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode("Customer not found")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Account not found")
	}
}
