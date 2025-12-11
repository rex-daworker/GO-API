// internal/api/repository/DAL/SQLite/sqlite.go
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
        CREATE TABLE IF NOT EXISTS parking_events (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            slot_id TEXT NOT NULL,
            vehicle_id TEXT,
            distance_cm REAL,
            status TEXT NOT NULL,
            action TEXT NOT NULL,
            threshold_cm INTEGER NOT NULL,
            updated_at TEXT NOT NULL,
            note TEXT
        );
    `)
    if err != nil {
        return nil, err
    }

    return db, nil
}
