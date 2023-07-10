package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var err error

type Output struct {
	Id              int     `json:"id"`
	ValorConvertido float64 `json:"valorConvertido"`
	SimboloMoeda    string  `json:"SimboloMoeda"`
}

func ConnectionDB() {
	dns := "user=root dbname=genesis password=root host=localhost sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("Error on connection database", err)
	}

	DB.AutoMigrate(Output{})
}
