package SQLite

import (
	"context"
	"database/sql"

	"goapi/internal/api/repository/models"
)

type parkingRepo struct {
	db *sql.DB
}

func NewParkingRepository(db *sql.DB) models.ParkingRepository {
	return &parkingRepo{db: db}
}

func (r *parkingRepo) Create(ev *models.ParkingEvent, ctx context.Context) error {
	_, err := r.db.ExecContext(ctx, `
        INSERT INTO parking_events 
        (slot_id, vehicle_id, distance_cm, status, action, threshold_cm, updated_at, note)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		ev.SlotID, ev.VehicleID, ev.DistanceCM, ev.Status, ev.Action, ev.ThresholdCM, ev.UpdatedAt, ev.Note,
	)
	return err
}

func (r *parkingRepo) ReadOne(id int, ctx context.Context) (*models.ParkingEvent, error) {
	row := r.db.QueryRowContext(ctx, `
        SELECT id, slot_id, vehicle_id, distance_cm, status, action, threshold_cm, updated_at, note
        FROM parking_events WHERE id = ?`, id,
	)
	var ev models.ParkingEvent
	err := row.Scan(
		&ev.ID, &ev.SlotID, &ev.VehicleID, &ev.DistanceCM, &ev.Status, &ev.Action, &ev.ThresholdCM, &ev.UpdatedAt, &ev.Note,
	)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

func (r *parkingRepo) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.ParkingEvent, error) {
	if rowsPerPage <= 0 {
		rowsPerPage = 10
	}
	offset := (page - 1) * rowsPerPage
	rows, err := r.db.QueryContext(ctx, `
        SELECT id, slot_id, vehicle_id, distance_cm, status, action, threshold_cm, updated_at, note
        FROM parking_events
        ORDER BY id DESC
        LIMIT ? OFFSET ?`, rowsPerPage, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*models.ParkingEvent
	for rows.Next() {
		var ev models.ParkingEvent
		if err := rows.Scan(
			&ev.ID, &ev.SlotID, &ev.VehicleID, &ev.DistanceCM, &ev.Status, &ev.Action, &ev.ThresholdCM, &ev.UpdatedAt, &ev.Note,
		); err != nil {
			return nil, err
		}
		result = append(result, &ev)
	}
	return result, rows.Err()
}

func (r *parkingRepo) Update(ev *models.ParkingEvent, ctx context.Context) (int64, error) {
	res, err := r.db.ExecContext(ctx, `
        UPDATE parking_events
        SET slot_id = ?, vehicle_id = ?, distance_cm = ?, status = ?, action = ?, threshold_cm = ?, updated_at = ?, note = ?
        WHERE id = ?`,
		ev.SlotID, ev.VehicleID, ev.DistanceCM, ev.Status, ev.Action, ev.ThresholdCM, ev.UpdatedAt, ev.Note, ev.ID,
	)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (r *parkingRepo) Delete(ev *models.ParkingEvent, ctx context.Context) (int64, error) {
	res, err := r.db.ExecContext(ctx, `DELETE FROM parking_events WHERE id = ?`, ev.ID)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
