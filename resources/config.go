package resources

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DNS = "root:Team@1000@tcp(localhost:3307)/rpay"

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(mysql.Open(DNS), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

func init() {
	Connect()
}
