package handler




import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var databas *gorm.DB
var err error

func init() {
    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }
}

func Dbconnect() {
    // Retrieve the database URL from environment
    Dbconnectur := os.Getenv("Dgconnect")
    if Dbconnectur == "" {
        log.Fatal("Dgconnect environment variable not set")
    }

    // Open database connection
    databas, err = gorm.Open(mysql.Open(Dbconnectur), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to initialize database, got error: %v", err)
    }

    // Automatically migrate schema
    databas.AutoMigrate(&CreateUserData{})
}
