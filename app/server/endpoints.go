package server

import (
    "fmt"
    "github.com/devforfu/blogapp/app/blog/posts"
    "github.com/devforfu/blogapp/app/blog/tips"
    "github.com/devforfu/blogapp/app/config"
    "github.com/gorilla/mux"
    log "github.com/sirupsen/logrus"
    "html/template"
    "net/http"
    "path/filepath"
    "strconv"
    "strings"
)

func Posts(w http.ResponseWriter, req *http.Request) {
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
}

func TipsList(w http.ResponseWriter, req *http.Request) {
    allTips, err := tips.FetchFromFolder(config.ServerConfig.TipsRoot)

    if err != nil {
        log.Debugf("failed to load tips: %s", err)
        http.NotFound(w, req)
        return
    }

    t := parseTemplates("tips", "main")
    err = t.ExecuteTemplate(w,"main", map[string]interface{}{
        "Assets": config.ServerConfig.Assets,
        "Tips": allTips,
    })

    if err != nil {
        log.Debugf("failedto to render the template: %s", err)
        http.NotFound(w, req)
        return
    }
}

func Tip(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    order, err := strconv.ParseInt(params["order"], 10, 32)
    if err != nil {
        log.Debugf("wrong tip number")
        http.NotFound(w, req)
    } else {
        tip, err := tips.LoadTip(int(order), config.ServerConfig.TipsRoot)
        if err != nil || tip == nil {
            log.Debugf("failed to load tip: %s", err)
            http.NotFound(w, req)
        } else {
            t := parseTemplates("tip", "main")
            err = t.ExecuteTemplate(w, "main", map[string]interface{}{
                "Assets": config.ServerConfig.Assets,
                "Tip":    tip,
            })
            if err != nil {
                log.Debugf("failed to render a tip: %s", err)
                http.NotFound(w, req)
            }
        }
    }
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

func parseTemplates(names ...string) *template.Template {
    filePaths := make([]string, 0)
    for _, name := range names {
        filePaths = append(filePaths, config.ServerConfig.GetTemplateFilePath(name))
    }
    return template.Must(template.ParseFiles(filePaths...))
}
