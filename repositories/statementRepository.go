package repositories

import (
	"github.com/lorytins/BankAcountApi/database"
	"github.com/lorytins/BankAcountApi/models"
)

func GetStatementByAccountId(accountId int) []models.Operation {
	var operations []models.Operation
	database.DB.Find(&operations, "destination_account_id = ? OR origin_account_id = ?", accountId, accountId)
	return operations
}
