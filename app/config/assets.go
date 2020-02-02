package config

import (
    "path/filepath"
)

var root = "/static"

func image(name string) string { return filepath.Join(root, "images", name) }
func css(name string) string   { return filepath.Join(root, "styles", name) }

type Assets struct {
    Images map[string]string
    Styles map[string]string
    FontsURL string
}

var DefaultAssets = Assets{
    Images: map[string]string{
        "TopHeaderBackground": image("blueprint.png"),
        "TopHeaderBackgroundOverlay": image("writings.png"),
        "Portrait": image("portrait.jpeg"),
    },
    Styles: map[string]string{
        "Reset": css("reset.css"),
        "PageHeader": css("topmost_header.css"),
        "PostsPreviewCards": css("posts.css"),
        "Navigation": css("navigation.css"),
        "About": css("about.css"),
    },
    FontsURL: "https://fonts.googleapis.com/css?family=Average|Ubuntu|Ubuntu+Mono&display=swap",
}