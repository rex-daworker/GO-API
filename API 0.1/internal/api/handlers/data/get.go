package data

import (
	"context"
	"encoding/json"
	service "goapi/internal/api/service/data"
	"log"
	"net/http"
	"strconv"
	"time"
)

// * The GET method retrieves all resources identified by a URI *
// * curl -X GET http://127.0.0.1:8080/data -i -u admin:password -H "Content-Type: application/json"
func GetHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger, ds service.DataService) {
	// Parse optional page query parameter; default to 0 when missing
	pageStr := r.URL.Query().Get("page")
	var page int
	if pageStr == "" {
		page = 0
	} else {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Invalid page specified."}`))
			return
		}
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	data, err := ds.ReadMany(page, 10, ctx)
	if err != nil {
		logger.Println("Could not get data:", err, data)
		http.Error(w, "Internal Server error.", http.StatusInternalServerError)
		return
	}
	if len(data) == 0 {
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
		w.Write([]byte(`{"error": "Internal Server error."}`))
		return
	}
}
