package main

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

    // Automatically migrate schema
    if err := databas.AutoMigrate(&CreateUserData{}); err != nil {
        log.Fatalf("failed to migrate schema: %v", err)
    }
    
    log.Println("Database connection and migration successful")
}
