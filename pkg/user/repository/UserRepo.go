package repository

import (
	"fmt"
	"gorm.io/gorm"
	dao "rpay/pkg/user/dao"
	config "rpay/resources"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetUserById(user_id string) dao.Login_Out {

	var result dao.Login_Out
	obj1 := db.Raw("SELECT USER_INFO_ID,USER_LOGIN_ID, CONCAT(first_name,' ',last_name) AS 'NAME' FROM rm_user_info WHERE USER_LOGIN_ID = ?;", user_id).Scan(&result)
	if obj1.Error != nil || result.NAME == "" {
		fmt.Println(obj1.Error)
		result.Status = 0
		return result
	}
	obj2 := db.Raw("select MONEY_ACCOUNT_BALANCE from rm_account WHERE ACCOUNT_ID = (select ACCOUNT_ID from rm_user_account WHERE USER_INFO_ID = ? );", result.USER_INFO_ID).Scan(&result.BALANCE)
	if obj2.Error != nil {
		fmt.Println(obj2.Error)
		result.Status = 0
		return result
	}
	result.Status = 1
	return result
}

func GetUserByEmail(query string) dao.UserQuery {
	var users []dao.User
	obj := db.Debug().Raw("SELECT USER_INFO_ID,USER_LOGIN_ID, CONCAT(first_name,' ',last_name) AS 'NAME' FROM rm_user_info WHERE user_email LIKE ? ;", "%"+query+"%").Scan(&users)
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
	obj := db.Debug().Raw("SELECT USER_INFO_ID,USER_LOGIN_ID, CONCAT(first_name,' ',last_name) AS 'NAME' FROM rm_user_info WHERE user_phone LIKE ? ;", "%"+query+"%").Scan(&users)
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
	obj := db.Debug().Raw("SELECT USER_INFO_ID,USER_LOGIN_ID, CONCAT(first_name,' ',last_name) AS 'NAME' FROM rm_user_info WHERE CONCAT(first_name,' ',last_name) LIKE ?;", "%"+query+"%").Scan(&users)
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
	db.Raw("select MONEY_ACCOUNT_ID from rm_account WHERE ACCOUNT_ID = (select ACCOUNT_ID from rm_user_account WHERE USER_INFO_ID = ?);", user_info_id).Scan(&res)
	return res
}

func GetUserAccountByLogId(user_id string) string {
	var res string
	db.Raw("select MONEY_ACCOUNT_ID from rm_account WHERE ACCOUNT_ID = (select ACCOUNT_ID from rm_user_account WHERE USER_INFO_ID = (SELECT user_info_id FROM rm_user_info WHERE USER_LOGIN_ID=?));", user_id).Scan(&res)
	return res
}

func GetUserAccountPk(user_id string) int64 {
	var res int64
	db.Raw("select ACCOUNT_ID from rm_user_account WHERE USER_INFO_ID = (SELECT user_info_id FROM rm_user_info WHERE USER_LOGIN_ID=?);", user_id).Scan(&res)
	return res
}
