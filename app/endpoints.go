package app

import (
    "blogapp/app/blog"
    "blogapp/app/cache"
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
    var notFoundOnError = func(err error) bool {
        if err != nil {
            log.Debugf("cannot resolve the template: %s", err.Error())
            http.NotFound(w, req)
            return true
        } else {
            return false
        }
    }

    var getPost = func(ref *blog.PostReference) *blog.Post {
       key := ref.Filename()
       post := cache.Default.Get(key)
       if post != nil { return post }
       post, err := blog.NewPost(ref)
       if notFoundOnError(err) { return nil }
       cache.Default.Set(key, post)
       return post
    }

    ref, err := parseReference(req)
    if notFoundOnError(err) { return }
    //post, err := blog.NewPost(ref)
    //if notFoundOnError(err) { return }
    if post := getPost(ref); post != nil {
        path := config.ServerConfig.GetTemplateFilePath("main")
        t, err := template.ParseFiles(path)
        if notFoundOnError(err) {
            return
        }
        content := fmt.Sprintf(`
{{ define "title" }}%s{{ end }}
{{ define "content" }}
%s
{{ end }}`, post.Preamble.Title, post.RenderedPage)
        t, err = t.Parse(content)
        if notFoundOnError(err) {
            return
        }
        err = t.ExecuteTemplate(w, "main", config.Assets)

        if err != nil {
            log.Debugf("failed to execute the template: %s", err)
        }
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
