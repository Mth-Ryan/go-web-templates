package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Mth-Ryan/go-web-templates/pkg/conf"
	"github.com/Mth-Ryan/go-web-templates/internal/infra/data"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func printUsage(scriptName string) {
	fmt.Printf(`Usage: %[1]s [up|down]

%[1]s will use the "migrations" folder in
the current directory to perform a full migration
operation using the golang-migrate/migrate lib.

up      Perform a full migration up
down    Perform a full migration down
	`, scriptName)
}

func getMigrate() *migrate.Migrate {
	appConf, err := conf.NewAppConf()
	if err != nil {
		log.Fatal(err)
	}

	db, err := data.NewDatabase(appConf)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db.Ctx, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	return m
}

func main() {
	args := os.Args
	scriptName := filepath.Base(args[0])

	if (len(args) < 2) {
		printUsage(scriptName)
		os.Exit(1)
	}

	switch args[1] {
	case "up":
		m := getMigrate()
		m.Up()
	case "down":
		m := getMigrate()
		m.Down()
	default:
		printUsage(scriptName)
		os.Exit(1)
	}
}
