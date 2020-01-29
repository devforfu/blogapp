package config

import (
    "fmt"
    util "github.com/devforfu/fastgoing"
    "github.com/sirupsen/logrus"
    log "github.com/sirupsen/logrus"
    "path/filepath"
)

const defaultPreambleSeparator = "<!--preamble-->"
const defaultDigestSeparator = "<!--more-->"

type Config struct {
    PagesRoot string
    TemplatesRoot string
    LoggingLevel logrus.Level
    PostPreambleSeparator string
    PostDigestSeparator string
}

var ServerConfig *Config

func FromEnvironment() *Config {
    var (
        cwd = util.WorkDir()
        pagesRoot = util.DefaultEnv("APP_PAGES_ROOT", filepath.Join(cwd, "pages"))
        templatesRoot = util.DefaultEnv("APP_TEMPLATES_ROOT", filepath.Join(cwd, "templates"))
        appVerbosity = util.DefaultEnv("APP_VERBOSITY", "debug")
    )
    var loggingLevel log.Level
    switch appVerbosity {
    case "debug": loggingLevel = log.DebugLevel
    case "info":  loggingLevel = log.InfoLevel
    case "warn":  loggingLevel = log.WarnLevel
    case "error": loggingLevel = log.ErrorLevel
    default:      loggingLevel = log.DebugLevel
    }
    return &Config{
        PagesRoot:             pagesRoot,
        TemplatesRoot:         templatesRoot,
        LoggingLevel:          loggingLevel,
        PostPreambleSeparator: defaultPreambleSeparator,
        PostDigestSeparator:defaultDigestSeparator}
}

func (c *Config) GetPostFilePath(name string) string {
    return filepath.Join(c.PagesRoot, fmt.Sprintf("%s.md", name))
}

func (c *Config) GetTemplateFilePath(name string) string {
    return filepath.Join(c.TemplatesRoot, fmt.Sprintf("%s.html", name))
}