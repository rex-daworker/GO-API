package data

import (
	"context"
	"encoding/json"
	"goapi/internal/api/repository/models"
	service "goapi/internal/api/service/data"
	"log"
	"net/http"
	"time"
)

// * When using PUT, the client sends a complete representation of a resource to replace the current version: Whole Resource Replacement. *
// * curl -X PUT http://127.0.0.1:8080/data -i -u admin:password -H "Content-Type: application/json" -d '{"id": 1, "device_id": "device1", "device_name": "device1", "reading": 2.0, "type": "type1", "date_time": "2021-01-01T00:00:00Z", "description": "updated description", "status": "active", "created_at": "2021-01-01T00:00:00Z"}'
func PutHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger, ds service.DataService) {
	var data models.Data

	// * Decode the JSON payload from the request body into the data struct
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		// * This is a User Error: format of body is invalid, response in JSON and with a 400 status code
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid request data. Please check your input."}`))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// * Try to update the data in the database
	if aff, err := ds.Update(&data, ctx); err != nil {
		switch err.(type) {
		case service.DataError:
			// * If the error is a DataError, handle it as a client error
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "` + err.Error() + `"}`))
			return
		default:
			// * If it is not a DataError, handle it as a server error
			logger.Println("Error creating data:", err, data)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Internal server error."}`))
			return
		}
	} else if aff == 0 {
		// * This is a User Error, response in JSON and with a 404 status code
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Resource not found."}`))
		return
	}

	// * Return the data to the user as JSON with a 200 OK status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Println("Error encoding data:", err, data)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal server error."}`))
		return
	}
}
