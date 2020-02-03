package server

import (
    "github.com/gorilla/mux"
    "net/http"
)

func New() *http.Server {
    router := mux.NewRouter()
    fs := http.FileServer(RestrictedFileSystem{http.Dir("static")})
    router.HandleFunc("/", Posts)
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
    router.HandleFunc(`/post/{year:[0-9]+}/{month:[0-9]+}/{day:[0-9]+}/{post:[a-zA-Z0-9\-]+}`, BlogPage)
    router.HandleFunc("/posts", Posts)
    server := &http.Server{Addr:"0.0.0.0:9090", Handler:router}
    return server
}
