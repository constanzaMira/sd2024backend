package database
 
import (
    "fmt"
    "log"
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "hello-run/model"
)
 
var DB *gorm.DB
 
func Connect() {
    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
 
    log.Printf("DB_HOST: %s", host)
    log.Printf("DB_USER: %s", user)
    log.Printf("DB_NAME: %s", dbname)
 
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
    log.Printf("DSN: %s", dsn)
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }
 
    log.Println("Database connected successfully")
 
    err = DB.AutoMigrate(&model.User{})
    if err != nil {
        log.Fatalf("Error migrating User model: %v", err)
    }
 
    err = DB.AutoMigrate(&model.Reservation{})
    if err != nil {
        log.Fatalf("Error migrating Reservation model: %v", err)
    }
}