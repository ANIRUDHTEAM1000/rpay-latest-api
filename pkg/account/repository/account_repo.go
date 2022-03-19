package repository

import (
	"fmt"
	"gorm.io/gorm"
	config "rpay/resources"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetBalanceByUid(uid string) float64 {
	var balance float64
	db.Raw("select MONEY_ACCOUNT_BALANCE from rm_account WHERE ACCOUNT_ID = (select ACCOUNT_ID from rm_user_account WHERE USER_INFO_ID = (SELECT user_info_id FROM rm_user_info WHERE USER_LOGIN_ID=?));", uid).Scan(&balance)
	fmt.Println(balance)
	return balance
}
