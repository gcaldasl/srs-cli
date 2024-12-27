package main

import "github.com/gcaldasl/srs-cli/internal/adapters/cli"

func main() {
	app := cli.NewCLI()
	app.Run()
}
