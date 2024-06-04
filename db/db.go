package db

import (
    "database/sql"
    "log"

    "go-blog/config"

    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
    var err error
    DB, err = sql.Open(config.DBDriver, config.DBName)
    if err != nil {
        log.Fatal(err)
    }

    createTable()
}

func createTable() {
    sqlStmt := `
    CREATE TABLE IF NOT EXISTS blogpost (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        content TEXT
    );
    `
    _, err := DB.Exec(sqlStmt)
    if err != nil {
        log.Fatalf("%q: %s\n", err, sqlStmt)
    }
}

func Close() {
    DB.Close()
}
