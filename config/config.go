package mysql

import (
	"fmt"
	"os"
	"popaket/migration"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//connection to database
func Connection() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		return nil
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	//migration
	db.AutoMigrate(&migration.Auths{})
	db.AutoMigrate(&migration.Logistics{})

	return db
}
