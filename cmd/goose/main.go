// This is custom goose binary.

package main

import (
	"context"
	"flag"
	"hitalent/app/Services/DatabaseManager"
	"log"
	"os"

	_ "hitalent/database/migrations"

	"github.com/pressly/goose/v3"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", ".", "directory with migration files")
)

func main() {
	if err := flags.Parse(os.Args[1:]); err != nil {
		log.Fatalf("goose: failed to parse flags: %v", err)
	}
	args := flags.Args()

	command := &args[0]

	connectionName := DatabaseManager.GetConnectionName("")
	driverName := DatabaseManager.GetDriverNameByConnectionName(connectionName)
	dsn := DatabaseManager.GetDsn(connectionName)

	db, err := goose.OpenDBWithDriver(driverName, dsn)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v", err)
		}
	}()

	arguments := args[1:]

	ctx := context.Background()
	if err := goose.RunContext(ctx, *command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
