package blog

import (
	"blogapp/app/config"
	util "github.com/devforfu/fastgoing"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func ListPosts() PostsList {
	r := util.MustRegexpMap(mdFilePattern)

	posts := make(PostsList, 0)
	var parseFiles = func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Debugf("failed to parse dir: %s", err)
		} else {
			if match := r.Search(path); len(match) > 0 {
				ref := &PostReference{
					Year:util.MustInt(match["year"]),
					Month:util.MustInt(match["month"]),
					Day:util.MustInt(match["day"]),
					Name:match["name"]}
				post, err := NewPost(ref)
				if err != nil {
					log.Warnf("failed to create post: %s", err)
				} else {
					posts = append(posts, post)
				}
			}
		}
		return nil
	}

	util.Check(filepath.Walk(config.ServerConfig.PagesRoot, parseFiles))
	return posts
}

const mdFilePattern = `(?P<year>\d{4})_(?P<month>\d{2})_(?P<day>\d{2})_(?P<name>[\w\W]+)\.md$`
