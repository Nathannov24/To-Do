# To-Do List App

[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=blue)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=blue)](https://github.com/labstack/echo)


# Table of Content

- [Description](#description)
- [Requirements](#Requirements)
- [How to use](#how-to-use)
- [Our Feature](#Our-Feature)
- [Endpoints](#endpoints)
- [Credits](#credits)


# Description

Project ini bertujuan untuk membuat aplikasi To-Do List.

# Requirements

* Visual Studio Code
* Postman
* Mysql Workbench

# How to use
- Install Go, Postman, MySQL Workbench
- Clone this repository in your $PATH:
```
git clone https://github.com/Nathannov24/To-Do.git
```
* CREATE DATABASE IF NOT EXISTS `todo`;
* CREATE .env in root and paste
```
DB_USERNAME=root
DB_PASSWORD="Your Mysql Password"
DB_HOST=localhost
DB_PORT=3306
DB_NAME=todo
```
* Run `main.go`
```
$ go run main.go
```
* Open Postman run with your local host, follow the routes in Visual Studio Code folder.


# Our Feature
* CREATE : Create Activity
* READ : See your list of activity
* UPDATE : Update your activity to done
* DELETE : Delete your activity if it done


# Endpoints

| Method | Endpoint | Description| Authentication | Authorization
|:-----|:--------|:----------| :----------:| :----------:|
| POST  | /post | Post your activity | No | No
| GET | / | See your list of activity | No | No
| DELETE | /delete/:id | Delete Activity by ID | No | No
| PUT | /update-status/:id | Update Status Activity to Done by ID | No | No
|---|---|---|---|---|

<br>


# Credits

https://github.com/Nathannov24
