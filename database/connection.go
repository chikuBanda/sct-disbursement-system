package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sct-system/models"
)

var DB *gorm.DB

func StartDbConnection() {
	dsn := "chiku:mygorestapp@tcp(127.0.0.1:3306)/sct_system_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	println("Connected to database")

	err = db.AutoMigrate(
		&models.Beneficiary{},
		&models.Transfer{},
		&models.TransactionHistory{},
	)
	if err != nil {
		panic("failed to auto migrate")
	}
	println("Database migrated")

	DB = db
}
