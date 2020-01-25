// https://dev.to/hauxe/golang-http-serve-static-files-correctly-2oj2
package app

import (
    "net/http"
    "strings"
)

// RestrictedFileSystem prevents from rendering folders structure via browser.
type RestrictedFileSystem struct {
    base http.FileSystem
}

func (fs RestrictedFileSystem) Open(path string) (file http.File, err error) {
    f, err := fs.base.Open(path)
    if err != nil { return }

    stat, err := f.Stat()
    if err != nil { return }

    if stat.IsDir() {
        _, err = fs.openIndexFile(path)
        if err != nil { return }
    }

    return f, nil
}

func (fs RestrictedFileSystem) openIndexFile(dir string) (http.File, error) {
    indexFilePath := strings.TrimSuffix(dir, "/") + "/main.html"
    return fs.base.Open(indexFilePath)
}
