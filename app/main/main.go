package main

import (
    "context"
    "github.com/devforfu/blogapp/app/assets"
    "github.com/devforfu/blogapp/app/config"
    "github.com/devforfu/blogapp/app/server"
    util "github.com/devforfu/fastgoing"
    log "github.com/sirupsen/logrus"
    "net/http"
    "os"
)

func main() {
    defer log.Exit(0)
    srv := server.New()
    log.Debugf("starting listening the address: %s", srv.Addr)
    if err := srv.ListenAndServe(); err != http.ErrServerClosed {
        log.Fatalf("server error: %s", err)
    }
    _ = srv.Shutdown(context.TODO())
}

func init() {
    config.ServerConfig = config.FromEnvironment()
    log.SetOutput(os.Stderr)
    log.SetLevel(config.ServerConfig.LoggingLevel)
    config.ServerConfig.Assets = assets.FromJSON(config.ServerConfig.StaticFilesMap)
    config.ServerConfig.Validate()
    for _, line := range util.MustVerboseWithSplit(config.ServerConfig) {
        log.Debug(line)
    }
}