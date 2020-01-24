package app

import (
    "net/http"
)

func New() *http.Server {
    mux := http.NewServeMux()
    server := &http.Server{Addr:"0.0.0.0:9090", Handler:mux}
    mux.Handle("/demo", http.HandlerFunc(Demo))
    return server
}
