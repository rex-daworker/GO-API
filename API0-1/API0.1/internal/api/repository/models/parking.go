// internal/api/repository/models/parking.go
package models

import "context"

type ParkingEvent struct {
    ID          int     `json:"id"`
    SlotID      string  `json:"slot_id"`       // e.g., "A1"
    VehicleID   string  `json:"vehicle_id"`    // optional: plate or session id
    DistanceCM  float64 `json:"distance_cm"`   // raw sensor reading
    Status      string  `json:"status"`        // "free", "occupied", "unknown"
    Action      string  `json:"action"`        // "open", "close", "none"
    ThresholdCM int     `json:"threshold_cm"`  // occupancy threshold
    UpdatedAt   string  `json:"updated_at"`    // ISO8601 UTC
    Note        string  `json:"note"`          // optional description
}

type ParkingRepository interface {
    Create(ev *ParkingEvent, ctx context.Context) error
    ReadOne(id int, ctx context.Context) (*ParkingEvent, error)
    ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*ParkingEvent, error)
    Update(ev *ParkingEvent, ctx context.Context) (int64, error)
    Delete(ev *ParkingEvent, ctx context.Context) (int64, error)
}
