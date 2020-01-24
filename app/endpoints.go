package app

import (
    "fmt"
    "github.com/gorilla/mux"
    log "github.com/sirupsen/logrus"
    "net/http"
)

func Demo(w http.ResponseWriter, req *http.Request) {
    content := GetPage("demo")
    _, _ = fmt.Fprint(w, content)
}

func BlogPage(w http.ResponseWriter, req *http.Request) {
    log.Debugf("got URL request: %s", req.URL.Path)
    vars := mux.Vars(req)
    log.Debugf("%v", vars)
}