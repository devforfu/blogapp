package app

import (
    "blogapp/app/blog"
    "blogapp/app/config"
    "fmt"
    util "github.com/devforfu/fastgoing"
    "github.com/gorilla/mux"
    log "github.com/sirupsen/logrus"
    "html/template"
    "net/http"
    "strconv"
)

func Home(w http.ResponseWriter, req *http.Request) {
    mainPath := config.ServerConfig.GetTemplateFilePath("main")
    homePath := config.ServerConfig.GetTemplateFilePath("home")
    t, _ := template.ParseFiles(mainPath, homePath)
    util.Check(t.ExecuteTemplate(w, "main", config.Assets))
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
    path := config.ServerConfig.GetTemplateFilePath("main")
    t, err := template.ParseFiles(path)
    notFoundIfError(err)
    content := fmt.Sprintf(`
{{ define "title" }}%s{{ end }}
{{ define "content" }}
%s
{{ end }}`, post.Preamble.Title, post.RenderedPage)
    t, err = t.Parse(content)
    notFoundIfError(err)
    err = t.ExecuteTemplate(w, "main", config.Assets)

    if err != nil {
        log.Debugf("failed to execute the template: %s", err)
    }
}

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
