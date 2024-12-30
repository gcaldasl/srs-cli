package main

import (
	"log"

	"github.com/gcaldasl/srs-cli/internal/adapters/primary/cli"
	"github.com/gcaldasl/srs-cli/internal/adapters/secondary/persistence"
	"github.com/gcaldasl/srs-cli/internal/core/ports"
	"github.com/gcaldasl/srs-cli/internal/core/services"
	"github.com/gcaldasl/srs-cli/internal/db"
)

func main() {
		conn, err := db.InitDB()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		var repo ports.CardRepository = persistence.NewSQLiteRepository(conn)
		srsService := services.NewSRSService(repo)
		cli := cli.NewCLI(srsService)
		cli.Run()
}
