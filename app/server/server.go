package server

import (
    "github.com/devforfu/blogapp/app/config"
    "github.com/gorilla/mux"
    "net/http"
)

func New() *http.Server {
    router := mux.NewRouter()
    fs := http.FileServer(RestrictedFileSystem{http.Dir(config.ServerConfig.StaticRoot)})
    router.HandleFunc("/", Posts)
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
    router.HandleFunc(`/posts/{year:[0-9]+}/{month:[0-9]+}/{day:[0-9]+}/{posts:[a-zA-Z0-9\-]+}`, Article)
    router.HandleFunc("/posts", Posts)
    server := &http.Server{Addr:"0.0.0.0:9090", Handler:router}
    return server
}
