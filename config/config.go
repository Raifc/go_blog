package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

var (
    DBDriver string
    DBName   string
)

func Init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Missing .env")
    }

    DBDriver = os.Getenv("DB_DRIVER")
    DBName = os.Getenv("DB_NAME")
}
