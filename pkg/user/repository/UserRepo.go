package repository

import (
	"fmt"
	dao "rpay/pkg/user/dao"
	config "rpay/resources"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetUserById(user_id string) dao.Login_Out {

	var result dao.Login_Out
	obj := db.Raw("SELECT USER_INFO_ID,USER_LOGIN_ID, CONCAT(FIRST_NAME,' ',LAST_NAME) AS 'NAME' FROM RM_USER_INFO WHERE USER_LOGIN_ID = ?;", user_id).Scan(&result)
	if obj.Error != nil || result.NAME == "" {
		fmt.Println(obj.Error)
		result.Status = 0
		return result
	}
	fmt.Print(result.USER_INFO_ID)
	obj = db.Raw("select MONEY_ACCOUNT_BALANCE from RM_ACCOUNT WHERE ACCOUNT_ID IN (select ACCOUNT_ID from RM_USER_ACCOUNT WHERE USER_INFO_ID = ? ) and ACCOUNT_TYPE_ID = 0;", result.USER_INFO_ID).Scan(&result.BALANCE)
	if obj.Error != nil {
		fmt.Println("hi")
		fmt.Println(obj.Error)
		result.Status = 0
		return result
	}
	obj = db.Raw("select MONEY_ACCOUNT_BALANCE from RM_ACCOUNT WHERE ACCOUNT_ID IN (select ACCOUNT_ID from RM_USER_ACCOUNT WHERE USER_INFO_ID = ? ) and ACCOUNT_TYPE_ID = 1;", result.USER_INFO_ID).Scan(&result.RAKUTEN_POINTS)
	if obj.Error != nil {
		fmt.Println("hi")
		fmt.Println(obj.Error)
		result.Status = 0
		return result
	}
	obj = db.Raw("select MONEY_ACCOUNT_BALANCE from RM_ACCOUNT WHERE ACCOUNT_ID IN (select ACCOUNT_ID from RM_USER_ACCOUNT WHERE USER_INFO_ID = ? ) and ACCOUNT_TYPE_ID = 2;", result.USER_INFO_ID).Scan(&result.CASH_BACK)
	if obj.Error != nil {
		fmt.Println("hi")
		fmt.Println(obj.Error)
		result.Status = 0
		return result
	}

	// update last_logged_in

	result.Status = 1
	return result
}

func GetUserByEmail(query string) dao.UserQuery {
	var users []dao.User
	obj := db.Debug().Raw("SELECT USER_INFO_ID,USER_LOGIN_ID, CONCAT(FIRST_NAME,' ',LAST_NAME) AS 'NAME' FROM RM_USER_INFO WHERE USER_EMAIL LIKE ? ;", "%"+query+"%").Scan(&users)
	if obj.Error != nil {
		fmt.Printf("Error")
		fmt.Printf(obj.Error.Error())
		return dao.UserQuery{}
	}
	for i := 0; i < len(users); i++ {
		users[i].MONEY_ACCOUNT_ID = getUserAccount(users[i].USER_INFO_ID)
	}
	var result dao.UserQuery
	result.Users = users
	if len(users) == 0 {
		result.Status = 0
	} else {
		result.Status = 1
	}
	return result
}

func GetUserByPhone(query string) dao.UserQuery {
	var users []dao.User
	obj := db.Debug().Raw("SELECT USER_INFO_ID,USER_LOGIN_ID, CONCAT(first_name,' ',last_name) AS 'NAME' FROM RM_USER_INFO WHERE USER_PHONE LIKE ? ;", "%"+query+"%").Scan(&users)
	if obj.Error != nil {
		fmt.Printf("Error")
		fmt.Printf(obj.Error.Error())
		return dao.UserQuery{}
	}
	for i := 0; i < len(users); i++ {
		users[i].MONEY_ACCOUNT_ID = getUserAccount(users[i].USER_INFO_ID)
	}
	var result dao.UserQuery
	result.Users = users
	if len(users) == 0 {
		result.Status = 0
	} else {
		result.Status = 1
	}
	return result
}

func GetUserByName(query string) dao.UserQuery {
	var users []dao.User
	obj := db.Debug().Raw("SELECT USER_INFO_ID,USER_LOGIN_ID, CONCAT(FIRST_NAME,' ',LAST_NAME) AS 'NAME' FROM RM_USER_INFO WHERE CONCAT(FIRST_NAME,' ',LAST_NAME) LIKE ?;", "%"+query+"%").Scan(&users)
	if obj.Error != nil {
		fmt.Printf("Error")
		fmt.Printf(obj.Error.Error())
		return dao.UserQuery{}
	}
	for i := 0; i < len(users); i++ {
		users[i].MONEY_ACCOUNT_ID = getUserAccount(users[i].USER_INFO_ID)
	}
	var result dao.UserQuery
	result.Users = users
	if len(users) == 0 {
		result.Status = 0
	} else {
		result.Status = 1
	}
	return result
}

// used general way
func getUserAccount(user_info_id int) string {
	var res string
	db.Raw("SELECT MONEY_ACCOUNT_ID from RM_ACCOUNT WHERE ACCOUNT_ID in (select ACCOUNT_ID from RM_USER_ACCOUNT WHERE USER_INFO_ID = ?) and ACCOUNT_TYPE_ID=0;", user_info_id).Scan(&res)
	return res
}

func GetUserAccountByLogId(user_id string) string {
	var res string
	db.Raw("select MONEY_ACCOUNT_ID from RM_ACCOUNT WHERE ACCOUNT_ID in (select ACCOUNT_ID from RM_USER_ACCOUNT WHERE USER_INFO_ID = (SELECT USER_INFO_ID FROM RM_USER_INFO WHERE USER_LOGIN_ID=?));", user_id).Scan(&res)
	return res
}

func GetUserAccountPk(user_id string) int64 {
	var res int64
	db.Raw("select ACCOUNT_ID from RM_ACCOUNT where ACCOUNT_ID in (select ACCOUNT_ID from RM_USER_ACCOUNT WHERE USER_INFO_ID = (SELECT USER_INFO_ID FROM RM_USER_INFO WHERE USER_LOGIN_ID=?)) and ACCOUNT_TYPE_ID=0;", user_id).Scan(&res)
	return res
}
