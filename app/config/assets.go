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
}

var DefaultAssets = Assets{
    Images: map[string]string{
        "TopHeaderBackground": image("blueprint.png"),
        "TopHeaderBackgroundOverlay": image("writings.png"),
    },
    Styles: map[string]string{
        "Reset": css("reset.css"),
        "PageHeader": css("page_header.css"),
    },
}