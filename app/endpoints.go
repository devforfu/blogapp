package app

import (
    "fmt"
    "net/http"
)

func Demo(w http.ResponseWriter, req *http.Request) {
    content := GetPage("demo")
    _, _ = fmt.Fprint(w, content)
}