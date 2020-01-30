package blog

import (
    "blogapp/app/config"
    "encoding/json"
    "fmt"
    util "github.com/devforfu/fastgoing"
    "gopkg.in/russross/blackfriday.v2"
    "html/template"
    "io"
    "io/ioutil"
    "path/filepath"
    "regexp"
    "strings"
    "time"
)

// PostReference keeps information that helps map URL suffix to post file name.
type PostReference struct {
    Year, Month, Day int
    Name string
}

// Filename converts publication date and post name into a name of Markdown
// file with post content.
func (ref *PostReference) Filename() string {
    filename := fmt.Sprintf("%d_%02d_%02d_%s.md", ref.Year, ref.Month, ref.Day, ref.Name)
    filename = strings.ReplaceAll(strings.ToLower(filename), "-", "_")
    return filename
}

// Post represent a single post content.
type Post struct {
    Preamble *PostPreamble
    PublicationDate time.Time
    RenderedPage string
}

func NewPost(ref *PostReference) (*Post, error) {
    path := filepath.Join(config.ServerConfig.PagesRoot, ref.Filename())
    markdownContent, err := ioutil.ReadFile(path)
    if err != nil { return nil, err }
    preamble, post, err := extractPreamble(string(markdownContent))
    if err != nil { return nil, err }
    rendered := blackfriday.Run([]byte(post))
    published := time.Date(ref.Year, time.Month(ref.Month), ref.Day, 0, 0,0,0, time.UTC)
    return &Post{Preamble:preamble, RenderedPage:string(rendered), PublicationDate:published}, nil
}

func (p *Post) RenderWith(baseTemplateName string, w io.Writer) {
    path := config.ServerConfig.GetTemplateFilePath(baseTemplateName)
    t := template.Must(template.ParseFiles(path))
    wrappedPage := fmt.Sprintf(wrappedContent, p.Preamble.Title, p.RenderedPage)
    t = template.Must(t.Parse(wrappedPage))
    util.Check(t.ExecuteTemplate(w, baseTemplateName, config.DefaultAssets))
}

func (p *Post) Digest() string {
    index := strings.Index(p.RenderedPage, config.ServerConfig.PostDigestSeparator)
    if index == -1 { return "" }
    digest := p.RenderedPage[:index]
    return digest
}

type PostsList []*Post
func (arr PostsList) Len() int           { return len(arr) }
func (arr PostsList) Less(i, j int) bool { return arr[i].PublicationDate.Unix() < arr[j].PublicationDate.Unix() }
func (arr PostsList) Swap(i, j int)      { arr[i], arr[j] = arr[j], arr[i] }

type PostPreamble struct {
    Category string  `json:"category"`
    Title string     `json:"title"`
    Tags []string    `json:"tags"`
    ImageName string `json:"image"`
    Identifier int   `json:"identifier"`
}

func extractPreamble(markdownContent string) (*PostPreamble, string, error) {
    sep := config.ServerConfig.PostPreambleSeparator
    index := strings.Index(markdownContent, sep)
    if index == -1 {
        return nil, "", fmt.Errorf("posts without preamble are not valid")
    }

    jsonPreambleOnly := markdownContent[:index]
    postContentOnly := markdownContent[index:]
    reg := regexp.MustCompile("^```json\n(?P<preamble>[\\w\\W]+)```")
    matched := reg.FindStringSubmatch(jsonPreambleOnly)

    var preamble PostPreamble
    util.Check(json.Unmarshal([]byte(matched[1]), &preamble))
    trimmed := strings.Trim(strings.ReplaceAll(postContentOnly, sep, ""), "\n\t ")
    return &preamble, trimmed, nil
}

const wrappedContent = `
{{ define "title" }}%s{{ end }}
{{ define "content" }}
%s
{{ end }}
`
