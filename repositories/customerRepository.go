package repositories

import (
	"github.com/lorytins/BankAcountApi/database"
	"github.com/lorytins/BankAcountApi/models"
)

func FindCustomerById(customerId int) (models.Customer, error) {
	var customer models.Customer
	err := database.DB.First(&customer, "id = ?", customerId)
	return customer, err.Error
}
