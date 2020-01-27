package main

import (
    "blogapp/app"
    "blogapp/app/config"
    "context"
    log "github.com/sirupsen/logrus"
    "net/http"
    "os"
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
    config.ServerConfig = config.FromEnvironment()
    log.SetOutput(os.Stdout)
    log.SetLevel(config.ServerConfig.LoggingLevel)
}

