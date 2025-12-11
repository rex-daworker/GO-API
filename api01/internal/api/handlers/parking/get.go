package parking

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"goapi/internal/api/service/parking"
)

slotID := r.URL.Query().Get("slot_id")
pageStr := r.URL.Query().Get("page")
page := 1
if pageStr != "" {
    if p, err := strconv.Atoi(pageStr); err == nil {
        page = p
    }
}

events, err := h.Service.ReadManyFiltered(slotID, page, 10, r.Context())

type GetHandler struct {
	Service parking.ParkingService
}

func NewGetHandler(service parking.ParkingService) *GetHandler {
	return &GetHandler{Service: service}
}

func (h *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	pageStr := query.Get("page")
	page := 1
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil {
			page = p
		}
	}

	events, err := h.Service.ReadMany(page, 10, r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch events", http.StatusInternalServerError)
		log.Println("ReadMany error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}
