package app

import (
    "github.com/gorilla/mux"
    "net/http"
)

func New() *http.Server {
    router := mux.NewRouter()
    fs := http.FileServer(http.Dir("static"))
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
    router.HandleFunc("/", Home)
    router.HandleFunc("/demo", Demo)
    router.HandleFunc(`/posts/{post:[a-zA-Z0-9\-]+}`, BlogPage)
    server := &http.Server{Addr:"0.0.0.0:9090", Handler:router}
    return server
}
