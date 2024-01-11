package db

import (
	"fmt"
	"sample/pkg/config"
	"sample/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {

	psqlInfo := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s", cfg.DBUser, cfg.DBName, cfg.DBPassword, cfg.DBHost, cfg.DBPort)
	// psqlInfo:="user=postgres dbname=company password=123 host=localhost port=5432 "
	db, dberr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if dberr != nil {
		fmt.Println("error in db connection", dberr)
	}

	db.AutoMigrate(&domain.Users{})

	// Create a admin in postgress database using terminal with credential of name and password;
	db.AutoMigrate(&domain.Admin{})
	return db, nil
}
