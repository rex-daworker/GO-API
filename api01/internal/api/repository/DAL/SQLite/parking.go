// internal/api/repository/DAL/SQLite/parking.go
package SQLite

import (
	"context"
	"database/sql"

	"goapi/internal/api/repository/models"
)

type ParkingRepositorySQLite struct {
	sqlDB *sql.DB
}

func NewParkingRepository(db *sql.DB) *ParkingRepositorySQLite {
	return &ParkingRepositorySQLite{sqlDB: db}
}

func (repo *ParkingRepositorySQLite) Create(ev *models.ParkingEvent, ctx context.Context) error {
	stmt := `INSERT INTO parking_events (slot_id, vehicle_id, distance_cm, status, action, threshold_cm, updated_at, note)
             VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := repo.sqlDB.ExecContext(ctx, stmt,
		ev.SlotID, ev.VehicleID, ev.DistanceCM, ev.Status, ev.Action, ev.ThresholdCM, ev.UpdatedAt, ev.Note)
	return err
}

func (repo *ParkingRepositorySQLite) ReadOne(id int, ctx context.Context) (*models.ParkingEvent, error) {
	stmt := `SELECT id, slot_id, vehicle_id, distance_cm, status, action, threshold_cm, updated_at, note
             FROM parking_events WHERE id = ?`
	row := repo.sqlDB.QueryRowContext(ctx, stmt, id)

	var ev models.ParkingEvent
	err := row.Scan(&ev.ID, &ev.SlotID, &ev.VehicleID, &ev.DistanceCM, &ev.Status, &ev.Action, &ev.ThresholdCM, &ev.UpdatedAt, &ev.Note)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

func (repo *ParkingRepositorySQLite) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.ParkingEvent, error) {
	stmt := `SELECT id, slot_id, vehicle_id, distance_cm, status, action, threshold_cm, updated_at, note
             FROM parking_events ORDER BY updated_at DESC LIMIT ? OFFSET ?`
	rows, err := repo.sqlDB.QueryContext(ctx, stmt, rowsPerPage, page*rowsPerPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*models.ParkingEvent
	for rows.Next() {
		var ev models.ParkingEvent
		if err := rows.Scan(&ev.ID, &ev.SlotID, &ev.VehicleID, &ev.DistanceCM, &ev.Status, &ev.Action, &ev.ThresholdCM, &ev.UpdatedAt, &ev.Note); err != nil {
			return nil, err
		}
		results = append(results, &ev)
	}
	return results, nil
}

func (repo *ParkingRepositorySQLite) Update(ev *models.ParkingEvent, ctx context.Context) (int64, error) {
	stmt := `UPDATE parking_events SET slot_id = ?, vehicle_id = ?, distance_cm = ?, status = ?, action = ?, threshold_cm = ?, updated_at = ?, note = ? WHERE id = ?`
	res, err := repo.sqlDB.ExecContext(ctx, stmt,
		ev.SlotID, ev.VehicleID, ev.DistanceCM, ev.Status, ev.Action, ev.ThresholdCM, ev.UpdatedAt, ev.Note, ev.ID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (repo *ParkingRepositorySQLite) Delete(ev *models.ParkingEvent, ctx context.Context) (int64, error) {
	stmt := `DELETE FROM parking_events WHERE id = ?`
	res, err := repo.sqlDB.ExecContext(ctx, stmt, ev.ID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
