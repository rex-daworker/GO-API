package server

import (
    "context"
    dataHandlers "goapi/internal/api/handlers/data"
    "goapi/internal/api/service"
    "log"
    "net/http"
)

type Server struct {
    mux    *http.ServeMux
    ctx    context.Context
    logger *log.Logger
}

func NewServer(ctx context.Context, sf *service.ServiceFactory, logger *log.Logger) *Server {
    mux := http.NewServeMux()

    // Build data service
    dataSvc, err := sf.CreateDataService(service.SQLiteDataService)
    if err != nil {
        logger.Println("Error creating data service:", err)
    }

    // Routes
    mux.Handle("/data", dataHandlers.NewPostHandler(dataSvc))   // POST /data
    mux.Handle("/data/", dataHandlers.NewPutHandler(dataSvc))   // PUT /data/{id}

    return &Server{
        mux:    mux,
        ctx:    ctx,
        logger: logger,
    }
}

func (s *Server) ListenAndServe(addr string) error {
    return http.ListenAndServe(addr, s.mux)
}

func (s *Server) Shutdown() error {
    // Extend with context-aware shutdown if needed
    return nil
}
