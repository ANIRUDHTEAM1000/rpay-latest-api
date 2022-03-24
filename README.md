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
    * ``` mariadb -u <user-name> -p<Password> ```
    * ``` use <database name> ```
    *  now copy the path till the database folder eg:  C:\Users\user1\OneDrive\Desktop\rpay-latest-api\resources\database
    * ``` source <paste the copied path>\create_tables.sql ``` 
    * ``` source <paste the copied path>\tempdata.sql ```
4. after step 3 our database is ready with all required tables and some temporary data.
5. open the integrated terminal in vsCode and navigate as cmd/main where our main.go is located.
6. now run the command
    * ``` go run main.go ```
7. Now the api is running on local host 8080
8. test the [api calls](#api-calls) on postman
9. Use the database client to see the data in tables.

## API calls



## Understanding the file structure
