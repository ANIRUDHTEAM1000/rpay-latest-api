package models

type RT_TRANSACTION_LEDGER struct {
	TRANSACTION_LEDGER_ID     int64
	TRANSACTION_ID            int64
	ACCOUNT_ID                int64
	LEDGER_TYPE_CODE          string
	LEDGER_TRANSACTION_AMOUNT float64
	CREATED_DATE              string
	CREATED_BY                string
	UPDATED_DATE              string
	UPDATED_BY                string
}
