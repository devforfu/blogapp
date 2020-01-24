package app

import (
    "fmt"
    "github.com/golang/glog"
    "gopkg.in/russross/blackfriday.v2"
    "io/ioutil"
    "os"
    "strings"
)

var templatesFolder = getRootTemplatesFolder()

func GetPage(name string) string {
    path := getTemplatePath(name)
    glog.Info(fmt.Sprintf("getting template: %s", path))
    data, err := ioutil.ReadFile(path)
    checkError(err)
    rendered := blackfriday.Run(data)
    return string(rendered)
}

func getTemplatePath(name string) string {
    withExt := fmt.Sprintf("%s.md", name)
    return strings.Join([]string{templatesFolder, withExt}, "/")
}

func getRootTemplatesFolder() string {
    if root := os.Getenv("APP_TEMPLATES_ROOT"); root == "" {
        cwd, err := os.Getwd()
        checkError(err)
        return strings.Join([]string{cwd, "pages"}, "/")
    } else {
        return root
    }
}

func checkError(err error) {
    if err != nil { panic(err) }
}
