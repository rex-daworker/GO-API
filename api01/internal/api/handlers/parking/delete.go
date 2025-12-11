package parking

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"goapi/internal/api/repository/models"
	"goapi/internal/api/service/parking"
)

type DeleteHandler struct {
	Service parking.ParkingService
}

func NewDeleteHandler(service parking.ParkingService) *DeleteHandler {
	return &DeleteHandler{Service: service}
}

func (h *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	ev := &models.ParkingEvent{ID: id}
	if _, err := h.Service.Delete(ev, r.Context()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Delete error:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Parking event deleted"))
}
