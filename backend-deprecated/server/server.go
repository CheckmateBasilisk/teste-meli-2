package server

import (
    "fmt"
    "net/http"
)

// creates a server, attaches injected routes and sets up addr:port
func NewServer(baseUrl string, port int, router http.Handler) *http.Server {
    server := &http.Server {
        Addr: fmt.Sprintf("%s:%d", baseUrl, port),
        Handler: router,
    }
    return server
}
