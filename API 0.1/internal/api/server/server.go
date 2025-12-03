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
    // POST /data
    mux.Handle("/data", method("POST", dataHandlers.NewPostHandler(dataSvc)))
    // PUT /data/{id}
    mux.Handle("/data/", method("PUT", dataHandlers.NewPutHandler(dataSvc)))

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
    // Add graceful shutdown logic if you expand beyond ServeMux
    return nil
}

// method wraps a handler and enforces HTTP method, returning 405 for others.
func method(allowed string, h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != allowed {
            http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
            return
        }
        h.ServeHTTP(w, r)
    })
}
