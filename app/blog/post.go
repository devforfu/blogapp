package blog

import (
    "blogapp/app/config"
    "encoding/json"
    "fmt"
    util "github.com/devforfu/fastgoing"
    "gopkg.in/russross/blackfriday.v2"
    "io"
    "io/ioutil"
    "path/filepath"
    "regexp"
    "strings"
)

// PostReference keeps information that helps map URL suffix to post file name.
type PostReference struct {
    Year, Month, Day int
    Name string
}

// Filename converts publication date and post name into a name of Markdown
// file with post content.
func (ref *PostReference) Filename() string {
    filename := fmt.Sprintf("%d_%d_%d_%s.md", ref.Year, ref.Month, ref.Day, ref.Name)
    filename = strings.ReplaceAll(strings.ToLower(filename), "-", "_")
    return filename
}

// Post represent a single post content.
type Post struct {
    Preamble *PostPreamble
    RenderedPage string
}

func NewPost(ref *PostReference) (*Post, error) {
    path := filepath.Join(config.ServerConfig.PagesRoot, ref.Filename())
    markdownContent, err := ioutil.ReadFile(path)
    if err != nil { return nil, err }
    preamble, post, err := extractPreamble(string(markdownContent))
    if err != nil { return nil, err }
    rendered := blackfriday.Run([]byte(post))
    return &Post{Preamble:preamble, RenderedPage:string(rendered)}, nil
}

func (p *Post) Write(w io.Writer) (int, error) {
    return fmt.Fprint(w, p.RenderedPage)
}

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