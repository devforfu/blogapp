package app

import (
    "fmt"
    log "github.com/sirupsen/logrus"
    "gopkg.in/russross/blackfriday.v2"
    "io/ioutil"
    "strings"
)

func GetPage(name string) (string, error) {
    path := getTemplatePath(name)
    log.Debugf("getting template: %s", path)
    data, err := ioutil.ReadFile(path)
    if err != nil { return "", err }
    rendered := blackfriday.Run(data)
    return string(rendered), nil
}

func getTemplatePath(name string) string {
    withExt := fmt.Sprintf("%s.md", name)
    return strings.Join([]string{ServerConfig.TemplatesRoot, "pages", withExt}, "/")
}