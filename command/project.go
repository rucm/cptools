package command

import (
	"github.com/urfave/cli"
)

func init() {
	commandList = append(commandList, cli.Command{
		Name:  "load",
		Usage: "",
	})

	commandList = append(commandList, cli.Command{
		Name:  "expand",
		Usage: "",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "all, a",
				Usage: "",
			},
		},
	})
}
