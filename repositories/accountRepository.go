package repositories

import (
	"log"
	"strconv"

	"github.com/lorytins/BankAcountApi/database"
	"github.com/lorytins/BankAcountApi/models"
	"github.com/lorytins/BankAcountApi/utils"
)

func FindAccountByAccountNumber(accountNumber string) (models.Account, error) {
	var account models.Account
	err := database.DB.First(&account, "account_number = ?", accountNumber)
	return account, err.Error
}

func CreateAccount(customer *models.Customer) (models.Account, []error) {
	var account models.Account
	var messageError []error

	err := database.DB.Create(&customer).Error
	if err != nil {
		log.Println(err)
		messageError = append(messageError, err)
	}else{
		account.AccountNumber = strconv.Itoa(utils.GenerateAccountNumber())
		account.BranchNumber = strconv.Itoa(utils.GenerateBranchNumber())
		account.CustomerId = customer.Id

		err := database.DB.Create(&account).Error
		if err != nil {
			log.Println(err)
			messageError = append(messageError, err)
		}
	}
	
	return account, messageError
}
