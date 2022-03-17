INSERT INTO rm_account_type values (0,"cash",current_timestamp(),'Admin',current_timestamp(),'Admin');

INSERT INTO rm_account VALUES(1,'12389499494',0,10000,CURRENT_TIMESTAMP(),'Admin',CURRENT_TIMESTAMP(),'Admin');
INSERT INTO rm_account VALUES(2,'83589828405',0,10000,CURRENT_TIMESTAMP(),'Admin',CURRENT_TIMESTAMP(),'Admin');
INSERT INTO rm_account VALUES(3,'40589249020',0,10000,CURRENT_TIMESTAMP(),'Admin',CURRENT_TIMESTAMP(),'Admin');


INSERT into rm_user_info VALUES(1,'18891A05D6','Anirudh',' ','G',123456789,1,'nishanth@mail.com',1,current_timestamp(),'Admin',current_timestamp(),'Admin');
INSERT into rm_user_info VALUES(2,'18891A05D5','Rakshith',' ','A',123456788,1,'nishanth@mail.com',1,current_timestamp(),'Admin',current_timestamp(),'Admin');
INSERT into rm_user_info VALUES(3,'18891A05D0','Nishanth',' ','B',123456787,1,'nishanth@mail.com',1,current_timestamp(),'Admin',current_timestamp(),'Admin');

INSERT INTO rm_user_account VALUES(1,1,1,CURRENT_TIMESTAMP(),'Admin',CURRENT_TIMESTAMP(),'Admin');
INSERT INTO rm_user_account VALUES(2,2,2,CURRENT_TIMESTAMP(),'Admin',CURRENT_TIMESTAMP(),'Admin');
INSERT INTO rm_user_account VALUES(3,3,3,CURRENT_TIMESTAMP(),'Admin',CURRENT_TIMESTAMP(),'Admin');