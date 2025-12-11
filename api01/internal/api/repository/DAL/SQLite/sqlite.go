package SQLite

import (
    "database/sql"

    _ "github.com/mattn/go-sqlite3"
)

func Open(path string) (*sql.DB, error) {
    return sql.Open("sqlite3", path)
}

func EnsureSchema(db *sql.DB) error {
    ddl := `
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
);`
    _, err := db.Exec(ddl)
    return err
}
