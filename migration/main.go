package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gadhittana01/socialmedia/config"
	"github.com/gadhittana01/socialmedia/helper"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
)

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
Usage:
  go run *.go <command> [args]
`

func main() {
	config := &config.GlobalConfig{}
	flag.Usage = usage
	flag.Parse()

	helper.LoadConfig(config)

	dbConn := config.DB

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", dbConn.Host, dbConn.Port),
		Database: dbConn.Name,
		User:     dbConn.User,
		Password: dbConn.Password,
	})

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		log.Fatal(err)
	}

	if newVersion != oldVersion {
		log.Println("migrated from version", oldVersion, "to", newVersion)
	} else {
		log.Println("version is", oldVersion)
	}
}

func usage() {
	log.Printf(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}
