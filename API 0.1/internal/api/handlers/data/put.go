package data

import (
	"encoding/json"
	"goapi/internal/api/repository/models"
	serviceData "goapi/internal/api/service/data"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type PutHandler struct {
	Service serviceData.DataService
}

func NewPutHandler(service serviceData.DataService) *PutHandler {
	return &PutHandler{Service: service}
}

func (h *PutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var payload models.Data

	// Extract ID from URL path (/data/{id})
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 2 {
		http.Error(w, "Missing ID in URL", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}
	payload.ID = id

	// Decode JSON body
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		log.Println("JSON decode error:", err)
		return
	}
	log.Printf("Received PUT payload for ID %d: %+v\n", id, payload)

	// Update via service (returns validation errors)
	if _, err := h.Service.Update(&payload, r.Context()); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Update error:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Data updated successfully"))
}
