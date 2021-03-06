package dao

// cardModel "rpay/pkg/cards/models"

type Login_Out struct {
	Status         int     `json:"status" binding:"required"`
	NAME           string  `json:"name" binding:"required"`
	USER_INFO_ID   int     `json:"-"`
	USER_LOGIN_ID  string  `json:"user_login_id" binding:"required"`
	BALANCE        float32 `json:"balance" binding:"required"`
	CASH_BACK      float32 `json:"cash_back" `
	RAKUTEN_POINTS int32   `json:"rakuten_points" `
}

type User struct {
	NAME             string `json:"name" binding:"required"`
	USER_INFO_ID     int    `json:"-" binding:"required"`
	USER_LOGIN_ID    string `json:"user_login_id" binding:"required"`
	MONEY_ACCOUNT_ID string `json:"-" binding:"required"`
}

type UserQuery struct {
	Status int    `json:"status" binding:"required"`
	Users  []User `json:"users" binding:"required"`
}
