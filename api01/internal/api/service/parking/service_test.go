package parking_test

import (
    "context"
    "testing"

    "goapi/internal/api/repository/models"
    "goapi/internal/api/service/parking"
)

type dummyRepo struct{}

func (d dummyRepo) Create(ev *models.ParkingEvent, ctx context.Context) error { return nil }
func (d dummyRepo) ReadOne(id int, ctx context.Context) (*models.ParkingEvent, error) { return nil, nil }
func (d dummyRepo) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.ParkingEvent, error) {
    return nil, nil
}
func (d dummyRepo) Update(ev *models.ParkingEvent, ctx context.Context) (int64, error) { return 1, nil }
func (d dummyRepo) Delete(ev *models.ParkingEvent, ctx context.Context) (int64, error) { return 1, nil }

func TestValidateOK(t *testing.T) {
    svc := parking.NewParkingServiceSQLite(dummyRepo{})

    ev := &models.ParkingEvent{
        SlotID:      "A1",
        Status:      "occupied",
        Action:      "close",
        ThresholdCM: 30,
        DistanceCM:  15,
        UpdatedAt:   "2025-12-11T18:00:00Z",
    }

    if err := svc.Validate(ev); err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
}
