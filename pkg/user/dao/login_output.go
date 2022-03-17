package dao

type Login_Out struct {
	Status        int     `json:"Status" binding:"required"`
	NAME          string  `json:"NAME" binding:"required"`
	USER_INFO_ID  int     `json:"-"`
	USER_LOGIN_ID string  `json:"USER_LOGIN_ID" binding:"required"`
	BALANCE       float32 `json:"BALANCE" binding:"required"`
}

type User struct {
	NAME             string `json:"NAME" binding:"required"`
	USER_INFO_ID     int    `json:"-" binding:"required"`
	USER_LOGIN_ID    string `json:"USER_ID" binding:"required"`
	MONEY_ACCOUNT_ID string `json:"-" binding:"required"`
}

type UserQuery struct {
	Status int    `json:"Status" binding:"required"`
	Users  []User `json:"Users" binding:"required"`
}
