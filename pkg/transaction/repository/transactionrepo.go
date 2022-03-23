package repository

import (
	"fmt"
	dao "rpay/pkg/transaction/dao"
	config "rpay/resources"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
}

func GenUUID() string {
	id := uuid.New()
	currenttime := time.Now()
	uuid := id.String()
	ddmmyyyy := currenttime.Format("02012006") //DDMMYYYY format
	hhmmss := currenttime.Format("150405")     //HHMMSS format
	return uuid + "_" + ddmmyyyy + "_" + hhmmss
}

func StartTransaction(sender int64, receiver int64, amount int64) dao.Transaction_output {
	var result = dao.Transaction_output{}
	uuid := GenUUID()
	fmt.Println("senders a_id", sender, "receiver a_id", receiver, uuid)
	var tid int
	txn := db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Debug().Exec("INSERT INTO RT_TRANSACTION(TRANSACTION_UNIQUE_ID,TRANSACTION_TYPE_CODE,TRANSACTION_AMOUNT,CREATED_BY,UPDATED_BY)"+
			" values(?,'b',?,'Admin','Admin');", uuid, amount).Error; err != nil {
			fmt.Println(err)
			return err
		}
		if err := tx.Debug().Raw("select TRANSACTION_ID from RT_TRANSACTION where TRANSACTION_UNIQUE_ID = ?", uuid).Scan(&tid).Error; err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(tid)
		if err := tx.Debug().Exec("INSERT INTO RT_TRANSACTION_LEDGER(TRANSACTION_ID,ACCOUNT_ID,LEDGER_TYPE_CODE,LEDGER_TRANSACTION_AMOUNT,CREATED_BY"+
			",UPDATED_BY) values (?,?,'D',?,'Admin','Admin');", tid, sender, -amount).Error; err != nil {
			fmt.Println(err)
			return err
		}
		if err := tx.Debug().Exec("INSERT INTO RT_TRANSACTION_LEDGER(TRANSACTION_ID,ACCOUNT_ID,LEDGER_TYPE_CODE,LEDGER_TRANSACTION_AMOUNT,CREATED_BY"+
			",UPDATED_BY) values (?,?,'C',?,'Admin','Admin');", tid, receiver, +amount).Error; err != nil {
			fmt.Println(err)
			return err
		}
		if err := tx.Debug().Exec("UPDATE RM_ACCOUNT set MONEY_ACCOUNT_BALANCE = MONEY_ACCOUNT_BALANCE-? where ACCOUNT_ID=? ;", amount, sender).Error; err != nil {
			fmt.Println(err)
			return err
		}
		if err := tx.Debug().Exec("UPDATE RM_ACCOUNT set MONEY_ACCOUNT_BALANCE = MONEY_ACCOUNT_BALANCE+? where ACCOUNT_ID=? ;", amount, receiver).Error; err != nil {
			fmt.Println(err)
			return err
		}
		if err := tx.Debug().Raw("select current_timestamp();").Scan(&result.Time).Error; err != nil {
			fmt.Println(err, result.Time)
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})
	if txn == nil {
		result.Status = 1
		result.Tno = uuid
		result.Amount = amount
	} else {
		result.Status = 0
	}
	return result
}
