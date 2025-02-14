package main

import (
	"fmt"
    "context"
	"io"
    "log"
	"net/http"

    _ "github.com/jackc/pgx/v5/stdlib"

    "api-meli/config"
)

func main() {
    conf := config.GetConfig(context.Background())


    //TODO: remove this router from here, send to /api/router
    router := http.NewServeMux()
    router.HandleFunc("/hello", sayHello)

    server := &http.Server{
        Addr: fmt.Sprintf(":%s", conf.Server.Port),
        Handler: router,
    }

    log.Printf("starting server at %s\n", server.Addr)
    err := server.ListenAndServe()
    if err != nil {
        log.Fatalf("Server startup failed: %v", err)
    }
}

func sayHello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello, World!")
}
