INSERT INTO RT_TRANSACTION(TRANSACTION_UNIQUE_ID,TRANSACTION_TYPE_CODE,TRANSACTION_AMOUNT,CREATED_BY,UPDATED_BY)
values('T1','b','100','Admin','Admin');

INSERT INTO RT_TRANSACTION_LEDGER(TRANSACTION_ID,ACCOUNT_ID,LEDGER_TYPE_CODE,LEDGER_TRANSACTION_AMOUNT,CREATED_BY,UPDATED_BY)
 values (1,1,'D',-100,'Admin','Admin');
INSERT INTO RT_TRANSACTION_LEDGER(TRANSACTION_ID,ACCOUNT_ID,LEDGER_TYPE_CODE,LEDGER_TRANSACTION_AMOUNT,CREATED_BY,UPDATED_BY)
 values (1,2,'C',+100,'Admin','Admin');

UPDATE RM_ACCOUNT set MONEY_ACCOUNT_BALANCE = MONEY_ACCOUNT_BALANCE+100 where MONEY_ACCOUNT_ID='accRakshith' ;
UPDATE RM_ACCOUNT set MONEY_ACCOUNT_BALANCE = MONEY_ACCOUNT_BALANCE-100 where MONEY_ACCOUNT_ID='accAnirudh' ;
