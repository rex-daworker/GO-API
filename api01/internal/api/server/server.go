// internal/api/server/server.go
package server

import (
    "context"
    "log"
    "net/http"

    "goapi/internal/api/handlers/parking"
    "goapi/internal/api/service"
)


type Server struct {
    mux    *http.ServeMux
    ctx    context.Context
    logger *log.Logger
}

func NewServer(ctx context.Context, sf *service.ServiceFactory, logger *log.Logger) *Server {
    mux := http.NewServeMux()

    parkingSvc, err := sf.CreateParkingService(service.SQLiteParkingService)
    if err != nil {
        logger.Println("Error creating parking service:", err)
    }

    // Protected routes
    mux.Handle("/parking", withAuth(method("POST", ph.NewPostHandler(parkingSvc))))
    mux.Handle("/parking/", withAuth(method("PUT", ph.NewPutHandler(parkingSvc))))

    return &Server{
        mux:    mux,
        ctx:    ctx,
        logger: logger,
    }
}

func (s *Server) ListenAndServe(addr string) error {
    return http.ListenAndServe(addr, s.mux)
}

func (s *Server) Shutdown() error { return nil }

func method(allowed string, h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != allowed {
            http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
            return
        }
        h.ServeHTTP(w, r)
    })
}

// Simple Basic Auth middleware
func withAuth(h http.Handler) http.Handler {
    const user = "admin_oghenerobo" // CHANGE HERE
    const pass = "StrongPass!2025"  // CHANGE HERE

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        u, p, ok := r.BasicAuth()
        if !ok || u != user || p != pass {
            w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        h.ServeHTTP(w, r)
    })
}
