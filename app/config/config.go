package config

import (
    "fmt"
    "github.com/sirupsen/logrus"
    "path/filepath"
)

type Config struct {
    PagesRoot string
    TemplatesRoot string
    LoggingLevel logrus.Level
}

var ServerConfig *Config

func (c *Config) GetPostFilePath(name string) string {
    return filepath.Join(c.PagesRoot, fmt.Sprintf("%s.md", name))
}

func (c *Config) GetTemplateFilePath(name string) string {
    return filepath.Join(c.TemplatesRoot, fmt.Sprintf("%s.html", name))
}