package main

import (
	"log"

	"github.com/gcaldasl/srs-cli/internal/adapters/cli"
	"github.com/gcaldasl/srs-cli/internal/adapters/db"
	"github.com/gcaldasl/srs-cli/internal/core/ports"
	"github.com/gcaldasl/srs-cli/internal/core/services"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
		conn, err := db.InitDB()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		repo := ports.NewSQLiteRepository(conn)
		srsService := services.NewSRSService(repo)
		cli := cli.NewCLI(srsService)
		cli.Run()
}
