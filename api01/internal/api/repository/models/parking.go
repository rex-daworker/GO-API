package models

import "context"

type ParkingEvent struct {
	ID          int     `json:"id"`
	SlotID      string  `json:"slot_id"`
	VehicleID   string  `json:"vehicle_id,omitempty"`
	DistanceCM  float64 `json:"distance_cm,omitempty"`
	Status      string  `json:"status"`
	Action      string  `json:"action"`
	ThresholdCM int     `json:"threshold_cm"`
	UpdatedAt   string  `json:"updated_at"`
	Note        string  `json:"note,omitempty"`
}

type ParkingRepository interface {
	Create(ev *ParkingEvent, ctx context.Context) error
	ReadOne(id int, ctx context.Context) (*ParkingEvent, error)
	ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*ParkingEvent, error)
	Update(ev *ParkingEvent, ctx context.Context) (int64, error)
	Delete(ev *ParkingEvent, ctx context.Context) (int64, error)
}
