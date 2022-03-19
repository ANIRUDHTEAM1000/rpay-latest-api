package services

import (
	accountRepo "rpay/pkg/account/repository"
)

func GetBalanceByUid(uid string) float64 {
	return accountRepo.GetBalanceByUid(uid)
}
