package main

import (
    "blogapp/app/config"
    "blogapp/app/server"
    "context"
    log "github.com/sirupsen/logrus"
    "net/http"
    "os"
)

func main() {
    srv := server.New()
    log.Debugf("starting listening the address: %s", srv.Addr)
    if err := srv.ListenAndServe(); err != http.ErrServerClosed {
        log.Fatalf("server error: %s", err)
    }
    _ = srv.Shutdown(context.TODO())
}

func init() {
    config.ServerConfig = config.FromEnvironment()
    log.SetOutput(os.Stdout)
    log.SetLevel(config.ServerConfig.LoggingLevel)
}

