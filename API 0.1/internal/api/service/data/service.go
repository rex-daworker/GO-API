package data

import (
	"context"
	"goapi/internal/api/repository/models"
	"log"
	"time"
)

type DataServiceSQLite struct {
	repo models.DataRepository
}

func NewDataServiceSQLite(repo models.DataRepository) *DataServiceSQLite {
	return &DataServiceSQLite{repo: repo}
}

func (ds *DataServiceSQLite) Create(data *models.Data, ctx context.Context) error {
	if err := ds.ValidateData(data); err != nil {
		return err // ✅ return actual validation error
	}
	return ds.repo.Create(data, ctx)
}

func (ds *DataServiceSQLite) ReadOne(id int, ctx context.Context) (*models.Data, error) {
	return ds.repo.ReadOne(id, ctx)
}

func (ds *DataServiceSQLite) ReadMany(page int, rowsPerPage int, ctx context.Context) ([]*models.Data, error) {
	return ds.repo.ReadMany(page, rowsPerPage, ctx)
}

func (ds *DataServiceSQLite) Update(data *models.Data, ctx context.Context) (int64, error) {
	if err := ds.ValidateData(data); err != nil {
		return 0, err
	}
	return ds.repo.Update(data, ctx)
}

func (ds *DataServiceSQLite) Delete(data *models.Data, ctx context.Context) (int64, error) {
	return ds.repo.Delete(data, ctx)
}

func (ds *DataServiceSQLite) ValidateData(data *models.Data) error {
	var errMsg string

	if data.DeviceID == "" || len(data.DeviceID) > 50 {
		errMsg += "DeviceID is required and must be less than 50 characters. "
	}
	if len(data.DeviceName) > 50 {
		errMsg += "DeviceName must be less than 50 characters. "
	}
	if len(data.Type) > 20 {
		errMsg += "Type must be less than 20 characters. "
	}
	if len(data.Description) > 100 {
		errMsg += "Description must be less than 100 characters. "
	}
	if _, err := time.Parse("2006-01-02T15:04:05Z", data.DateTime); err != nil {
		errMsg += "DateTime must be in format 2006-01-02T15:04:05Z. "
	}
	if data.Reading < 0 { 
		errMsg += "Reading must be >= 0. "
	}
	if data.Status != "active" && data.Status != "inactive" {
		errMsg += "Status must be 'active' or 'inactive'. "
	}
	if data.CreatedAt == "" {
		data.CreatedAt = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	}

	if errMsg != "" {
		log.Println("Validation failed:", errMsg) // ✅ helpful log
		return DataError{Message: errMsg}
	}
	return nil
}
