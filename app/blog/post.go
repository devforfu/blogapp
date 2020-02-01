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
    "strings"
    "time"
)

// PostReference keeps information that helps map URL suffix to post file name.
type PostReference struct {
    Year, Month, Day int
    Name string
}

func (ref *PostReference) URL() string {
    url := fmt.Sprintf("/post/%d/%02d/%02d/%s", ref.Year, ref.Month, ref.Day, ref.Name)
    return strings.ToLower(url)
}

// Filename converts publication date and post name into a name of Markdown
// file with post content.
func (ref *PostReference) Filename() string {
    filename := fmt.Sprintf("%d-%02d-%02d-%s.md", ref.Year, ref.Month, ref.Day, ref.Name)
    return strings.ToLower(filename)
}

// Post represent a single post content.
type Post struct {
    Preamble *PostPreamble
    PublicationDate time.Time
    RenderedPage template.HTML
    URL string
    IsForeign bool
    Logo string
}

func NewPost(ref *PostReference) (*Post, error) {
    path := filepath.Join(config.ServerConfig.PagesRoot, ref.Filename())

    markdownContent, err := ioutil.ReadFile(path)
    if err != nil { return nil, err }

    preamble, pageContent, err := extractPreamble(string(markdownContent))
    if err != nil { return nil, err }

    rendered := blackfriday.Run([]byte(pageContent))
    published := util.DateUTC(ref.Year, ref.Month, ref.Day)

    var logo string
    var isForeign bool
    if match := config.RegexURL.Search(ref.URL()); len(match) > 0 {
        logo = match["origin"]
        isForeign = true
    } else {
        logo = ""
        isForeign = false
    }

    post := &Post{
        Preamble:preamble,
        RenderedPage:template.HTML(rendered),
        PublicationDate:published,
        URL:ref.URL(),
        IsForeign:isForeign,
        Logo:logo}

    return post, nil
}

func (p *Post) RenderWith(baseTemplateName string, w io.Writer) {
    path := config.ServerConfig.GetTemplateFilePath(baseTemplateName)
    t := template.Must(template.ParseFiles(path))
    wrappedPage := fmt.Sprintf(config.FormatWrappedPostContent, p.Preamble.Title, p.RenderedPage)
    t = template.Must(t.Parse(wrappedPage))
    data := map[string]interface{}{"Assets": config.DefaultAssets}
    util.Check(t.ExecuteTemplate(w, baseTemplateName, data))
}

func (p *Post) Digest() template.HTML {
    index := strings.Index(string(p.RenderedPage), config.ServerConfig.PostDigestSeparator)
    if index == -1 { return "" }
    digest := p.RenderedPage[:index]
    return template.HTML(digest)
}

func (p *Post) VerbosePublicationDate() string {
    return p.PublicationDate.Format(config.FormatVerboseDate)
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
    matched := config.RegexJSONPreamble.Search(jsonPreambleOnly)["preamble"]

    var preamble PostPreamble
    util.Check(json.Unmarshal([]byte(matched), &preamble))
    trimmed := strings.Trim(strings.ReplaceAll(postContentOnly, sep, ""), "\n\t ")
    return &preamble, trimmed, nil
}
