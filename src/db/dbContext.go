package db

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm/logger"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createPostgresPool(connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(0) // No limit on connection lifetime

	return db, nil
}

var DB *gorm.DB

func InitDB() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectionString := os.Getenv("DATABASE_URL")
	DB, err := createPostgresPool(connectionString)
	if err != nil {
		log.Fatalf("Error creating database connection: %v", err)
	}
	if DB == nil {
		log.Fatal("Database connection is nil")
		return nil
	}

	log.Println("Database connection established successfully")
	return nil
}
