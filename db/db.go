package db

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
    "go-blog/config"
)

var DB *sql.DB

func Init() {
    var err error
    dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

    DB, err = sql.Open(config.DBDriver, dataSourceName)
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Failed to ping the database:", err)
    }

    createTable()
}

func createTable() {
    query := `
    CREATE TABLE IF NOT EXISTS blogpost (
        id SERIAL PRIMARY KEY,
        title VARCHAR(100) NOT NULL,
        content TEXT NOT NULL
    );
    `
    _, err := DB.Exec(query)
    if err != nil {
        log.Fatal("Failed to create table:", err)
    }
}

func Close() {
    if DB != nil {
        DB.Close()
    }
}
