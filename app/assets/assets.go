package assets

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "path/filepath"
    "strings"
)

type Assets struct {
    Images map[string]string    `json:"images"`
    Styles map[string]string    `json:"styles"`
    JS map[string]string        `json:"scripts"`
    Fonts []string              `json:"required_fonts"`
    SyntaxHighlightTheme string `json:"syntax_highlight_theme"`
}

const googleFontsURL = "https://fonts.googleapis.com/css?family=%s&display=swap"
const root = "/static"

func image(name string) string { return filepath.Join(root, "images", name) }
func css(name string) string   { return filepath.Join(root, "styles", name) }
func js(name string) string    { return filepath.Join(root, "js", name)     }

// FontsURL returns Google Fonts URL to load the required web app's fonts.
func (a *Assets) FontsURL() string {
    return fmt.Sprintf(googleFontsURL, strings.Join(a.Fonts, "|"))
}

// FromJSON loads assets configuration from JSON file.
func FromJSON(filename string) *Assets {
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatalf("Failed reading assets file: %s", err)
    }
    var assets Assets
    err = json.Unmarshal(content, &assets)
    if err != nil {
        log.Fatalf("Failed to un-marshal assets file: %s", err)
    }
    for k, v := range assets.Images { assets.Images[k] = image(v) }
    for k, v := range assets.Styles { assets.Styles[k] = css(v) }
    for k, v := range assets.JS     { assets.JS[k] = js(v) }
    return &assets
}