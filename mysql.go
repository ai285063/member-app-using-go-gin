package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

const (
	USERNAME = "root"
	PASSWORD = "root"
	NETWORK  = "tcp"
	// SERVER   = "127.0.0.1"
	// docker-compose 裡面有自己的 dns，api 如果在 docker 裡面  不能用127.0.0.1
	SERVER   = "mysql"
	PORT     = 3306
	DATABASE = "practice"
)

func ConnectMysql() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	var err error
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("MySQL connection failed: " + err.Error())
	} else {
		log.Println("MySQL connected.")
	}

	if err := MysqlDB.AutoMigrate(&User{}); err != nil {
		panic("MySql create table failed: " + err.Error())
	}
}
