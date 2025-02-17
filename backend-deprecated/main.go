package main

import (
    "fmt"
    "encoding/json"
    "net/http"

    "github.com/CheckmateBasilisk/teste-meli-2/backend/core"
)

func main() {
    url := "localhost"
    port := 6969

    app := core.NewCore(url, port)
    err := app.Run()
    if err != nil {
        fmt.Printf("error: %w", err)
    }
}

func readFoos(w http.ResponseWriter, r *http.Request) {
    foos := [...]int{1,2,3}
    response , _ := json.Marshal(foos)

    w.Write([]byte(string(response)))
}
