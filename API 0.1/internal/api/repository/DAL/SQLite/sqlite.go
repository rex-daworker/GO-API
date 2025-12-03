package SQLite

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func NewSqlite(filename string) (*sql.DB, error) {
    db, err := sql.Open("sqlite3", filename)
    if err != nil {
        return nil, err
    }

    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS data (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            device_id TEXT,
            device_name TEXT,
            reading REAL,
            type TEXT,
            date_time TEXT,
            description TEXT,
            status TEXT,
            created_at TEXT
        );
    `)
    if err != nil {
        return nil, err
    }

    return db, nil
}
