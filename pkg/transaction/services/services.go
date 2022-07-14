package services

import (
	"fmt"
	"math"
	accountRepo "rpay/pkg/account/repository"
	accountService "rpay/pkg/account/services"
	"rpay/pkg/transaction/dao"
	transactionRepo "rpay/pkg/transaction/repository"
	userRepo "rpay/pkg/user/repository"
)

func StartTransaction(sender string, receiver string, amount int64) dao.Transaction_output {
	if accountService.GetBalanceByUid(sender) >= float64(amount) {
		senderPk := userRepo.GetUserAccountPk(sender)
		receiverPk := userRepo.GetUserAccountPk(receiver)
		result := transactionRepo.StartTransaction(senderPk, receiverPk, amount)
		result.SenderName = sender
		result.ReceiverName = receiver
		return result
	} else {
		return dao.Transaction_output{0, "", "", 0, sender, receiver}
	}
}

func GetTransactionList(userId string, pageNumber int) dao.TransactionsList {
	var transactionList dao.TransactionsList
	// getting accountId using userId of user
	accountId := transactionRepo.GetAccountIdFromUserId(userId)
	fmt.Print(userId, pageNumber)
	transactions := transactionRepo.GetTransactions(accountId, pageNumber)
	transactionList.Status = 1
	for i := len(transactions) - 1; i >= 0; i-- {
		var transaction dao.Transaction
		//fetch name from account_id
		transaction.Name = accountRepo.GetNameFromAccountId(transactions[i].ACCOUNT_ID)
		// fetch transaction_number from transaction_id
		transaction.TransactionNumber = transactionRepo.GetTransactionNumberFromId(transactions[i].TRANSACTION_ID)
		if transaction.TransactionNumber == "" {
			continue
		}
		transaction.Amount = transactions[i].LEDGER_TRANSACTION_AMOUNT
		transaction.Date = transactions[i].CREATED_DATE
		transactionList.Transactions = append(transactionList.Transactions, transaction)
	}
	x := transactionRepo.GetTotalTransactions(accountId) / 10
	transactionList.TotalPages = math.Ceil(float64(x))
	return transactionList
}
