package core

import (
    "fmt"
	"github.com/CheckmateBasilisk/teste-meli-2/backend/routes"
	"github.com/CheckmateBasilisk/teste-meli-2/backend/server"
	//"github.com/CheckmateBasilisk/teste-meli-2/backend/routes"
)

// has the config for accessing db and running server
type Core struct {
    server struct {
        baseUrl string
        port int
    }
    db struct {

    }
}

// setup config
func NewCore(baseUrl string, port int) *Core {
    core := &Core{}

    core.server.baseUrl = baseUrl
    core.server.port = port

    return core
}

// instantiate things based on config and run
func (core *Core) Run() error {
    server := server.NewServer(core.server.baseUrl, core.server.port, routes.BuildRouter())

    err := server.ListenAndServe()
    if err != nil {
        return fmt.Errorf("error serving %w", err)
    }
    return nil
}
