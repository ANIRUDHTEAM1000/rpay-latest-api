package services

import (
	models "rpay/pkg/rakuten_points/models"
	rpRepo "rpay/pkg/rakuten_points/repository"
)

func GetRakutenPoints(uid string) float32 {
	return rpRepo.GetRakutenPoints(uid)
}

func RedeemRP(uid string, amount float32) models.RedeemOutput {
	return rpRepo.RedeemCB(uid, amount)
}
