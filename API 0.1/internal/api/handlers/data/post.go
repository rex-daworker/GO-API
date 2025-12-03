package data

import (
    "encoding/json"
    "goapi/internal/api/repository/models"
    serviceData "goapi/internal/api/service/data"
    "log"
    "net/http"
)

type PostHandler struct {
    Service serviceData.DataService
}

func NewPostHandler(service serviceData.DataService) *PostHandler {
    return &PostHandler{Service: service}
}

func (h *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var payload models.Data

    // Decode JSON body
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
        log.Println("JSON decode error:", err)
        return
    }
    log.Printf("Received POST payload: %+v\n", payload)

    // Create via service (returns validation errors)
    if err := h.Service.Create(&payload, r.Context()); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        log.Println("Create error:", err)
        return
    }

    w.WriteHeader(http.StatusCreated)
    _, _ = w.Write([]byte("Data created successfully"))
}
