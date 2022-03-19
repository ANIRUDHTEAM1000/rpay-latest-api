package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	config "rpay/resources"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
}

func GenUUID() string {
	id := uuid.New()
	return id.String()
}

func StartTransaction(sender int, receiver int, amount int64) string {
	uuid := GenUUID()
	fmt.Println("senders a_id", sender, "receiver a_id", receiver, uuid)
	var tid int

	db.Debug().Exec("INSERT INTO RT_TRANSACTION(TRANSACTION_UNIQUE_ID,TRANSACTION_TYPE_CODE,TRANSACTION_AMOUNT,CREATED_BY,UPDATED_BY)"+
		" values(?,'b',?,'Admin','Admin');", uuid, amount)
	db.Debug().Raw("select TRANSACTION_ID from RT_TRANSACTION where TRANSACTION_UNIQUE_ID = ?", uuid).Scan(&tid)
	fmt.Println(tid)
	db.Debug().Exec("INSERT INTO RT_TRANSACTION_LEDGER(TRANSACTION_ID,ACCOUNT_ID,LEDGER_TYPE_CODE,LEDGER_TRANSACTION_AMOUNT,CREATED_BY"+
		",UPDATED_BY) values (?,?,'D',?,'Admin','Admin');", tid, sender, -amount)
	db.Debug().Exec("INSERT INTO RT_TRANSACTION_LEDGER(TRANSACTION_ID,ACCOUNT_ID,LEDGER_TYPE_CODE,LEDGER_TRANSACTION_AMOUNT,CREATED_BY"+
		",UPDATED_BY) values (?,?,'C',?,'Admin','Admin');", tid, receiver, +amount)
	db.Debug().Exec("UPDATE rm_account set MONEY_ACCOUNT_BALANCE = MONEY_ACCOUNT_BALANCE-? where ACCOUNT_ID=? ;", amount, sender)
	db.Debug().Exec("UPDATE rm_account set MONEY_ACCOUNT_BALANCE = MONEY_ACCOUNT_BALANCE+? where ACCOUNT_ID=? ;", amount, receiver)

	return "success"
}

//senderAcc := userRepo.GetUserAccountByLogId(sender)
//receiverAcc := userRepo.GetUserAccountByLogId(receiver)
//db.Debug().Raw("INSERT INTO rt_transaction values(1,'XYZ','T',CURRENT_TIMESTAMP(),?,CURRENT_TIMESTAMP(),'Admin',CURRENT_TIMESTAMP(),'Admin');", amount)
//db.Debug().Raw("INSERT INTO rt_transaction_ledger values(1,1,?,'D',?,CURRENT_TIMESTAMP(),'Admin',CURRENT_TIMESTAMP(),'Admin');", senderAcc, -amount)
//db.Debug().Raw("INSERT INTO rt_transaction_ledger values(2,1,?,'C',?);", receiverAcc, +amount)
//db.Debug().Raw("UPDATE rm_account set money_account_balance = money_account_balance+? where money_account_id=?;", amount, receiverAcc)
//db.Debug().Raw("UPDATE rm_account set money_account_balance = money_account_balance-? where money_account_id=?;", amount, senderAcc)
//return "success"
