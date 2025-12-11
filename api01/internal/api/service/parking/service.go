// internal/api/service/parking/service.go
package parking

import (
    "context"
    "strings"
    "time"

    "goapi/internal/api/repository/models"
)


type ParkingServiceSQLite struct {
    repo models.ParkingRepository
}

func NewParkingServiceSQLite(repo models.ParkingRepository) *ParkingServiceSQLite {
    return &ParkingServiceSQLite{repo: repo}
}

func (ps *ParkingServiceSQLite) Create(ev *models.ParkingEvent, ctx context.Context) error {
    if err := ps.Validate(ev); err != nil {
        return err
    }
    return ps.repo.Create(ev, ctx)
}

func (ps *ParkingServiceSQLite) ReadOne(id int, ctx context.Context) (*models.ParkingEvent, error) {
    return ps.repo.ReadOne(id, ctx)
}

func (ps *ParkingServiceSQLite) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.ParkingEvent, error) {
    return ps.repo.ReadMany(page, rowsPerPage, ctx)
}

func (ps *ParkingServiceSQLite) Update(ev *models.ParkingEvent, ctx context.Context) (int64, error) {
    if err := ps.Validate(ev); err != nil {
        return 0, err
    }
    return ps.repo.Update(ev, ctx)
}

func (ps *ParkingServiceSQLite) Delete(ev *models.ParkingEvent, ctx context.Context) (int64, error) {
    return ps.repo.Delete(ev, ctx)
}

func (ps *ParkingServiceSQLite) Validate(ev *models.ParkingEvent) error {
    var errs []string

    if ev.SlotID == "" || len(ev.SlotID) > 20 {
        errs = append(errs, "slot_id is required and must be <= 20 chars")
    }
    if ev.Status != "free" && ev.Status != "occupied" && ev.Status != "unknown" {
        errs = append(errs, "status must be 'free', 'occupied', or 'unknown'")
    }
    if ev.Action != "open" && ev.Action != "close" && ev.Action != "none" {
        errs = append(errs, "action must be 'open', 'close', or 'none'")
    }
    if ev.ThresholdCM < 5 || ev.ThresholdCM > 200 {
        errs = append(errs, "threshold_cm must be between 5 and 200")
    }
    if ev.DistanceCM < 0 || ev.DistanceCM > 1000 {
        errs = append(errs, "distance_cm must be between 0 and 1000")
    }
    if ev.UpdatedAt == "" {
        ev.UpdatedAt = time.Now().UTC().Format("2006-01-02T15:04:05Z")
    }
    if len(errs) > 0 {
        return ParkingError{Message: strings.Join(errs, "; ")}
    }
    return nil
}
