package configs

import (
	"fmt"
	"os"
	"prakerja/eventmanagement/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	if env := os.Getenv("WEB_ENV"); env != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			panic("gagal load file")
		}
	}
}

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func ConnectDatabase() {
	var dbConfig DBConfig = DBConfig{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database Error")
	}
	migration()
}

func migration() {
	DB.AutoMigrate(&models.Event{})
}
