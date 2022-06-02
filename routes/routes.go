package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lorytins/BankAcountApi/controllers"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/api/customer/", controllers.NewCustomerHandler).Methods("Post")
	router.HandleFunc("/api/account/{accountNumber}", controllers.GetAccountByAccountNumberHandler).Methods("Get")
	router.HandleFunc("/api/deposit", controllers.DepositHandler).Methods("Post")
	router.HandleFunc("/api/withdrawal", controllers.WithdrawalHandler).Methods("Post")
	router.HandleFunc("/api/transfer", controllers.TransferHandler).Methods("Post")
	router.HandleFunc("/api/statement", controllers.StatementHandler).Methods("Get")

	log.Fatal(http.ListenAndServe(":8000", router))
}
