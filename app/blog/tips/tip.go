package tips

import (
    "fmt"
    log "github.com/sirupsen/logrus"
    "gopkg.in/russross/blackfriday.v2"
    "html/template"
    "io/ioutil"
    "path/filepath"
    "sort"
    "strconv"
    "strings"
)

type Tip struct {
    Order int
    Title string
    FileName string
    Content template.HTML
}

func (t *Tip) IsLoaded() bool {
    return t.Content != ""
}

func (t *Tip) Load() error {
    markdown, err := ioutil.ReadFile(t.FileName)
    if err != nil { return err }
    t.Content = template.HTML(blackfriday.Run(markdown))
    return nil
}

func (t *Tip) EnumeratedTitle() string {
    return fmt.Sprintf("Tip #%d: %s", t.Order, t.Title)
}

func LoadTip(order int, folder string) (tip *Tip, err error) {
    tips, err := FetchFromFolder(folder)
    if err != nil {
        return nil, err
    }
    if order > len(tips) || order < 1 {
        return nil, fmt.Errorf("invalid index: %d", order)
    }
    chosenTip := tips[order-1]
    err = chosenTip.Load()
    if err != nil {
        return nil, err
    }
    return chosenTip, nil
}

type List []*Tip
func (arr List) Len() int           { return len(arr) }
func (arr List) Less(i, j int) bool { return arr[i].Order < arr[j].Order }
func (arr List) Swap(i, j int)      { arr[i], arr[j] = arr[j], arr[i] }

func FetchFromFolder(dirName string) (List, error) {
    files, err := ioutil.ReadDir(dirName)

    if err != nil {
        log.Warnf("failed to parse folder with tips: %s", dirName)
        return nil, err
    }

    tips := make(List, 0)
    for _, file := range files {
        basename := file.Name()
        name := strings.TrimSuffix(basename, filepath.Ext(basename))
        parts := strings.Split(name, "_")
        order, err := strconv.ParseInt(parts[0], 10, 32)
        if err != nil {
            log.Warnf("file name should start with a number: %s", basename)
        } else {
            tips = append(tips, &Tip{
                Order: int(order),
                Title: strings.Title(strings.Join(parts[1:], " ")),
                FileName: filepath.Join(dirName, basename),
            })
        }
    }

    sort.Sort(tips)

    return tips, nil
}