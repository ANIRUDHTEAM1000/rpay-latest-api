package models

import "google.golang.org/genproto/googleapis/type/decimal"

type RM_ACCOUNT struct {
	ACCOUNT_ID            int             `gorm:"primaryKey;not null;autoIncrement"`
	MONEY_ACCOUNT_ID      string          `gorm:"type:varchar(100);not null"`
	ACCOUNT_TYPE_ID       int             `gorm:"type:int;default:0"`
	MONEY_ACCOUNT_BALANCE decimal.Decimal `gorm:"type:decimal(7,6);default:0.00"`
	CREATED_DATE          int64           `gorm:"autoCreateTime;not null"`
	CREATED_BY            string          `gorm:"type:varchar(50);not null"`
	UPDATED_DATE          int64           `gorm:"autoCreateTime;not null"`
	UPDATED_BY            string          `gorm:"type:varchar(50);not null"`
}
