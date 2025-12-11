// internal/api/handlers/parking/put.go
package parking

import (
    "encoding/json"
    "github.com/rex-daworker/GO-API/internal/api/repository/models"
    svc "github.com/rex-daworker/GO-API/internal/api/service/parking"
    "log"
    "net/http"
    "strconv"
    "strings"
)

type PutHandler struct {
    Service svc.ParkingService
}

func NewPutHandler(service svc.ParkingService) *PutHandler {
    return &PutHandler{Service: service}
}

func (h *PutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var ev models.ParkingEvent

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
    ev.ID = id

    if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
        http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
        log.Println("JSON decode error:", err)
        return
    }

    if _, err := h.Service.Update(&ev, r.Context()); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        log.Println("Update error:", err)
        return
    }

    w.WriteHeader(http.StatusOK)
    _, _ = w.Write([]byte("Parking event updated"))
}
