package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

var (
    DBDriver   string
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
)

func Init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Missing .env")
    }

    DBDriver = os.Getenv("DB_DRIVER")
    DBHost = os.Getenv("DB_HOST")
    DBPort = os.Getenv("DB_PORT")
    DBUser = os.Getenv("DB_USER")
    DBPassword = os.Getenv("DB_PASSWORD")
    DBName = os.Getenv("DB_NAME")
}
