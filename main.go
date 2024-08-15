package main

import (
	"adr/commands"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:    "adr",
		Usage:   "Work with Architecture Decision Records (ADRs)",
		Version: "0.2.0",
	}

	app.Commands = []*cli.Command{
		&commands.InitCommand,
		&commands.NewCommand,
		&commands.TemplatesCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
