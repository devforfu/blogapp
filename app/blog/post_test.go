package blog

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestReadPreamble(t *testing.T) {
    jsonPreamble := `{
    "category": "blog",
    "title": "Test Title",
    "tags": ["tag1", "tag2"],
    "identifier": 0
}
`
    postContent := `
# Example

An example of post content that goes right after *preamble*.
`
    markdownFileContent := fmt.Sprintf("```yaml\n%s```\n%s", jsonPreamble, postContent)

    preamble := ExtractPreamble(markdownFileContent)

    assert.NotNil(t, preamble)
}