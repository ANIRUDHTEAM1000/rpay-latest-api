package services

import (
	transactionRepo "rpay/pkg/transaction/repository"
	userRepo "rpay/pkg/user/repository"
)

func StartTransaction(sender string, receiver string, amount int64) string {
	senderPk := userRepo.GetUserAccountPk(sender)
	receiverPk := userRepo.GetUserAccountPk(receiver)
	return transactionRepo.StartTransaction(senderPk, receiverPk, amount)
}
