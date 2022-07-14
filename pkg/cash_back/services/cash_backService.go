package services

import (
	models "rpay/pkg/cash_back/models"
	cbRepo "rpay/pkg/cash_back/repository"
)

func GetCashBack(uid string) float32 {
	return cbRepo.GetCashBack(uid)
}

func RedeemCB(uid string, amount float32) models.RedeemOutput {
	return cbRepo.RedeemCB(uid, amount)
}
