package parking

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"goapi/internal/api/repository/models"
	"goapi/internal/api/service/parking"
)

type GetHandler struct {
	Service parking.ParkingService
}

func NewGetHandler(service parking.ParkingService) *GetHandler {
	return &GetHandler{Service: service}
}

func (h *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Query params: ?page=1&slot_id=A1
	q := r.URL.Query()

	// Page parsing with default = 1
	page := 1
	if ps := strings.TrimSpace(q.Get("page")); ps != "" {
		if p, err := strconv.Atoi(ps); err == nil && p > 0 {
			page = p
		}
	}

	slotID := strings.TrimSpace(q.Get("slot_id"))

	// Fetch a page from the service
	events, err := h.Service.ReadMany(page, 10, r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch events", http.StatusInternalServerError)
		log.Println("ReadMany error:", err)
		return
	}

	// Optional filtering by slot_id at the handler level
	if slotID != "" {
		filtered := make([]*models.ParkingEvent, 0, len(events))
		for _, ev := range events {
			if ev != nil && ev.SlotID == slotID {
				filtered = append(filtered, ev)
			}
		}
		events = filtered
	}

	// Respond JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(events); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
