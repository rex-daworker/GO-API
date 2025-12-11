package parking

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"goapi/internal/api/service/parking"
)

type GetHandler struct {
	Service parking.ParkingService
}

func NewGetHandler(service parking.ParkingService) *GetHandler {
	return &GetHandler{Service: service}
}

func (h *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Parse query params
	query := r.URL.Query()

	page := 1
	if pageStr := query.Get("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil {
			page = p
		}
	}

	slotID := query.Get("slot_id")

	// Call service method
	var events []*parking.ParkingEvent
	var err error

	if slotID != "" {
		events, err = h.Service.ReadManyFiltered(slotID, page, 10, r.Context())
	} else {
		events, err = h.Service.ReadMany(page, 10, r.Context())
	}

	if err != nil {
		http.Error(w, "Failed to fetch events", http.StatusInternalServerError)
		log.Println("ReadMany error:", err)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(events); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
