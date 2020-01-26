package blog

import (
    "encoding/json"
    util "fastgoing"
    "fmt"
    "gopkg.in/russross/blackfriday.v2"
    "io"
    "io/ioutil"
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
    markdownContent, err := ioutil.ReadFile(ref.Filename())
    if err != nil { return nil, err }
    preamble, post := ExtractPreamble(string(markdownContent))
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

func ExtractPreamble(markdownContent string) (*PostPreamble, string) {
    reg := regexp.MustCompile("^```json\n[\\d\\D]+```")
    jsonPreamble := reg.FindString(markdownContent)
    lines := strings.Split(jsonPreamble, "\n")
    jsonOnly := strings.Join(lines[1:len(lines)-1], "\n")
    var preamble PostPreamble
    util.Check(json.Unmarshal([]byte(jsonOnly), &preamble))
    postOnly := reg.ReplaceAllString(markdownContent, "")
    trimmed := strings.Trim(postOnly, " \t\n")
    return &preamble, trimmed
}