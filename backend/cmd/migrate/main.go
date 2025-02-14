package main

import (
    "fmt"
    "log"
    "os"
    "flag"
    "context"

    _ "github.com/jackc/pgx/v5/stdlib"
    "github.com/pressly/goose/v3"

    "api-meli/config"
)


var flags = flag.NewFlagSet("migrate", flag.ExitOnError)
var dir   = flags.String("dir", "migrations", "directory with migration files")

// this basically calls goose from inside go (and inside the project)
func main(){
    conf := config.GetConfig(context.Background())

    const dialect  = "pgx"
    var dbString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.Db.Host, conf.Db.User, conf.Db.Password, conf.Db.Name, conf.Db.Port)

    //FIXME: remove this
    fmt.Printf("loaded config:%v\n", conf)
    fmt.Printf("dbString: %s", dbString)

	flags.Parse(os.Args[1:])
	args := flags.Args()

	command := args[0]

    //setting db connection
	db, err := goose.OpenDBWithDriver(dialect, dbString)
    if err != nil {
        log.Fatalf("goose: failed to open DB: %v", err.Error())
    }
    defer func() { // defers closing the db
        err = db.Close()
        if err != nil {
            log.Fatalf("goose: failed to close DB: %v", err.Error())
        }
    }()

    ctx := context.Background()
    err = goose.RunContext(ctx, command, db, *dir, args[1:]...)
    if err != nil {
        log.Fatalf("migrate %v [failed]\n\t%v", command, err)
    }
}

// TODO: make better error msgs...

