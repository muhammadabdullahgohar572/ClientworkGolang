package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var urlbfconnect = "root:ilove1382005#@tcp(localhost:3306)/practice2?parseTime=true"

var databas *gorm.DB

var err  error;

func Dbconnect() {

	databas,err =gorm.Open(mysql.Open(urlbfconnect),&gorm.Config{})
	if err != nil {
		panic("DataBaseconnection Error")
	}
	databas.AutoMigrate(&CreateUserData{}) 
}



