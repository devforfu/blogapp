package blog

import (
    "fmt"
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
    filename := fmt.Sprintf("%d_%d_%d_%s", ref.Year, ref.Month, ref.Day, ref.Name)
    filename = strings.ReplaceAll(strings.ToLower(filename), "-", "_")
    return filename
}

// Post represent a single post content.
type Post struct {
    Preamble PostPreamble
}

type PostPreamble struct {
    Category string
    Link string
    Title string
    Tags []string
    ImageName string
    Identifier int
}

func NewPost(ref *PostReference) *Post {
    //filename := ref.Filename()
    // markdownContent, err := ioutil.ReadFile(filename)
    //util.Check(err)
    return nil
}


func ExtractPreamble(postContent string) *PostPreamble {
    return nil
}