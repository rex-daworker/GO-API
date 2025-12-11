package parking

import (
    "encoding/json"
    "log"
    "net/http"

    "goapi/internal/api/repository/models"
    "goapi/internal/api/service/parking"
)

type PostHandler struct {
    Service parking.ParkingService
}

func NewPostHandler(service parking.ParkingService) *PostHandler {
    return &PostHandler{Service: service}
}

func (h *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var ev models.ParkingEvent

    if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
        http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
        log.Println("JSON decode error:", err)
        return
    }

    if err := h.Service.Create(&ev, r.Context()); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        log.Println("Create error:", err)
        return
    }

    w.WriteHeader(http.StatusCreated)
    _, _ = w.Write([]byte("Parking event created"))
}
