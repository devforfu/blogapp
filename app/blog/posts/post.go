package posts

import (
    "encoding/json"
    "fmt"
    "github.com/devforfu/blogapp/app/config"
    util "github.com/devforfu/fastgoing"
    log "github.com/sirupsen/logrus"
    "gopkg.in/russross/blackfriday.v2"
    "html/template"
    "io/ioutil"
    "os"
    "path/filepath"
    "sort"
    "strings"
    "time"
)

type Post struct {
    Year, Month, Day int
    Name string
    Meta Meta
    rawDigest string
    rawContent string
}

type Meta struct {
    Category string   `json:"category"`
    Title string      `json:"title"`
    Tags []string     `json:"tags"`
    ImageName string  `json:"image"`
    ForeignURL string `json:"foreign_url"`
    Identifier int    `json:"identifier"`
}

func FromMarkdown(filePath string) (*Post, error) {
    if match := config.RegexMDFile.Search(filePath); len(match) > 0 {
        post := Post{
            Year:         util.MustInt(match["year"]),
            Month:        util.MustInt(match["month"]),
            Day:          util.MustInt(match["day"]),
            Name:         match["name"],
        }
        err := post.Parse(filePath)
        if err != nil {
            return nil, err
        } else {
            return &post, nil
        }
    } else {
        return nil, fmt.Errorf("failed to parse file: %s", filePath)
    }
}

// Parse reads markdown content and populates posts's properties required
// to render a template.
func (p *Post) Parse(filePath string) error {
    markdown, err := ioutil.ReadFile(filePath)

    if err != nil { return err }

    err = p.parseFileContent(
        string(markdown),
        config.ServerConfig.PostPreambleSeparator,
        config.ServerConfig.PostDigestSeparator)

    if err != nil { return err }

    return nil
}

func (p *Post) Digest() template.HTML {
    return template.HTML(blackfriday.Run([]byte(p.rawDigest)))
}

func (p *Post) Content() template.HTML {
    return template.HTML(blackfriday.Run([]byte(p.rawContent)))
}

func (p *Post) IsForeign() bool {
    return p.Meta.ForeignURL != ""
}

func (p *Post) URL() string {
    if p.IsForeign() {
        return p.Meta.ForeignURL
    } else {
        url := fmt.Sprintf("/posts/%d/%02d/%02d/%s", p.Year, p.Month, p.Day, p.Name)
        return strings.ToLower(url)
    }
}

func (p *Post) Logo() string {
    if p.IsForeign() {
        origin := config.RegexURL.Search(p.URL())["origin"]
        return strings.Title(origin)
    } else {
        return ""
    }
}

func (p *Post) PublicationDate() time.Time {
    return util.DateUTC(p.Year, p.Month, p.Day)
}

func (p *Post) VerbosePublicationDate() string {
    return p.PublicationDate().Format(config.FormatVerboseDate)
}

func (p *Post) parseFileContent(content, preambleSep, digestSep string) error {
    trimmed, err := p.parseMeta(content, preambleSep)
    if err != nil { return err }

    trimmed, err = p.parseDigest(trimmed, digestSep)
    if err != nil { return err }

    p.rawContent = trimmed

    return nil
}

func (p *Post) parseMeta(content, preambleSep string) (string, error) {
    index := strings.Index(content, preambleSep)
    if index == -1 {
        return "", fmt.Errorf("preamble is not found")
    }

    jsonPreamble := content[:index]
    trimmed := content[index + len(preambleSep):]
    matched := config.RegexJSONPreamble.Search(jsonPreamble)["preamble"]

    var meta Meta
    err := json.Unmarshal([]byte(matched), &meta)
    if err != nil {
        return "", err
    }

    p.Meta = meta

    return trimmed, nil
}

func (p *Post) parseDigest(content, digestSep string) (string, error) {
    index := strings.Index(content, digestSep)
    if index == -1 {
        return "", fmt.Errorf("digest is not found")
    }

    p.rawDigest = content[:index]
    trimmed := content[index + len(digestSep):]

    return trimmed, nil
}

type List []*Post
func (arr List) Len() int      { return len(arr) }
func (arr List) Less(i, j int) bool { return arr[i].PublicationDate().Unix() < arr[j].PublicationDate().Unix() }
func (arr List) Swap(i, j int) { arr[i], arr[j] = arr[j], arr[i] }

func FetchFromFolder(dirName string) List {
    posts := make(List, 0)

    var parseFiles = func(path string, info os.FileInfo, err error) error {
        if err != nil {
            log.Debugf("failed to parse dir: %s", err)
        } else {
            if match := config.RegexMDFile.Search(path); len(match) > 0 {
                post, err := FromMarkdown(path)
                if err != nil {
                    log.Warnf("failed to create a posts from file: %s", err)
                } else {
                    posts = append(posts, post)
                }
            }
        }
        return nil
    }

    if err := filepath.Walk(dirName, parseFiles); err != nil {
        log.Errorf("failed to parse posts pages: %s", err)
    }

    sort.Sort(sort.Reverse(posts))

    return posts
}