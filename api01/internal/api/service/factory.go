package service

import (
    "context"
    "database/sql"
    "log"

    "goapi/internal/api/repository/DAL/SQLite"
    "goapi/internal/api/service/parking"
)

type ParkingServiceType int

const (
    SQLiteParkingService ParkingServiceType = iota
)

type ServiceFactory struct {
    db     *sql.DB
    logger *log.Logger
    ctx    context.Context
}

func NewServiceFactory(db *sql.DB, logger *log.Logger, ctx context.Context) *ServiceFactory {
    return &ServiceFactory{db: db, logger: logger, ctx: ctx}
}

func (sf *ServiceFactory) CreateParkingService(serviceType ParkingServiceType) (parking.ParkingService, error) {
    switch serviceType {
    case SQLiteParkingService:
        repo := SQLite.NewParkingRepository(sf.db)
        svc := parking.NewParkingServiceSQLite(repo)
        return svc, nil
    default:
        return nil, parking.ParkingError{Message: "Invalid parking service type"}
    }
}
