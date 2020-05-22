package config

import (
    "fmt"
    "github.com/devforfu/blogapp/app/assets"
    util "github.com/devforfu/fastgoing"
    "github.com/sirupsen/logrus"
    log "github.com/sirupsen/logrus"
    "path/filepath"
)

const defaultPreambleSeparator = "<!--preamble-->"
const defaultDigestSeparator = "<!--more-->"

// Config stores information about server parameters and templates
//
// The configuration defines locations where blog's templates and pages are
// stored. The current working directory is parsed by default. However, if the
// app is installed as a binary file, the explicitly set environment variables
// should be used.
type Config struct {
    PagesRoot string
    TemplatesRoot string
    StaticRoot string
    StaticFilesMap string
    LoggingLevel logrus.Level
    PostPreambleSeparator string
    PostDigestSeparator string
    Assets *assets.Assets
}

// Validate ensures that the given configuration doesn't lead to errors.
func (c *Config) Validate() {
    paths := []string{c.TemplatesRoot, c.PagesRoot, c.StaticRoot, c.StaticFilesMap}
    for _, path := range paths {
        if ok, _ := util.Exists(c.TemplatesRoot); !ok {
            log.Fatalf("file is not found: %s", path)
        }
    }
}

// ServerConfig stores global app's configuration
var ServerConfig *Config

// FromEnvironment constructs configuration from environment variables.
func FromEnvironment() *Config {
    var (
        cwd = util.WorkDir()
        pagesRoot = util.DefaultEnv("APP_PAGES_ROOT", filepath.Join(cwd, "pages"))
        templatesRoot = util.DefaultEnv("APP_TEMPLATES_ROOT", filepath.Join(cwd, "templates"))
        staticRoot = util.DefaultEnv("APP_STATIC_ROOT", filepath.Join(cwd, "static"))
        staticFilesMap = util.DefaultEnv("APP_STATIC_FILES_MAP", filepath.Join(cwd, "assets.json"))
        appVerbosity = util.DefaultEnv("APP_VERBOSITY", "debug")
        postPreambleSeparator = util.DefaultEnv("APP_POST_PREAMBLE_SEP", defaultPreambleSeparator)
        postDigestSeparator = util.DefaultEnv("APP_POST_DIGEST_SEP", defaultDigestSeparator)
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
        StaticRoot:            staticRoot,
        StaticFilesMap:        staticFilesMap,
        LoggingLevel:          loggingLevel,
        PostPreambleSeparator: postPreambleSeparator,
        PostDigestSeparator:   postDigestSeparator}
}

func (c *Config) GetPostFilePath(name string) string {
    return filepath.Join(c.PagesRoot, fmt.Sprintf("%s.md", name))
}

func (c *Config) GetTemplateFilePath(name string) string {
    return filepath.Join(c.TemplatesRoot, fmt.Sprintf("%s.html", name))
}