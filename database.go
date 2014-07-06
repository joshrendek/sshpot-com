package main

import (
	"os"

	"github.com/jinzhu/gorm"
)

var DB gorm.DB

func SetupDB() {
	dbString := os.Getenv("DO_POSTGRESQL")
	if dbString == "" {
		dbString = os.Getenv("HEROKU_POSTGRESQL_BLUE_URL")
		if dbString == "" {
			dbString = "user=joshrendek port=5432 dbname=ssh_honey sslmode=disable"
		}
	}
	DB, err = gorm.Open("postgres", dbString)

	if err != nil {
		panic(err)
	}

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(50)
	DB.DB().Ping()
	DB.LogMode(true)

	DB.CreateTable(SshLogin{})
	DB.CreateTable(ApiStat{})
	DB.CreateTable(HoneypotServer{})

}
