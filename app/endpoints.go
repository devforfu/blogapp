package app

import (
    util "fastgoing"
    "fmt"
    "github.com/gorilla/mux"
    log "github.com/sirupsen/logrus"
    "net/http"
    "strings"
)

func Demo(w http.ResponseWriter, req *http.Request) {
    content, _ := GetPage("demo")
    _, _ = fmt.Fprint(w, content)
}

func Home(w http.ResponseWriter, req *http.Request) {
    t, err := GetTemplate("index")
    util.Check(err)
    util.Check(t.Execute(w, Assets))
}

func BlogPage(w http.ResponseWriter, req *http.Request) {
    log.Debugf("got URL request: %s", req.URL.Path)
    vars := mux.Vars(req)
    postName, ok := vars["post"]
    if ok {
        content, err := GetPage(normalize(postName))
        if err == nil {
            _, _ = fmt.Fprint(w, content)
            return
        }
    }
    http.NotFound(w, req)
}

func normalize(name string) string {
    return strings.ReplaceAll(strings.ToLower(name), "-", "_")
}