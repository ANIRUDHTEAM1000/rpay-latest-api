package services

import (
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
