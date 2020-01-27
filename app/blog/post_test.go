package blog

import (
    "blogapp/app/config"
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

func init() {
    config.ServerConfig = config.FromEnvironment()
}

func TestReadPreamble(t *testing.T) {
    jsonPreamble := `
{
    "category": "blog",
    "title": "Test Title",
    "tags": ["tag1", "tag2"],
    "identifier": 0
}
`
    postContent := `
<!--preamble-->

# Example

An example of post content that goes right after *preamble*.
` + "```python\nx=1\n```"

    expectedPost := "# Example\n\n" +
                    "An example of post content that goes right after *preamble*.\n" +
                    "```python\nx=1\n```"

    markdownFileContent := fmt.Sprintf("```json%s```%s", jsonPreamble, postContent)

    preamble, postWithoutPreamble, err := ExtractPreamble(markdownFileContent)

    assert.Nil(t, err)
    assert.NotNil(t, preamble)
    assert.Equal(t, expectedPost, postWithoutPreamble)
}