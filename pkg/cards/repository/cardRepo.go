package repository

import (
	cardModel "rpay/pkg/cards/models"
	config "rpay/resources"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetCards(uid string) []cardModel.Account {
	cards := []cardModel.Account{}
	db.Raw("SELECT MONEY_ACCOUNT_BALANCE,MONEY_ACCOUNT_ID,SECURITY_CODE,IS_ACTIVE,CREATED_DATE  as ExpireOn,ACCOUNT_TYPE_ID as Type from RM_ACCOUNT where ACCOUNT_ID in ( select ACCOUNT_ID from RM_USER_ACCOUNT where USER_INFO_ID = ( select USER_INFO_ID from RM_USER_INFO where USER_LOGIN_ID=? ) ) and ACCOUNT_TYPE_ID not in (1,2);", uid).Scan(&cards)
	var card_holder string
	db.Raw("SELECT CONCAT(FIRST_NAME,' ',LAST_NAME) from RM_USER_INFO where USER_LOGIN_ID = ?;", uid).Scan(&card_holder)
	for i := 0; i < len(cards); i++ {
		cards[i].CardHolder = card_holder
		db.Raw("SELECT ACCOUNT_TYPE_DESC from RM_ACCOUNT_TYPE where ACCOUNT_TYPE_ID=?", cards[i].Type).Scan(&cards[i].Type)
	}
	return cards
}
