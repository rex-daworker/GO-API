package service

import (
	"context"
	"database/sql"
	"goapi/internal/api/repository/DAL/SQLite"
	data "goapi/internal/api/service/data"
	"log"
)

type DataServiceType int

const (
	SQLiteDataService DataServiceType = iota
)

type ServiceFactory struct {
	db     *sql.DB
	logger *log.Logger
	ctx    context.Context
}

func NewServiceFactory(db *sql.DB, logger *log.Logger, ctx context.Context) *ServiceFactory {
	return &ServiceFactory{db: db, logger: logger, ctx: ctx}
}

func (sf *ServiceFactory) CreateDataService(serviceType DataServiceType) (data.DataService, error) {
	switch serviceType {
	case SQLiteDataService:
		repo := SQLite.NewDataRepository(sf.db)
		ds := data.NewDataServiceSQLite(repo)
		return ds, nil
	default:
		return nil, data.DataError{Message: "Invalid data service type."}
	}
}
