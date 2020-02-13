package config

import (
    "fmt"
    "path/filepath"
    "strings"
)

var root = "/static"

func image(name string) string { return filepath.Join(root, "images", name) }
func css(name string) string   { return filepath.Join(root, "styles", name) }
func js(name string) string    { return filepath.Join(root, "js", name)     }

type Assets struct {
    Images map[string]string
    Styles map[string]string
    JS map[string]string
    FontsURL string
}

var googleFonts = strings.Join([]string{
    "Average",
    "Source+Serif+Pro:400,700",
    "Quantico",
    "Roboto",
    "Oswald",
    "Roboto+Slab:400,700",
    "Ubuntu",
    "Ubuntu+Mono",
}, "|")

var fontsURL = fmt.Sprintf("https://fonts.googleapis.com/css?family=%s&display=swap", googleFonts)

var DefaultAssets = Assets{
    Images: map[string]string{
        "TopHeaderBackground": image("blueprint.png"),
        "TopHeaderBackgroundOverlay": image("writings.png"),
        "LogoMedium": image("logo/medium.png"),
        "LogoTowardsdatascience": image("logo/medium.png"),
        "LogoHackernoon": image("logo/hackernoon.png"),
        "NetworkGithub": image("links/github.png"),
        "NetworkStackOverflow": image("links/stackoverflow.png"),
        "NetworkMedium": image("links/medium.png"),
        "NetworkUpwork": image("links/upwork.png"),
        "NetworkLinkedIn": image("links/linkedin.png"),
    },
    Styles: map[string]string{
        "Reset": css("reset.css"),
        "PageHeader": css("topmost_header.css"),
        "PostsPreviewCards": css("posts.css"),
        "Navigation": css("navigation.css"),
        "Code": css("code.css"),
    },
    JS: map[string]string {
        "Main": js("lib.js"),
    },
    FontsURL: fontsURL,
}