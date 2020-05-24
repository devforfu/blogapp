package server

import (
    "github.com/gorilla/mux"
    log "github.com/sirupsen/logrus"
    "net/http"
)

// LoggingMiddleware logs all requests to the API.
func LoggingMiddleware(_ *mux.Router) mux.MiddlewareFunc {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
            clientAddr := req.Header.Get("X-Forwarded-For")
            if clientAddr == "" {
                clientAddr = req.RemoteAddr
            }
            log.Debugf("addr=%s | endpoint=%s", clientAddr, req.URL)
            next.ServeHTTP(w, req)
        })
    }
}
