package services

import (
	"rpay/pkg/transaction/dao"
	transactionRepo "rpay/pkg/transaction/repository"
	userRepo "rpay/pkg/user/repository"
)

func StartTransaction(sender string, receiver string, amount int64) dao.Transaction_output {
	senderPk := userRepo.GetUserAccountPk(sender)
	receiverPk := userRepo.GetUserAccountPk(receiver)
	result := transactionRepo.StartTransaction(senderPk, receiverPk, amount)
	result.SenderName = sender
	result.ReceiverName = receiver
	return result
}
