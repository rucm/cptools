package main

import (
	"log"
	"os"

	"github.com/rucm/cptools/command"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()

	app.Name = "cptools"
	app.Version = "0.1.0"

	app.Commands = command.GetCommandList()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
