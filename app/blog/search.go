package blog

import (
	"blogapp/app/config"
	util "github.com/devforfu/fastgoing"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func ListPosts() []*Post {
	r := RegexpMap{regexp.MustCompile(mdFilePattern)}
	util.Check(filepath.Walk(config.ServerConfig.PagesRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil { return err }
		if match := r.Search(path); len(match) > 0 {
			log.Debugf("matched: %v", match)
		}
		return nil
	}))
	return nil
}


type RegexpMap struct {
	compiled *regexp.Regexp
}

func (r *RegexpMap) Search(value string) map[string]string {
	matched := r.compiled.FindStringSubmatch(value)
	params := make(map[string]string)
	for i, name := range r.compiled.SubexpNames() {
		if i > 0 && i <= len(matched) {
			params[name] = matched[i]
		}
	}
	return params
}

func MustInt(number string) int {
	n, err := strconv.ParseInt(number, 10, 32)
	if err != nil { panic(err) }
	return int(n)
}

const mdFilePattern = `(?P<year>\d{4})_(?P<month>\d{2})_(?P<day>\d{2})_(?P<name>[\w\W]+)\.md$`