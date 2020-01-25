package app

import util "fastgoing"

type assets struct {
    root string
}

var images = map[string]string{
    "imageTopHeaderBackground": "blueprint.png",
    "imageTopHeaderBackgroundOverlay": "writings.png",
}

var Assets = &assets{root:"/static"}

func (a *assets) Image(id string) string {
    name, ok := images[id]
    if !ok { return "" }
    return util.JoinPaths(a.root, "images", name)
}