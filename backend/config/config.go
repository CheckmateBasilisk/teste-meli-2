package config

import (
	"context"
	"log"
    "fmt"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
    Server struct {

        Port string `env:"SERVER_PORT, required"`
        Debug bool `env:"SERVER_DEBUG, default=false"`
    }
    Db struct {
        Host string `env:"DB_HOST, required"`
        Port string `env:"DB_PORT, required"`
        User string `env:"DB_USER, required"`
        Name string `env:"DB_NAME, required"`
        Password string `env:"DB_PASSWORD, required"`
        Debug bool `env:"DB_DEBUG, default=false"`
        Url string
    }
}

// returns config from env variables
func GetConfig(ctx context.Context) Config {
    var c Config

    err := envconfig.Process(ctx, &c)
    if err != nil {
        log.Fatalf("Error loading config variables: %v\n", err)
    }
    c.Db.Url = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", c.Db.User, c.Db.Password, c.Db.Host, c.Db.Port, c.Db.Name)
    return c
}
