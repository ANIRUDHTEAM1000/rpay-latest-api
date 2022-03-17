package models

type RM_ACCOUNT_TYPE struct {
	ACCOUNT_TYPE_ID   int    `gorm:"primaryKey;not null"`
	ACCOUNT_TYPE_DESC string `gorm:"type:varchar(100);not null"`
	CREATED_DATE      int64  `gorm:"autoCreateTime;not null"`
	CREATED_BY        string `gorm:"type:varchar(50);not null"`
	UPDATED_DATE      int64  `gorm:"autoCreateTime;not null"`
	UPDATED_BY        string `gorm:"type:varchar(50);not null"`
}
