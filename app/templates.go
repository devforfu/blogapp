package app

import (
    "fmt"
    log "github.com/sirupsen/logrus"
    "gopkg.in/russross/blackfriday.v2"
    "io/ioutil"
    "path/filepath"
)

func GetPageContent(name string) (string, error) {
    path := getPageFileContent(name)
    log.Debugf("getting blog page: %s", path)
    data, err := ioutil.ReadFile(path)
    if err != nil { return "", err }
    rendered := blackfriday.Run(data)
    return string(rendered), nil
}

func getPageFileContent(name string) string {
    return filepath.Join(ServerConfig.PagesRoot, fmt.Sprintf("%s.md", name))
}

func getTemplateFileContent(name string) string {
    return filepath.Join(ServerConfig.TemplatesRoot, fmt.Sprintf("%s.html", name))
}