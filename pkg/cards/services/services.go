package services

import (
	cardRepo "rpay/pkg/cards/repository"
	cardModel "rpay/pkg/cards/models"
)

func GetCards(uid string) []cardModel.Account {
	return cardRepo.GetCards(uid)
}
