package resources

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DNS = "root:password@tcp(db:3306)/rpay"

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
