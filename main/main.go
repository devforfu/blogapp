package main

import (
    "blogapp/app"
    "context"
    "log"
    "net/http"
)

func main() {
    server := app.New()
    if err := server.ListenAndServe(); err != http.ErrServerClosed {
        log.Fatalf("server error: %s", err)
    }
    _ = server.Shutdown(context.TODO())
}
