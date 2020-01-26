package app

import (
    util "fastgoing"
    "fmt"
    "github.com/gorilla/mux"
    log "github.com/sirupsen/logrus"
    "html/template"
    "net/http"
    "strconv"
    "strings"
)

func Home(w http.ResponseWriter, req *http.Request) {
    mainPath := getTemplateFileContent("main")
    homePath := getTemplateFileContent("home")
    t, _ := template.ParseFiles(mainPath, homePath)
    util.Check(t.ExecuteTemplate(w, "main", Assets))
}

func BlogPage(w http.ResponseWriter, req *http.Request) {
    log.Debugf("got URL request: %s", req.URL.Path)
    ref, err := parseReference(req)
    if err != nil {
        log.Debugf("cannot resolve the template: %s", err.Error())
        http.NotFound(w, req)
    } else {
        content, err := GetPageContent(ref.GetPageName())
        if err != nil {
            http.NotFound(w, req)
        }
        _, _ = fmt.Fprint(w, content)
    }
}

type PageReference struct {
    Year, Month, Day int
    Name string
}

func (ref *PageReference) GetPageName() string {
    return normalize(fmt.Sprintf("%d_%d_%d_%s", ref.Year, ref.Month, ref.Day, ref.Name))
}

func parseReference(r *http.Request) (ref *PageReference, err error) {
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
        ref = &PageReference{
            Year:int(year),
            Month:int(month),
            Day:int(day),
            Name:name}
        return ref, nil
    }
}

func normalize(name string) string {
    return strings.ReplaceAll(strings.ToLower(name), "-", "_")
}