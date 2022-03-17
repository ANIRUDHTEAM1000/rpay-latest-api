package models

type RM_USER_INFO struct {
	USER_INFO_ID        int64  `gorm:"primaryKey;type:int;not null;autoIncrement"`
	USER_LOGIN_ID       string `gorm:"type:varchar(50);not null"`
	FIRST_NAME          string `gorm:"type:varchar(100);not null"`
	MIDDLE_NAME         string `gorm:"type:varchar(100);not null;default:'' "`
	LAST_NAME           string `gorm:"type:varchar(100);not null"`
	USER_PHONE          int    `gorm:"not null"`
	USER_PHONE_VERIFIED int    `gorm:"type:tinyint(1);default:1"`
	USER_EMAIL          string `gorm:"type:varchar(100);unique_index"`
	USER_EMAIL_VERIFIED int    `gorm:"type:tinyint(1);default:1"`
	CREATED_DATE        int64  `gorm:"autoCreateTime;not null"`
	CREATED_BY          string `gorm:"type:varchar(50);not null"`
	UPDATED_DATE        int64  `gorm:"autoCreateTime;not null"`
	UPDATED_BY          string `gorm:"type:varchar(50);not null"`
}
