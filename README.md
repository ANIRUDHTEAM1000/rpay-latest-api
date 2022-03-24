# RPay Wallet Engine API

### Table of Contents

1) [Project Descrption](#project-description)
2) [Requirements](#requirements)
3) [Installation](#installation)
4) [How to run the API](#how-to-run-the-api)
5) [all API calls](#api-calls)
5) [Understanding the file structure](#understanding-the-file-structure)

## Project Description


## Requirements
- golang
- mariadb
- any IDE (Visual Studio code) and database client
- Post man
- Docker


## Installation
### Golang Installation  
1) Download the Go Installer from https://go.dev/doc/install
2) Select the tab from the Go Install section on the page according to the operating system to setup the enviornment variable
3) You are ready to code.

### Mariadb Installation 
- For MacOS https://mariadb.com/resources/blog/installing-mariadb-10-1-16-on-mac-os-x-with-homebrew/
- For Windows https://www.mariadbtutorial.com/getting-started/install-mariadb/
- For Linux https://linuxize.com/post/how-to-install-mariadb-on-ubuntu-18-04/


## How to run the API
Clone the repository ( or ) download the zip and unzip in your computer.

To run the API we have two ways.
- Using docker
- without using docker

### 1) Using Docker
1) Start the docker using docker desktop. Make sure that the left bottom of docker desktop is green as shown below.
![Image](https://assets.digitalocean.com/67852/FTHGxfU.png)
1) Go to the terminal and navigate till the directory where the docker-compose file is located.
2) run the command     
  ``` docker-compose up --build ```
4) Now the api is running on local host 8080
5) test the [api calls](#api-calls) on postman

### 2) without using docker
1. open the cloned/downloaded api folder in your IDE editor.
2. open the resources/config.go and change the database name, user name and password with your database name and password.
3. open the terminal and enter into mariadb by using your name and password to do so run
    * ###
        ``` mariadb -u <user-name> -p<Password> ```
    * ``` use <databas- name> ```
    *  now copy the path till the database folder eg:  C:\Users\user1\OneDrive\Desktop\rpay-latest-api\resources\database
    * ``` source <paste_the_copied_path>\create_tables.sql ``` 
    * ``` source <paste_the_copied_path>\tempdata.sql ```
4. after step 3 our database is ready with all required tables and some temporary data.
5. open the integrated terminal in vsCode and navigate as cmd/main where our main.go is located.
6. now run the command
    * ``` go run main.go ```
7. Now the api is running on local host 8080
8. test the [api calls](#api-calls) on postman
9. Use the database client to see the data in tables.

## Get the user data using his USER_LOGIN_ID

`walletengine/user/<USER_LOGIN_ID>` (method : GET)

### Response
    {
    "status": 1,
    "name": "Anirudh G",
    "user_login_id": "18891A05D6",
    "balance": 7000
    }

## Get list of users searched using 

`walletengine/user/query<name-phone-email>` (method : GET)

### Response
    {
        "status": 1,
        "users": 
        [
            {
            "name": "Anirudh G",
            "user_login_id": "18891A05D6"
            },
            {
            "name": "Rakshith A",
            "user_login_id": "18891A05D5"
            },
            {
                "name": "Nishanth B",
                "user_login_id": "18891A05D0"
            }
        ]
    }

## Get the user data using his USER_LOGIN_ID

`walletengine/transfer` (method : POST)

### json input as follows
    {
    "sender":"18891A05D6",
    "receiver":"18891A05D0",
    "amount":1000
    }

### Response
    {
    "status": 1,
    "transaction_number": "b7bca2b1-79ca-4204-b144-a6b5a37f2aae_24032022_052205",
    "time": "2022-03-24 05:22:05",
    "amount": 1000,
    "sender": "18891A05D6",
    "receiver": "18891A05D0"
    }
    

## Understanding the file structure
```tree
.
├── README.md
├── app.Dockerfile
├── cmd
│   └── main
│       └── main.go
├── db.Dockerfile
├── docker-compose.yml
├── entrypoint.sh
├── go.mod
├── go.sum
├── pkg
│   ├── account
│   │   ├── models
│   │   │   ├── account.go
│   │   │   ├── account_type.go
│   │   │   └── user_account.go
│   │   ├── repository
│   │   │   └── account_repo.go
│   │   ├── routes
│   │   │   └── routes.go
│   │   └── services
│   │       └── services.go
│   ├── app
│   │   └── routes.go
│   ├── transaction
│   │   ├── dao
│   │   │   └── transaction_input.go
│   │   ├── repository
│   │   │   └── transactionrepo.go
│   │   ├── routes
│   │   │   └── routes.go
│   │   └── services
│   │       └── services.go
│   └── user
│       ├── dao
│       │   └── login_output.go
│       ├── models
│       │   └── user.go
│       ├── repository
│       │   └── UserRepo.go
│       ├── routes
│       │   └── routes.go
│       └── services
│           └── services.go
└── resources
    ├── config.go
    └── database
        ├── create_tables.sql
        ├── tables
        │   ├── a.sql
        │   ├── b.sql
        │   ├── c.sql
        │   ├── d.sql
        │   ├── e.sql
        │   ├── f.sql
        │   └── g.sql
        ├── tempdata.sql
        └── transaction.sql
```
