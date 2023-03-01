package internal

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {
	dsn := "host=localhost user=gorm password=gorm dbname=cash port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf(err.Error())
	}

	err = DB.AutoMigrate(&Income{}, &Outcome{})
	if err != nil {
		log.Printf(err.Error())
	}
}
