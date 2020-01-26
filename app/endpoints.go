package app

import (
    "blogapp/app/blog"
    util "fastgoing"
    "fmt"
    "github.com/gorilla/mux"
    log "github.com/sirupsen/logrus"
    "html/template"
    "net/http"
    "strconv"
)

func Home(w http.ResponseWriter, req *http.Request) {
    mainPath := getTemplateFileContent("main")
    homePath := getTemplateFileContent("home")
    t, _ := template.ParseFiles(mainPath, homePath)
    util.Check(t.ExecuteTemplate(w, "main", Assets))
}

func BlogPage(w http.ResponseWriter, req *http.Request) {
    var notFoundIfError = func(err error) {
        if err != nil {
            log.Debugf("cannot resolve the template: %s", err.Error())
            http.NotFound(w, req)
        }
    }
    ref, err := parseReference(req)
    notFoundIfError(err)
    post, err := blog.NewPost(ref)
    notFoundIfError(err)
    _, _ = post.Write(w)
}

//func BlogPage(w http.ResponseWriter, req *http.Request) {
//    log.Debugf("got URL request: %s", req.URL.Path)
//    ref, err := parseReference(req)
//    if err != nil {
//        log.Debugf("cannot resolve the template: %s", err.Error())
//        http.NotFound(w, req)
//    } else {
//        content, err := GetPageContent(ref.GetPageName())
//        if err != nil {
//            http.NotFound(w, req)
//        }
//        _, _ = fmt.Fprint(w, content)
//    }
//}

func parseReference(r *http.Request) (ref *blog.PostReference, err error) {
    params := mux.Vars(r)

    year, err := strconv.ParseInt(params["year"], 10, 32)
    if err != nil { return }

    month, err := strconv.ParseInt(params["month"], 10, 32)
    if err != nil { return }

    day, err := strconv.ParseInt(params["day"], 10, 32)
    if err != nil { return }

    name, ok := params["post"]
    if !ok {
        return nil, fmt.Errorf("post parameter is not found")
    } else {
        ref = &blog.PostReference{
            Year:int(year),
            Month:int(month),
            Day:int(day),
            Name:name}
        return ref, nil
    }
}
