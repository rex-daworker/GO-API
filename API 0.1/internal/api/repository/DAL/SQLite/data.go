package SQLite

import (
    "context"
    "database/sql"
    "goapi/internal/api/repository/models"
)

type DataRepositorySQLite struct {
    sqlDB *sql.DB
}

func NewDataRepository(db *sql.DB) *DataRepositorySQLite {
    return &DataRepositorySQLite{sqlDB: db}
}

func (repo *DataRepositorySQLite) Create(data *models.Data, ctx context.Context) error {
    stmt := `INSERT INTO data (
        device_id, device_name, reading, type, date_time, description, status, created_at
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

    _, err := repo.sqlDB.ExecContext(ctx, stmt,
        data.DeviceID, data.DeviceName, data.Reading,
        data.Type, data.DateTime, data.Description,
        data.Status, data.CreatedAt)
    return err
}

func (repo *DataRepositorySQLite) ReadOne(id int, ctx context.Context) (*models.Data, error) {
    stmt := `SELECT id, device_id, device_name, reading, type, date_time, description, status, created_at 
             FROM data WHERE id = ?`
    row := repo.sqlDB.QueryRowContext(ctx, stmt, id)

    var d models.Data
    err := row.Scan(&d.ID, &d.DeviceID, &d.DeviceName, &d.Reading,
        &d.Type, &d.DateTime, &d.Description, &d.Status, &d.CreatedAt)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    return &d, nil
}

func (repo *DataRepositorySQLite) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Data, error) {
    stmt := `SELECT id, device_id, device_name, reading, type, date_time, description, status, created_at 
             FROM data LIMIT ? OFFSET ?`
    rows, err := repo.sqlDB.QueryContext(ctx, stmt, rowsPerPage, page*rowsPerPage)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []*models.Data
    for rows.Next() {
        var d models.Data
        if err := rows.Scan(&d.ID, &d.DeviceID, &d.DeviceName, &d.Reading,
            &d.Type, &d.DateTime, &d.Description, &d.Status, &d.CreatedAt); err != nil {
            return nil, err
        }
        results = append(results, &d)
    }
    return results, nil
}

func (repo *DataRepositorySQLite) Update(data *models.Data, ctx context.Context) (int64, error) {
    stmt := `UPDATE data SET 
        device_id = ?, device_name = ?, reading = ?, type = ?, date_time = ?, description = ?, status = ?, created_at = ?
        WHERE id = ?`
    res, err := repo.sqlDB.ExecContext(ctx, stmt,
        data.DeviceID, data.DeviceName, data.Reading,
        data.Type, data.DateTime, data.Description,
        data.Status, data.CreatedAt, data.ID)
    if err != nil {
        return 0, err
    }
    return res.RowsAffected()
}

func (repo *DataRepositorySQLite) Delete(data *models.Data, ctx context.Context) (int64, error) {
    stmt := `DELETE FROM data WHERE id = ?`
    res, err := repo.sqlDB.ExecContext(ctx, stmt, data.ID)
    if err != nil {
        return 0, err
    }
    return res.RowsAffected()
}
