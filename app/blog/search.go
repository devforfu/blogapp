package blog

import (
	"blogapp/app/config"
	util "github.com/devforfu/fastgoing"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func ListPosts() []*Post {
	r := util.MustRegexMap(mdFilePattern)

	posts := make([]*Post, 0)
	var parseFiles = func(path string, info os.FileInfo, err error) error {
		if err != nil { return err }
		if match := r.Search(path); len(match) > 0 {
			ref := &PostReference{
				Year:util.MustInt(match["year"]),
				Month:util.MustInt(match["month"]),
				Day:util.MustInt(match["day"]),
				Name:match["name"]}
			post, err := NewPost(ref)
			if err != nil {
				log.Warnf("failed to load the post: %s", path)
			} else {
				posts = append(posts, post)
			}
		}
		return nil
	}

	util.Check(filepath.Walk(config.ServerConfig.PagesRoot, parseFiles))
	return posts
}

const mdFilePattern = `(?P<year>\d{4})_(?P<month>\d{2})_(?P<day>\d{2})_(?P<name>[\w\W]+)\.md$`