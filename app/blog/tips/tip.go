package tips

import (
    "fmt"
    "gopkg.in/russross/blackfriday.v2"
    "html/template"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

type Tip struct {
    Order int
    Title string
    Content template.HTML
}

func LoadTip(order int, folder string) (tip *Tip, err error) {
    var matched string

    err = filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
        if err != nil { return err }
        prefix := fmt.Sprintf("%02d", order)
        if strings.HasPrefix(filepath.Base(path), prefix) {
            matched = path
        }
        return nil
    })

    if err != nil { return }
    if matched == "" { return nil, fmt.Errorf("tip #%d is not found: ", order)}

    markdown, err := ioutil.ReadFile(matched)

    if err != nil { return }

    parts := strings.Split(strings.ReplaceAll(matched, ".md", ""), "_")
    title := strings.Title(strings.Join(parts[1:], " "))
    content := template.HTML(blackfriday.Run(markdown))
    tip = &Tip{Order: order, Title: title, Content: content}

    return
}