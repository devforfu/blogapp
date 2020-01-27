package main

import (
    "blogapp/app"
    "context"
    util "github.com/devforfu/fastgoing"
    log "github.com/sirupsen/logrus"
    "net/http"
    "os"
    "path/filepath"
)

func main() {
    server := app.New()
    log.Debugf("starting listening the address: %s", server.Addr)
    if err := server.ListenAndServe(); err != http.ErrServerClosed {
        log.Fatalf("server error: %s", err)
    }
    _ = server.Shutdown(context.TODO())
}

func init() {
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
    log.SetOutput(os.Stdout)
    log.SetLevel(loggingLevel)
    app.ServerConfig = &app.Config{
        PagesRoot: pagesRoot,
        TemplatesRoot: templatesRoot,
        LoggingLevel:loggingLevel}
}
