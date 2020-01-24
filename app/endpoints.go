package app

import (
    "fmt"
    "net/http"
)

func Demo(w http.ResponseWriter, req *http.Request) {
    content := GetPage("demo")
    _, _ = fmt.Fprint(w, content)
}

func BlogPage(w http.ResponseWriter, req *http.Request) {
    //glog.Infof("got URL request: %s", req.URL.Path)
    return
}