package config

import (
    "fmt"
    util "github.com/devforfu/fastgoing"
    "github.com/sirupsen/logrus"
    log "github.com/sirupsen/logrus"
    "path/filepath"
)

<<<<<<< HEAD
const defaultPreambleString = "<!--preamble-->"
=======
const defaultPreambleSeparator = "<!--preamble-->"
const defaultDigestSeparator = "<!--more-->"
>>>>>>> 388be5b5df69f7eb32879a041cd10a09bdcfdf01

type Config struct {
    PagesRoot string
    TemplatesRoot string
    LoggingLevel logrus.Level
    PostPreambleSeparator string
<<<<<<< HEAD
=======
    PostDigestSeparator string
>>>>>>> 388be5b5df69f7eb32879a041cd10a09bdcfdf01
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
<<<<<<< HEAD
        PagesRoot:pagesRoot,
        TemplatesRoot:templatesRoot,
        LoggingLevel:loggingLevel,
        PostPreambleSeparator:defaultPreambleString}
=======
        PagesRoot:             pagesRoot,
        TemplatesRoot:         templatesRoot,
        LoggingLevel:          loggingLevel,
        PostPreambleSeparator: defaultPreambleSeparator,
        PostDigestSeparator:defaultDigestSeparator}
>>>>>>> 388be5b5df69f7eb32879a041cd10a09bdcfdf01
}

func (c *Config) GetPostFilePath(name string) string {
    return filepath.Join(c.PagesRoot, fmt.Sprintf("%s.md", name))
}

func (c *Config) GetTemplateFilePath(name string) string {
    return filepath.Join(c.TemplatesRoot, fmt.Sprintf("%s.html", name))
}