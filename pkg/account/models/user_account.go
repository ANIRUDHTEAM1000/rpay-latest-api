package models

import user "rpay/pkg/user/models"

type RM_USER_ACCOUNT struct {
	USER_ACCOUNT_ID int `gorm:"primaryKey;type:int;not null;autoIncrement"`

	USER_INFO_ID user.RM_USER_INFO
	RM_USER_INFO user.RM_USER_INFO `gorm:"foreignkey:USER_INFO_ID"`

	ACCOUNT_ID RM_ACCOUNT `gorm:"type:int(11);not null"`
	RM_ACCOUNT RM_ACCOUNT `gorm:"foreignkey:ACCOUNT_ID"`

	CREATED_DATE int64  `gorm:"autoCreateTime;not null"`
	CREATED_BY   string `gorm:"type:varchar(50);not null"`
	UPDATED_DATE int64  `gorm:"autoCreateTime;not null"`
	UPDATED_BY   string `gorm:"type:varchar(50);not null"`
}
