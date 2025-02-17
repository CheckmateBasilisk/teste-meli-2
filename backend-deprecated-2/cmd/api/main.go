package main

import (
	"fmt"
    "context"
    "log"
	"net/http"

    _ "github.com/jackc/pgx/v5/stdlib"

    "api-meli/config"
    "api-meli/api/router"
)

//  @title          MELI STORE API
//  @version        1.0
//  @description    This is an api for a digital storefront. Part of meli test.

//  @contact.name   Lucas Bagatini
//  @contact.url    lalilulelo

//  @host       localhost:8080
//  @basePath   /v1
func main() {
    conf := config.GetConfig(context.Background())
    router := router.NewRouter()
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

