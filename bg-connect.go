package handler

import (
    "log"
    "os"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    _ "github.com/joho/godotenv/autoload" // Auto-load .env file in local development
)

var databas *gorm.DB

func Dbconnect() {
    dbURL := os.Getenv("Dgconnect")
    if dbURL == "" {
        log.Fatal("Database URL (Dgconnect) is not set")
    }

    var err error
    databas, err = gorm.Open(mysql.Open(dbURL), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    // Optional: Automatically migrate schema
    databas.AutoMigrate(&CreateUserData{})
}
