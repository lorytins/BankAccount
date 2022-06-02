package models

import (
	"math"
	"time"

	"github.com/lorytins/BankAcountApi/utils"
)

type Operation struct {
	Id                   int
	OriginAccountId      int
	DestinationAccountId int
	OperationType        string
	Amount               int
	ServiceCharge        int
	OperationDate        string
}

func NewTransferOperation(originAccountId int, destinationAccountId int, amount int) Operation {
	return Operation{
		OriginAccountId:      originAccountId,
		DestinationAccountId: destinationAccountId,
		ServiceCharge:        100,
		Amount:               amount,
		OperationType:        "Transfer",
		OperationDate:        time.Now().Format("2006-01-02 3:4:5"),
	}
}

func NewDepositOperation(destinationAccountId int, amount int) Operation {
	percentChargeValue := utils.CalcPercent(amount, 1.0)
	serviceCharge := int(math.Ceil(percentChargeValue))
	return Operation{
		DestinationAccountId: destinationAccountId,
		ServiceCharge:        serviceCharge,
		Amount:               amount,
		OperationType:        "Deposit",
		OperationDate:        time.Now().Format("2006-01-02 3:4:5"),
	}
}

func NewWithdrawaOperation(OriginAccountId int, amount int) Operation {
	return Operation{
		OriginAccountId: OriginAccountId,
		ServiceCharge:   400,
		Amount:          amount,
		OperationType:   "Withdrawal",
		OperationDate:   time.Now().Format("2006-01-02 3:4:5"),
	}
}
