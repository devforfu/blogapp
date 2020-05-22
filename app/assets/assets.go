package assets

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)

type Assets struct {
    Images map[string]string `json:"images"`
    Styles map[string]string `json:"styles"`
    JS map[string]string     `json:"scripts"`
    Fonts []string           `json:"required_fonts"`
}

const googleFontsURL = "https://fonts.googleapis.com/css?family=%s&display=swap"

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
    return &assets
}

var DefaultAssets *Assets
