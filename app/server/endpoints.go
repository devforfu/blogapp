package server

import (
    "fmt"
    "github.com/devforfu/blogapp/app/blog/posts"
    "github.com/devforfu/blogapp/app/config"
    "github.com/gorilla/mux"
    log "github.com/sirupsen/logrus"
    "html/template"
    "net/http"
    "path/filepath"
    "strconv"
    "strings"
)

//func Home(w http.ResponseWriter, req *http.Request) {
//    t := parseTemplates("main", "home")
//    data := map[string]interface{}{"Assets": config.ServerConfig.Assets}
//    util.Check(t.ExecuteTemplate(w, "main", data))
//}

func Posts(w http.ResponseWriter, req *http.Request) {
    //posts := blog.ListPosts()
    //sort.Sort(sort.Reverse(posts))
    //t := parseTemplates( "posts", "main")
    //data := map[string]interface{} {
    //    "Assets":config.ServerConfig.Assets,
    //    "Posts":posts,
    //}
    //util.Check(t.ExecuteTemplate(w, "main", data))

    allPosts := posts.FetchFromFolder(config.ServerConfig.PagesRoot)
    t := parseTemplates("posts", "main")
    err := t.ExecuteTemplate(w, "main", map[string]interface{} {
        "Assets": config.ServerConfig.Assets,
        "Posts": allPosts,
    })
    if err != nil {
        log.Debugf("failed to render posts list: %s", err)
        http.NotFound(w, req)
    }
}

func Article(w http.ResponseWriter, req *http.Request) {
    post, err := getPostFromRequest(req)
    if err != nil {
        log.Debugf("cannot resolve the template: %s", err)
        http.NotFound(w, req)
    } else {
        t := parseTemplates("article", "main")
        err = t.ExecuteTemplate(w, "main", map[string]interface{} {
            "Assets": config.ServerConfig.Assets,
            "Post": post,
        })
        if err != nil {
            log.Debugf("failed to render a post: %s", err)
            http.NotFound(w, req)
        }
    }

    //var notFoundOnError = func(err error) bool {
    //    if err != nil {
    //        log.Debugf("cannot resolve the template: %s", err.Error())
    //        http.NotFound(w, req)
    //        return true
    //    } else {
    //        return false
    //    }
    //}
    //
    //var getPost = func(ref *blog.PostReference) *blog.Post {
    //    key := ref.Filename()
    //    post := cache.Default.Get(key)
    //    if post != nil { return post }
    //    post, err := blog.NewPost(ref)
    //    if notFoundOnError(err) { return nil }
    //    cache.Default.Set(key, post)
    //    return post
    //}
    //
    //ref, err := parseReference(req)
    //if notFoundOnError(err) { return }
    //if post := getPost(ref); post != nil {
    //    post.RenderWith("main", w)
    //}


}

func getPostFromRequest(r *http.Request) (post *posts.Post, err error) {
    params := mux.Vars(r)

    year, err := strconv.ParseInt(params["year"], 10, 32)
    if err != nil { return }

    month, err := strconv.ParseInt(params["month"], 10, 32)
    if err != nil { return }

    day, err := strconv.ParseInt(params["day"], 10, 32)
    if err != nil { return }

    if name, ok := params["posts"]; ok {
        fileName := strings.ToLower(
            fmt.Sprintf("%d-%02d-%02d-%s.md", year, month, day, name))
        filePath := filepath.Join(config.ServerConfig.PagesRoot, fileName)
        return posts.FromMarkdown(filePath)
    } else {
        return nil, fmt.Errorf("'posts' parameter is not found")
    }
}

//func parseReference(r *http.Request) (ref *blog.PostReference, err error) {
//    params := mux.Vars(r)
//
//    year, err := strconv.ParseInt(params["year"], 10, 32)
//    if err != nil { return }
//
//    month, err := strconv.ParseInt(params["month"], 10, 32)
//    if err != nil { return }
//
//    day, err := strconv.ParseInt(params["day"], 10, 32)
//    if err != nil { return }
//
//    name, ok := params["posts"]
//    if !ok {
//        return nil, fmt.Errorf("posts parameter is not found")
//    } else {
//        ref = &blog.PostReference{
//            Year:int(year),
//            Month:int(month),
//            Day:int(day),
//            Name:name}
//        return ref, nil
//    }
//}

func parseTemplates(names ...string) *template.Template {
    filePaths := make([]string, 0)
    for _, name := range names {
        filePaths = append(filePaths, config.ServerConfig.GetTemplateFilePath(name))
    }
    return template.Must(template.ParseFiles(filePaths...))
}
