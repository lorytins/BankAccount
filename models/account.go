package models

import (
	"github.com/lorytins/BankAcountApi/database"
)

type Account struct {
	Id            int
	CustomerId    int
	AccountNumber string
	BranchNumber  string
	Balance       int
}

func (a *Account) Deposit(operation Operation) string{
	if operation.Amount < 0 {
		return "The deposit amount is invalid"
	} else {
		amount := operation.Amount - operation.ServiceCharge
		a.doDeposit(amount, &operation)
		database.DB.Select("OperationType", "DestinationAccountId",
			"Amount", "ServiceCharge", "OperationDate", "BalanceAccount").Create(&operation)
	}
	return ""
}

func (a *Account) Withdrawal(operation Operation) string{
	if operation.Amount > a.Balance || operation.Amount < 0{
		return "The balance is insufficient."
	} else {
		amount := operation.Amount + operation.ServiceCharge
		a.doWithdrawal(amount, &operation)
		database.DB.Select("OperationType", "OriginAccountId",
			"Amount", "ServiceCharge", "OperationDate", "BalanceAccount").Create(&operation)
	}
	return ""
}

func (a *Account) Transfer(operation Operation, destinationAccount Account) string{
	if operation.Amount > a.Balance || operation.Amount < 0 {
		return "The balance is insufficient."
	} else {
		amountWithdrawal := operation.Amount + operation.ServiceCharge
		a.doWithdrawal(amountWithdrawal, &operation)
		destinationAccount.doDeposit(operation.Amount, &operation)
		database.DB.Create(&operation)
	}
	return ""
}

func (a *Account) doWithdrawal(amount int, operation *Operation) {
	a.Balance -= amount
	database.DB.Save(&a)
}

func (a *Account) doDeposit(amount int, operation *Operation) {
	a.Balance += amount
	database.DB.Save(&a)
}
