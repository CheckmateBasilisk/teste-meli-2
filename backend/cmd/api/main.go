package main

import (
	"fmt"
    "context"
    "log"
	"net/http"

    "github.com/jackc/pgx/v5/pgxpool"

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
    ctx := context.Background()
    // getting config values
    conf := config.GetConfig(ctx)

    dbPool, dbErr := pgxpool.New(ctx, conf.Db.Url)
    if dbErr != nil {
        log.Fatalf("Database connection failed: %v\n", dbErr)
    }
	defer dbPool.Close()

    // building server
    router := router.NewRouter(ctx, dbPool)
    server := &http.Server{
        Addr: fmt.Sprintf(":%s", conf.Server.Port),
        Handler: router,
    }

    //starting server
    log.Printf("starting server at %s\n", server.Addr)
    err := server.ListenAndServe()
    if err != nil {
        log.Fatalf("Server startup failed: %v", err)
    } }
