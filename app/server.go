package app

import (
    "github.com/gorilla/mux"
    "net/http"
)

func New() *http.Server {
    router := mux.NewRouter()
    router.Handle("/demo", http.HandlerFunc(Demo))
    router.Handle(`/posts/{post:[a-zA-Z0-9\-]+}`, http.HandlerFunc(BlogPage))
    server := &http.Server{Addr:"0.0.0.0:9090", Handler:router}
    return server
}
