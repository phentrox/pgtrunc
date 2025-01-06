package main

import (
	"github.com/phentrox/db-pg-trunc/internal/config"
	"github.com/phentrox/db-pg-trunc/internal/postgresql"
	"github.com/phentrox/db-pg-trunc/internal/truncate"
)

func runTruncate() {
	println("Reading Config ...")
	configEntity := config.ReadConfig("pgtrunc.yaml")

	println("Connecting to Database ...")
	postgresql.OpenPostgreSqlConnection(configEntity)

	println("Truncating ...")
	truncate.Truncate(configEntity.Schemas)
	println("Truncation Successfull!")

	println("Closing Database Connection ...")
	postgresql.ClosePostgreSqlConnection()
}
