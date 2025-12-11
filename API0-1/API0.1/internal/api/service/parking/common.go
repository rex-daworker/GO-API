package parking
// internal/api/service/parking/common.go

import (
    "context"
    "github.com/rex-daworker/GO-API/internal/api/repository/models"
)

type ParkingService interface {
    Create(ev *models.ParkingEvent, ctx context.Context) error
    ReadOne(id int, ctx context.Context) (*models.ParkingEvent, error)
    ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.ParkingEvent, error)
    Update(ev *models.ParkingEvent, ctx context.Context) (int64, error)
    Delete(ev *models.ParkingEvent, ctx context.Context) (int64, error)
    Validate(ev *models.ParkingEvent) error
}

type ParkingError struct {
    Message string
}

func (e ParkingError) Error() string { return e.Message }
