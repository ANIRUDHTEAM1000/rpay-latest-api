package repository

import (
	"fmt"
	config "rpay/resources"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetBalanceByUid(uid string) float64 {
	var balance float64
	db.Raw("select MONEY_ACCOUNT_BALANCE from RM_ACCOUNT WHERE ACCOUNT_ID = (select ACCOUNT_ID from RM_USER_ACCOUNT WHERE USER_INFO_ID = (SELECT USER_INFO_ID FROM RM_USER_INFO WHERE USER_LOGIN_ID=?));", uid).Scan(&balance)
	fmt.Println(balance)
	return balance
}
