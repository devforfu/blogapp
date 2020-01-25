package app

import (
    util "fastgoing"
    "fmt"
    log "github.com/sirupsen/logrus"
    "gopkg.in/russross/blackfriday.v2"
    "html/template"
    "io/ioutil"
)

func GetPage(name string) (string, error) {
    path := getPageFileContent(name)
    log.Debugf("getting blog page: %s", path)
    data, err := ioutil.ReadFile(path)
    if err != nil { return "", err }
    rendered := blackfriday.Run(data)
    return string(rendered), nil
}

func GetTemplate(name string) (*template.Template, error) {
    path := getTemplateFileContent(name)
    log.Debugf("getting template: %s", path)
    return template.ParseFiles(path)
}

func getPageFileContent(name string) string {
    return util.JoinPaths(ServerConfig.PagesRoot, fmt.Sprintf("%s.md", name))
}

func getTemplateFileContent(name string) string {
    return util.JoinPaths(ServerConfig.TemplatesRoot, fmt.Sprintf("%s.html", name))
}