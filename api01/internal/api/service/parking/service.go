package parking

import (
    "context"
    "strings"
    "time"

    "goapi/internal/api/repository/models"
)

type repo interface {
    Create(ev *models.ParkingEvent, ctx context.Context) error
    ReadOne(id int, ctx context.Context) (*models.ParkingEvent, error)
    ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.ParkingEvent, error)
    Update(ev *models.ParkingEvent, ctx context.Context) (int64, error)
    Delete(ev *models.ParkingEvent, ctx context.Context) (int64, error)
}

type parkingServiceSQLite struct {
    repo repo
}

func NewParkingServiceSQLite(r repo) ParkingService {
    return &parkingServiceSQLite{repo: r}
}

func (s *parkingServiceSQLite) Create(ev *models.ParkingEvent, ctx context.Context) error {
    // Default updated_at if empty
    if strings.TrimSpace(ev.UpdatedAt) == "" {
        ev.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
    }
    return s.repo.Create(ev, ctx)
}

func (s *parkingServiceSQLite) ReadOne(id int, ctx context.Context) (*models.ParkingEvent, error) {
    return s.repo.ReadOne(id, ctx)
}

func (s *parkingServiceSQLite) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.ParkingEvent, error) {
    return s.repo.ReadMany(page, rowsPerPage, ctx)
}

func (s *parkingServiceSQLite) Update(ev *models.ParkingEvent, ctx context.Context) (int64, error) {
    if strings.TrimSpace(ev.UpdatedAt) == "" {
        ev.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
    }
    return s.repo.Update(ev, ctx)
}

func (s *parkingServiceSQLite) Delete(ev *models.ParkingEvent, ctx context.Context) (int64, error) {
    return s.repo.Delete(ev, ctx)
}

func (s *parkingServiceSQLite) Validate(ev *models.ParkingEvent) error {
    // Minimal validation rules
    if strings.TrimSpace(ev.SlotID) == "" {
        return ParkingError{Message: "slot_id is required"}
    }
    if ev.ThresholdCM <= 0 {
        return ParkingError{Message: "threshold_cm must be > 0"}
    }
    if strings.TrimSpace(ev.Status) == "" {
        return ParkingError{Message: "status is required"}
    }
    if strings.TrimSpace(ev.Action) == "" {
        return ParkingError{Message: "action is required"}
    }
    return nil
}
