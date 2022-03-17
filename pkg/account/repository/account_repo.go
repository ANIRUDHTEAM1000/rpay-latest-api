package repository

import (
	"gorm.io/gorm"
	config "rpay/resources"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
}
