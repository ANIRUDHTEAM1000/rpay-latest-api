package repository

import (
	"fmt"
	models "rpay/pkg/cash_back/models"
	transactionRepo "rpay/pkg/transaction/repository"
	config "rpay/resources"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetCashBack(uid string) float32 {
	var cash_back float32
	db.Raw("SELECT MONEY_ACCOUNT_BALANCE from RM_ACCOUNT WHERE ACCOUNT_ID in (select ACCOUNT_ID from RM_USER_ACCOUNT WHERE USER_INFO_ID = (SELECT USER_INFO_ID FROM RM_USER_INFO WHERE USER_LOGIN_ID=?)) and ACCOUNT_TYPE_ID=2; ", uid).Scan(&cash_back)
	return cash_back
}

func RedeemCB(uid string, amount float32) models.RedeemOutput {

	result := models.RedeemOutput{}
	var tid string
	var senderAccountId int32
	var receiverAccountId int32
	uuid := transactionRepo.GenUUID()

	cash_back := GetCashBack(uid)
	if amount > cash_back {
		result.Status = 2
		return result
	}

	txn := db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Debug().Exec("INSERT INTO RT_TRANSACTION(TRANSACTION_UNIQUE_ID,TRANSACTION_TYPE_CODE,TRANSACTION_AMOUNT,CREATED_BY,UPDATED_BY) values(?,'c',?,'Admin','Admin');", uuid, amount).Error; err != nil {
			fmt.Println(err)
			return err
		}

		if err := tx.Debug().Raw("select TRANSACTION_ID from RT_TRANSACTION where TRANSACTION_UNIQUE_ID = ?;", uuid).Scan(&tid).Error; err != nil {
			fmt.Println(err)
			return err
		}

		if err := tx.Debug().Raw("select ACCOUNT_ID from RM_ACCOUNT where ACCOUNT_ID in (select ACCOUNT_ID from RM_USER_ACCOUNT where USER_INFO_ID = (select USER_INFO_ID from RM_USER_INFO where USER_LOGIN_ID=?)) and ACCOUNT_TYPE_ID='2';", uid).Scan(&senderAccountId).Error; err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Print(tid)
		if err := tx.Debug().Raw("select ACCOUNT_ID from RM_ACCOUNT where ACCOUNT_ID in (select ACCOUNT_ID from RM_USER_ACCOUNT where USER_INFO_ID = (select USER_INFO_ID from RM_USER_INFO where USER_LOGIN_ID=?)) and ACCOUNT_TYPE_ID='0';", uid).Scan(&receiverAccountId).Error; err != nil {
			fmt.Println(err)
			return err
		}

		if err := tx.Debug().Exec("INSERT INTO RT_TRANSACTION_LEDGER(TRANSACTION_ID,ACCOUNT_ID,LEDGER_TYPE_CODE,LEDGER_TRANSACTION_AMOUNT,CREATED_BY,UPDATED_BY) values (?,?,'D',?,'Admin','Admin');", tid, senderAccountId, -amount).Error; err != nil {
			fmt.Println(err)
			return err
		}

		if err := tx.Debug().Exec("INSERT INTO RT_TRANSACTION_LEDGER(TRANSACTION_ID,ACCOUNT_ID,LEDGER_TYPE_CODE,LEDGER_TRANSACTION_AMOUNT,CREATED_BY,UPDATED_BY) values (?,?,'C',?,'Admin','Admin');", tid, receiverAccountId, amount).Error; err != nil {
			fmt.Println(err)
			return err
		}

		if err := tx.Debug().Exec("UPDATE RM_ACCOUNT set MONEY_ACCOUNT_BALANCE = MONEY_ACCOUNT_BALANCE+? where ACCOUNT_ID=? ;", amount, receiverAccountId).Error; err != nil {
			fmt.Println(err)
			return err
		}

		if err := tx.Debug().Exec("UPDATE RM_ACCOUNT set MONEY_ACCOUNT_BALANCE = MONEY_ACCOUNT_BALANCE-? where ACCOUNT_ID=? ;", amount, senderAccountId).Error; err != nil {
			fmt.Println(err)
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})
	if txn == nil {
		result.Status = 1
	} else {
		result.Status = 0
	}
	return result
}